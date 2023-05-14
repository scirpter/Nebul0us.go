package event_handle

import (
	"fmt"
	"neb/src/base/client/specs"
	"neb/src/base/enums"
	"neb/src/base/models"
	"neb/src/utils/cnsl"
	"strings"
)

type EventType string

const (
	OnConnected                           EventType = "onConnected"
	OnPlasmaExploitConnected              EventType = "onPlasmaExploitConnected"
	OnClanChatMessage                     EventType = "onClanChatMessage"
	OnPlayerJoin                          EventType = "onPlayerJoin"
	OnGameChatMessage                     EventType = "onGameChatMessage"
	OnSessionStats                        EventType = "onSessionStats"
	OnJoinResult                          EventType = "onJoinResult"
	OnPlasmaExploitJoinResult             EventType = "onPlasmaExploitJoinResult"
	OnEnterGameResult                     EventType = "onEnterGameResult"
	OnPlasmaExploitEnterGameResult        EventType = "onPlasmaExploitEnterGameResult"
	OnBattleRoyaleStart                   EventType = "onBattleRoyaleStart"
	OnPlayerSplit                         EventType = "onPlayerSplit"
	OnBattleRoyaleListResult              EventType = "onBattleRoyaleListResult"
	OnPlasmaExploitBattleRoyaleListResult EventType = "onPlasmaExploitBattleRoyaleStatusUpdate"
)

type EventData map[string]interface{}

type Event struct {
	Type EventType
	Data EventData
}

type Listener chan *Event

type EventDispatcher struct {
	net           *specs.Net
	blobData      *specs.BlobData
	statery       *specs.Statery
	enterGameFunc *func(doDelay bool)
	Listeners     map[EventType]*Listener
}

func NewEventDispatcher(net *specs.Net, blobData *specs.BlobData, statery *specs.Statery, enterGameFunc *func(doDelay bool)) *EventDispatcher {
	return &EventDispatcher{net: net, blobData: blobData, statery: statery, Listeners: make(map[EventType]*Listener), enterGameFunc: enterGameFunc}
}

func (ed *EventDispatcher) AddListener(eventType EventType, listener Listener) {
	ed.Listeners[eventType] = &listener
}

func (ed *EventDispatcher) RemoveListener(eventType EventType) {
	delete(ed.Listeners, eventType)
}

func (ed *EventDispatcher) Dispatch(event Event) {
	if listener, ok := ed.Listeners[event.Type]; ok {
		*listener <- &event
	}
}

func (ed *EventDispatcher) Run() {
	ed.AddListener(OnConnected, make(Listener))
	ed.AddListener(OnPlasmaExploitConnected, make(Listener))
	ed.AddListener(OnClanChatMessage, make(Listener))
	ed.AddListener(OnPlayerJoin, make(Listener))
	ed.AddListener(OnGameChatMessage, make(Listener))
	ed.AddListener(OnSessionStats, make(Listener))
	ed.AddListener(OnJoinResult, make(Listener))
	ed.AddListener(OnPlasmaExploitJoinResult, make(Listener))
	ed.AddListener(OnEnterGameResult, make(Listener))
	ed.AddListener(OnPlasmaExploitEnterGameResult, make(Listener))
	ed.AddListener(OnBattleRoyaleStart, make(Listener))
	ed.AddListener(OnPlayerSplit, make(Listener))
	ed.AddListener(OnBattleRoyaleListResult, make(Listener))
	ed.AddListener(OnPlasmaExploitBattleRoyaleListResult, make(Listener))

	go func() {
		for {
			select {

			case <-*ed.Listeners[OnConnected]:
				masterProxyPrefix := "FREE"
				slaveProxyPrefix := "FREE"

				if ed.net.MasterProxy != nil {
					masterProxyPrefix = strings.ToUpper(*ed.net.MasterProxy.Prefix)
				}
				if ed.net.SlaveProxy != nil {
					slaveProxyPrefix = strings.ToUpper(*ed.net.SlaveProxy.Prefix)
				}
				cnsl.Ok(fmt.Sprintf("%s connected -> (M %s & S %s)", *ed.blobData.Name, masterProxyPrefix, slaveProxyPrefix))
				ed.net.ConnectionState = enums.CONNECTION_STATE_CONNECTED

				if ed.statery.IsPlasmaExploitEnabled {
					*ed.Listeners[OnPlasmaExploitConnected] <- &Event{Type: OnPlasmaExploitConnected}
				}

			case e := <-*ed.Listeners[OnClanChatMessage]:
				playerName := e.Data["playerName"].(*string)
				message := e.Data["message"].(*string)
				playerClanRank := e.Data["playerClanRank"].(enums.CLAN_RANK)
				_ = playerClanRank
				playerAccountID := e.Data["playerAccountID"].(int32)

				if ed.net.Uniquifier == 0 {
					cnsl.LogLowPrio(fmt.Sprintf("%s (%d) sent clan chat -> %s", *playerName, playerAccountID, *message))
				}

			case e := <-*ed.Listeners[OnPlayerJoin]:
				player := e.Data["player"].(*models.Player)
				shortPacketType := e.Data["shortPacketType"].(string)
				_ = shortPacketType

				if ed.net.Uniquifier == 0 {
					cnsl.Log(fmt.Sprintf("%s (%d) joined", *player.Name, player.AccountID))
				}

			case e := <-*ed.Listeners[OnGameChatMessage]:
				playerName := e.Data["playerName"].(*string)
				message := e.Data["message"].(*string)
				playerAccountID := e.Data["playerAccountID"].(int32)

				if ed.net.Uniquifier == 0 {
					cnsl.LogLowPrio(fmt.Sprintf("%s (%d) sent game chat -> %s", *playerName, playerAccountID, *message))
				}

			case <-*ed.Listeners[OnSessionStats]:
				if ed.statery.IsAutoRejoining &&
					ed.net.ConnectionState == enums.CONNECTION_STATE_CONNECTED &&
					// lobby has ended, ignore packet
					(!(ed.net.World.TimeLeft >= 65500 && ed.net.World.TimeLeft < 65535) || (ed.net.World.TimeLeft == 32767)) {
					(*ed.enterGameFunc)(true)
				}

			case e := <-*ed.Listeners[OnJoinResult]:
				player := e.Data["player"].(*models.Player)

				if ed.statery.IsPlasmaExploitEnabled && *player.Name == *ed.blobData.Name {
					*ed.Listeners[OnPlasmaExploitJoinResult] <- &Event{Type: OnPlasmaExploitJoinResult}
				}

			case e := <-*ed.Listeners[OnEnterGameResult]:
				result := e.Data["result"].(enums.JOIN_RESULT)
				roomName := e.Data["roomName"].(*string)
				gameMode := e.Data["gameMode"].(enums.GAME_MODE)
				_ = gameMode
				_ = roomName

				if ed.net.Uniquifier == 0 {
					cnsl.Log(fmt.Sprintf("enter game result -> %s", result.String()))

					if result != enums.JOIN_RESULT_SUCCESS {
						ed.net.World = ed.net.CachedWorld.Copy()
					}
					ed.net.CachedWorld = nil
				}

				// if ed.statery.IsPlasmaExploitEnabled {
				// 	*ed.Listeners[OnPlasmaExploitEnterGameResult] <- &Event{Type: OnPlasmaExploitEnterGameResult, Data: e.Data}
				// }

			case <-*ed.Listeners[OnBattleRoyaleStart]:

			case e := <-*ed.Listeners[OnPlayerSplit]:
				player := e.Data["player"].(*models.Player)
				_ = player

			case e := <-*ed.Listeners[OnBattleRoyaleListResult]:
				if ed.statery.IsPlasmaExploitEnabled {
					*ed.Listeners[OnPlasmaExploitBattleRoyaleListResult] <- &Event{Type: OnPlasmaExploitBattleRoyaleListResult, Data: e.Data}
				}
			}
		}
	}()
}
