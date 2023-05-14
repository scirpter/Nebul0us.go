package client

import (
	"encoding/hex"
	"fmt"
	"math"
	"math/rand"
	event_handle "neb/src/base/client/events"
	"neb/src/base/client/specs"
	"neb/src/base/enums"
	"neb/src/base/models"
	"neb/src/base/packets"
	"neb/src/common"
	mnet "neb/src/net"
	"neb/src/utils/cnsl"
	"net"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	app                    *App
	blobData               *specs.BlobData
	controlData            *specs.ControlData
	net                    *specs.Net
	eventDispatcher        *event_handle.EventDispatcher
	cachedEnterGameRequest *packets.EnterGameRequest
	Statery                *specs.Statery
}

func NewClient(app *App, uniquifier uint8, name string, ticket string, serverIP enums.SERVER, slaveProxy *mnet.Proxy, masterProxy *mnet.Proxy) *Client {
	randSrc := rand.NewSource(time.Now().UnixNano())
	RNG := rand.New(randSrc)

	c := &Client{
		app: app,
		blobData: &specs.BlobData{
			Name:              &name,
			AccountID:         -1,
			Token:             &ticket,
			Skin:              app.InitialSkin,
			ColorCycle:        enums.COLOR_CYCLE_NONE,
			NameFont:          enums.FONT_DEFAULT,
			Halo:              common.ValuedUndefined,
			Hat:               common.ValuedUndefinedHigh,
			Particle:          common.ValuedUndefinedHigh,
			ProfileVisibility: enums.PROFILE_VISIBILITY_APPEAR_OFFLINE,
			EjectSkin:         common.ValuedUndefined,
		},
		controlData: &specs.ControlData{
			Speed:       0,
			Angle:       0,
			EjectCt:     0,
			SplitCt:     0,
			DropCt:      0,
			ControlTick: 0,
		},
		net: &specs.Net{
			Uniquifier:      uniquifier,
			ConnectionState: enums.CONNECTION_STATE_DISCONNECTED,
			ServerIP:        serverIP,
			Sock:            nil,
			World:           models.NewWorld(),
			Cr2Token1:       0,
			Cr2Token2:       0,
			RngToken1:       RNG.Uint32(),
			RngToken2:       0xff000000 | RNG.Uint32()&0x00ffffff,
			SlaveProxy:      slaveProxy,
			MasterProxy:     masterProxy,
			CachedWorld:     nil,
		},
		Statery: specs.NewStatery(),
	}

	enterGameFunc := c.EnterGame
	c.eventDispatcher = event_handle.NewEventDispatcher(c.net, c.blobData, c.Statery, &enterGameFunc)
	go c.eventDispatcher.Run()

	if slaveProxy != nil && masterProxy != nil {
		slaveProxy.Connect(fmt.Sprintf("%s:%d", string(serverIP), common.SERVER_PORT))
		masterProxy.Connect(fmt.Sprintf("%s:%d", string(serverIP), common.SERVER_PORT))
	}
	sock, _ := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4zero, Port: 0})
	c.net.Sock = sock
	return c
}

func (c *Client) Reset() {
	randSrc := rand.NewSource(time.Now().UnixNano())
	RNG := rand.New(randSrc)

	c.controlData.Speed = 0
	c.controlData.Angle = 0
	c.controlData.EjectCt = 0
	c.controlData.SplitCt = 0
	c.controlData.DropCt = 0
	c.controlData.ControlTick = 0

	c.net.ConnectionState = enums.CONNECTION_STATE_DISCONNECTED
	c.net.World = models.NewWorld()
	c.net.Cr2Token1 = 0
	c.net.Cr2Token2 = 0
	c.net.RngToken1 = RNG.Uint32()
	c.net.RngToken2 = 0xff000000 | RNG.Uint32()&0x00ffffff
	c.net.CachedWorld = nil
}

func (c *Client) RunBackgroundTasks() {
	go c.RunUDPHandleLoop()
	go c.KeepAliveLoop()
	go c.TickLoop()
	go c.EmoteLoop()
	go c.DNDLoop()
	go c.SafePlasmaFarmingLoop()
	go c.PlasmaExploitLoop()
}

func (c *Client) RunUDPHandleLoop() {
	// unproxied
	go func() {
		for {
			buf := make([]byte, 8192)
			_, _, err := c.net.Sock.ReadFromUDP(buf)

			if err != nil {
				cnsl.Error(fmt.Sprintf("error reading from UDP socket -> %s", err.Error()))
				time.Sleep(1 * time.Second)
				continue
			}

			if c.app.InitialDoTracePackets && c.net.Uniquifier == 0 {
				c.app.tracedPackets = append(c.app.tracedPackets, &buf)
			}

			HandlePacket(c, &buf)
		}
	}()

	// proxied
	if c.net.SlaveProxy == nil || c.net.MasterProxy == nil {
		return
	}
	// slave
	go func() {
		for {
			buf := make([]byte, 8192)
			conn := *c.net.SlaveProxy.Conn
			_, err := conn.Read(buf)

			if err != nil {
				cnsl.Error(fmt.Sprintf("proxy error reading from UDP socket -> %s", err.Error()))
				time.Sleep(1 * time.Second)
				continue
			}

			if c.app.InitialDoTracePackets && c.net.Uniquifier == 0 {
				c.app.tracedPackets = append(c.app.tracedPackets, &buf)
			}

			HandlePacket(c, &buf)
		}
	}()

	// master
	go func() {
		for {
			buf := make([]byte, 8192)
			conn := *c.net.MasterProxy.Conn
			_, err := conn.Read(buf)

			if err != nil {
				cnsl.Error(fmt.Sprintf("proxy error reading from UDP socket -> %s", err.Error()))
				time.Sleep(1 * time.Second)
				continue
			}

			if c.app.InitialDoTracePackets && c.net.Uniquifier == 0 {
				c.app.tracedPackets = append(c.app.tracedPackets, &buf)
			}

			HandlePacket(c, &buf)
		}
	}()
}

func (c *Client) AreTokensPresent() bool {
	return c.net.Cr2Token1 != 0 && c.net.Cr2Token2 != 0
}

func (c *Client) IsIdling() bool {
	return c.controlData.Speed == 0 &&
		c.controlData.SplitCt == 0 &&
		c.controlData.EjectCt == 0 &&
		c.controlData.DropCt == 0
}

func (c *Client) SendPacket(packet *[]byte) {
	if c.net.SlaveProxy == nil || c.net.MasterProxy == nil {
		c.net.Sock.WriteToUDP(*packet, &net.UDPAddr{IP: net.ParseIP(string(c.net.ServerIP)), Port: int(common.SERVER_PORT)})
		return
	}

	// we proxy when:
	// - connecting
	// - joining a new lobby
	if c.Statery.WhichProxyType == mnet.PROXY_MASTER {
		conn := *c.net.MasterProxy.Conn
		conn.Write(*packet)
	} else if c.Statery.WhichProxyType == mnet.PROXY_SLAVE {
		conn := *c.net.SlaveProxy.Conn
		conn.Write(*packet)
	}
}

func (c *Client) TickLoop() {
	for {
		if !c.AreTokensPresent() || c.Statery.TaskLock {
			time.Sleep(1 * time.Second)
			continue
		}

		packet := packets.NewControlPacket(c.net, c.controlData)

		var split, eject, drop bool = false, false, false

		// NOTE: our split/eject/drop is stacking faster than we are handling it
		if c.controlData.ControlTick%4 == 0 {
			if c.controlData.SplitCt > 0 {
				split = true
			}
			if c.controlData.EjectCt > 0 {
				eject = true
			}
			if c.controlData.DropCt > 0 {
				drop = true
			}
		}

		player := c.Player()
		if player == nil {
			time.Sleep(1 * time.Second)
			continue
		}

		// Run all functions
		// This will allow us to change the order of execution
		// without changing the code
		RUN_FUNCTIONS(c, player)

		// set the packet split/eject/drop
		packet.Split = split
		packet.Eject = eject
		packet.Drop = drop

		// send the packet
		c.SendPacket(packet.Write())

		// wait until next tick
		time.Sleep(72216 * time.Microsecond) // 72.216 ms
	}
}

func (c *Client) TargetHoles(player *models.Player, holeTypes []enums.HOLE_TYPE) {
	x, y := player.GetAvgPos()
	var holes []*models.Hole

	c.net.World.HoleMutx.RLock()
	for _, hole := range c.net.World.Holes {
		for _, holeType := range holeTypes {
			if hole.Typeof == holeType {
				hole.Dist = float32(math.Pow(float64(hole.X-x), 2) + math.Pow(float64(hole.Y-y), 2))
				holes = append(holes, hole)
			}
		}
	}
	c.net.World.HoleMutx.RUnlock()

	if len(holes) > 0 {
		sort.Slice(holes, func(i, j int) bool {
			return holes[i].Dist < holes[j].Dist
		})
		distX := holes[0].X - x
		distY := holes[0].Y - y
		angle := math.Atan2(float64(distY), float64(distX))
		c.Move(float32(angle), 1)
	}
}

func (c *Client) TargetPlasma(player *models.Player) {
	x, y := player.GetAvgPos()
	var items []*models.Item

	c.net.World.ItemMutx.RLock()
	for _, item := range c.net.World.Items {
		if item.Typeof == enums.ITEM_TYPE_COIN || item.Typeof == enums.ITEM_TYPE_CAKE_PLASMA || item.Typeof == enums.ITEM_TYPE_CAKE_XP {
			item.Dist = float32(math.Pow(float64(item.X-x), 2) + math.Pow(float64(item.Y-y), 2))
			items = append(items, item)
		}
	}
	c.net.World.ItemMutx.RUnlock()

	if len(items) > 0 {
		sort.Slice(items, func(i, j int) bool {
			return items[i].Dist < items[j].Dist
		})
		distX := items[0].X - x
		distY := items[0].Y - y
		angle := math.Atan2(float64(distY), float64(distX))
		c.Move(float32(angle), 1)
	}
}

func (c *Client) TargetPlayer(player *models.Player, target *models.Player) float64 {
	myX, myY := player.GetAvgPos()
	myRad := player.GetRadius()
	targetX, targetY := target.GetAvgPos()
	targetRad := target.GetRadius()
	distX := targetX - myX
	distY := targetY - myY
	dist := math.Pow(float64(distX), 2) + math.Pow(float64(distY), 2)
	circleDist := dist - (myRad + targetRad)
	angle := math.Atan2(float64(distY), float64(distX))
	c.Move(float32(angle), 1)
	return circleDist
}

func (c *Client) KeepAliveLoop() {
	for {
		if !c.AreTokensPresent() || (c.Statery.WhichProxyType == mnet.PROXY_MASTER && c.net.SlaveProxy != nil) {
			time.Sleep(1 * time.Second)
			continue
		}
		packet := packets.NewKeepAlivePacket(c.net)
		c.SendPacket(packet.Write())

		time.Sleep(505502 * time.Microsecond) // ~505.502 ms
	}
}

func (c *Client) EnterGame(doDelay bool) {
	if doDelay {
		go func() {
			// delay because holes are reloaded after 1-1.5s if the bot dies
			time.Sleep(1400 * time.Millisecond)
			packet := packets.NewJoinRequestPacket(c.net, c.blobData)
			c.SendPacket(packet.Write())
		}()
	} else {
		packet := packets.NewJoinRequestPacket(c.net, c.blobData)
		c.SendPacket(packet.Write())
	}
}

func (c *Client) Connect() {
	packet := packets.NewConnectRequest3Packet(c.net, c.blobData, c.app.InitialGameMode, c.app.InitialDoConnectInPrivateGameSearch, c.app.InitialGoMayhem)

	c.Statery.WhichProxyType = mnet.PROXY_MASTER
	c.SendPacket(packet.Write())
	c.net.ConnectionState = enums.CONNECTION_STATE_CONNECTING
	c.Statery.WhichProxyType = mnet.PROXY_SLAVE
}

func (c *Client) Disconnect(waitThreaded bool) {
	packet := packets.NewDisconnectPacket(c.net)
	if waitThreaded {
		go func() {
			c.SendPacket(packet.Write())
			time.Sleep(1 * time.Second)
			c.Reset()
		}()
	} else {
		c.SendPacket(packet.Write())
		time.Sleep(1 * time.Second)
		c.Reset()
	}
}

func (c *Client) Player() *models.Player {
	var player *models.Player

	c.net.World.PlayerMutx.RLock()
	for _, value := range c.net.World.Players {
		if *value.Name == *c.blobData.Name || (value.AccountID == c.blobData.AccountID && value.AccountID != -1) {
			player = value
			break
		}
	}
	c.net.World.PlayerMutx.RUnlock()
	return player
}

func (c *Client) Move(angle float32, speed uint8) {
	c.controlData.Angle = angle
	c.controlData.Speed = speed
}

// alias use item
func (c *Client) Eject() {
	c.controlData.EjectCt++
}

func (c *Client) UseItem() {
	c.controlData.EjectCt++
}

func (c *Client) Split() {
	c.controlData.SplitCt++
}

func (c *Client) Drop() {
	c.controlData.DropCt++
}

func (c *Client) FeedPlayer(playerID *int32, playerName *string) {
	if playerID != nil {
		c.Statery.EmotionalSupportPlayerAccountID = *playerID
	} else if playerName != nil {
		c.Statery.EmotionalSupportPlayerName = playerName
	} else {
		c.Statery.EmotionalSupportPlayerAccountID = -1
		c.Statery.EmotionalSupportPlayerName = nil
	}
}

func (c *Client) SetPlasmaFarmingTarget(playerID int32, playerName *string) {
	if playerID != -1 {
		c.Statery.PlasmaFarmingTargetPlayerAccountID = playerID
	} else if playerName != nil {
		c.Statery.PlasmaFarmingTargetPlayerName = playerName
	} else {
		c.Statery.PlasmaFarmingTargetPlayerAccountID = -1
		c.Statery.PlasmaFarmingTargetPlayerName = nil
	}
}

func (c *Client) Rejoin() {
	if c.cachedEnterGameRequest == nil || !c.AreTokensPresent() {
		return
	}
	// NOTE: didnt proxy test this yet. does this even need one?
	c.SendPacket(c.cachedEnterGameRequest.Write())
}

func (c *Client) Emote(emoteID uint8) {
	if !c.AreTokensPresent() {
		return
	}
	packet := packets.NewEmoteRequestPacket(c.net, emoteID)
	c.SendPacket(packet.Write())
}

func (c *Client) JoinGame(accountID int32, worldName string) {
	packet := packets.NewEnterGameRequestPacket(c.net, accountID, &worldName)
	keepAlivePacket := packets.NewKeepAlivePacket(c.net)

	c.net.CachedWorld = c.net.World.Copy()
	c.net.World.Reset()
	c.cachedEnterGameRequest = packet

	c.Statery.WhichProxyType = mnet.PROXY_MASTER
	c.SendPacket(keepAlivePacket.Write())
	c.SendPacket(packet.Write())
	c.Statery.WhichProxyType = mnet.PROXY_SLAVE
}

func (c *Client) Chat(message *string) {
	if !c.AreTokensPresent() {
		return
	}
	packet := packets.NewGameChatMessagePacket(c.net, c.blobData, nil, nil, message)
	c.SendPacket(packet.Write())
}

func (c *Client) EmoteLoop() {
	for {
		if !c.AreTokensPresent() || !c.Statery.IsEmoteLooping || c.Statery.TaskLock {
			time.Sleep(1 * time.Second)
			continue
		}
		if c.Statery.EmoteLoopID != 0 {
			c.Emote(c.Statery.EmoteLoopID)
		} else {
			// random number from 1-89
			c.Emote(uint8(rand.Intn(92) + 1))
		}
		time.Sleep(150 * time.Millisecond)
	}
}

func (c *Client) DNDLoop() {
	for {
		if !c.AreTokensPresent() || !c.Statery.IsDNDModeEnabled {
			time.Sleep(1 * time.Second)
			continue
		}

		c.Move(float32(c.Statery.DNDModeAngle), 1)
		c.controlData.SplitCt = 1
		controlPacket := packets.NewControlPacket(c.net, c.controlData)
		controlPacket.Split = true
		c.SendPacket(controlPacket.Write())

		time.Sleep(300 * time.Millisecond) // 72.216 ms
	}
}

func (c *Client) SafePlasmaFarmingLoop() {
	for {
		if !c.AreTokensPresent() || !c.Statery.IsFarmingPlasma || c.net.World.GameMode == enums.GAME_MODE_PLASMA_HUNT || c.Statery.TaskLock {
			// Wait until ready to farm plasma
			time.Sleep(1 * time.Second)
			continue
		}
		if c.Statery.PlasmaFarmingSafeResetTimer <= 0 {
			// Reset game after a while
			c.EnterGame(false)
			c.Statery.PlasmaFarmingSafeResetTimer = specs.PLASMA_FARMING_SAFE_RESET_TIMER_DEFAULT
		}
		// Count down
		c.Statery.PlasmaFarmingSafeResetTimer--
		time.Sleep(1 * time.Second)
	}
}

func (c *Client) Rename(name *string) {
	// Search through all the clients
	for _, client := range c.app.Clients {
		// Search through all the players
		client.net.World.PlayerMutx.Lock()
		for _, player := range client.net.World.Players {
			// If the player has the same name as the client
			if *player.Name == *c.blobData.Name {
				// Rename the player
				player.Name = name
			}
		}
		client.net.World.PlayerMutx.Unlock()
	}
	// Rename the client
	c.blobData.Name = name
}

func (c *Client) InjectToken() {
	// Get the token from the shared preferences
	token := GetSharedPrefsToken()

	if token == nil {
		// Error already handled in GetSharedPrefsToken()
		return
	}

	if *token == "nada" {
		// Token not found in shared preferences
		cnsl.Error("no token found. make sure you are running this script on mobile and that you have root. if done correctly, you will get the token of your last logged in account.")
		return
	}

	// Disconnect from the current server
	c.Disconnect(false)

	// Inject the token
	c.blobData.Token = token
	accountID := strings.Split(*token, ",")[0]
	actualAccountID, _ := strconv.Atoi(accountID)
	c.blobData.AccountID = int32(actualAccountID)

	// Connect to the server
	c.Connect()

	cnsl.Ok(fmt.Sprintf("token injected! bot %s is now playing on account %s", *c.blobData.Name, accountID))
}

func (c *Client) GetNetStats() {
	// Print out the network tokens.
	netstats :=
		"\n%s\n" +
			"%s\n" +
			"CR2_1: %s\n" +
			"CR2_2: %s\n" +
			"RNG_1: %s\n" +
			"RNG_2: %s\n"
	cnsl.LogHighPrio(fmt.Sprintf(
		netstats,
		*c.blobData.Name,
		strings.Repeat("-", len(*c.blobData.Name)),
		strconv.FormatUint(uint64(c.net.Cr2Token1), 16),
		strconv.FormatUint(uint64(c.net.Cr2Token2), 16),
		strconv.FormatUint(uint64(c.net.RngToken1), 16),
		strconv.FormatUint(uint64(c.net.RngToken2), 16),
	))
}

func (c *Client) PacketSearch(searchQuery *string) {
	// Iterate over all packets in the tracedPackets array
	for _, packet := range c.app.tracedPackets {

		// Convert the packet to a string
		packetString := hex.EncodeToString(*packet)

		// If the search string is contained in the string representation of the packet
		if strings.Contains(packetString, *searchQuery) {
			// Log the packet
			cnsl.LogHighPrio(fmt.Sprintf("found packet -> %s", packetString))
			goto FOUND
		}
	}
	cnsl.Error("no packet found. maybe packet tracing is disabled?")

FOUND:
}

func (c *Client) BattleRoyaleAction(doRegister bool) {
	if !c.AreTokensPresent() {
		return
	}
	packet := packets.NewBattleRoyaleActionPacket(c.net, doRegister)
	c.SendPacket(packet.Write())
}

func (c *Client) BattleRoyaleListRequest() {
	if !c.AreTokensPresent() {
		return
	}
	packet := packets.NewBattleRoyaleListRequestPacket(c.net)
	c.SendPacket(packet.Write())
}

func (c *Client) SetPlasmaExploit(isEnabled bool) {
	// If the exploit is disabled, disable it and disconnect from the server.
	// If the exploit is enabled, disable it, disconnect from the server, and
	// then immediately reconnect.
	if !isEnabled && c.AreTokensPresent() {
		c.BattleRoyaleAction(false)
		c.Disconnect(false)
	} else if isEnabled && c.AreTokensPresent() {
		// assuming the client is already connected
		c.Disconnect(false)
	}

	// Set the client's exploit state and lock the client's task queue.
	c.Statery.IsPlasmaExploitEnabled = isEnabled
	c.Statery.TaskLock = isEnabled
}

func (c *Client) PlasmaExploitLoop() {
	/* THE PLASMA EXPLOIT CODE IS NOT AVAILABLE TO THE PUBLIC. */
}

func (c *Client) SetDND(direction *string) {
	if direction == nil && c.Statery.IsDNDModeEnabled {
		c.Statery.IsDNDModeEnabled = false
		c.Statery.TaskLock = false
		return
	} else if direction == nil {
		return
	}
	c.Statery.DNDModeAngle = common.CornerDirectionStringToFloat(direction)
	c.Statery.IsDNDModeEnabled = true
	c.Statery.TaskLock = true
}

func (c *Client) GetProxyStats() {
	cnsl.LogHighPrio(fmt.Sprintf("%s -> type %s alias %s", *c.blobData.Name, *c.net.SlaveProxy.Prefix, mnet.PROXY_MASTER.String()))
	cnsl.LogHighPrio(fmt.Sprintf("%s -> type %s alias %s", *c.blobData.Name, *c.net.MasterProxy.Prefix, mnet.PROXY_SLAVE.String()))
}

func (c *Client) Test() {
	cnsl.Error("no testing implemented")
}
