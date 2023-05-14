package main

import (
	"fmt"
	"neb/src/auth"
	"neb/src/base/client"
	"neb/src/base/models/commands"
	"neb/src/common"
	"neb/src/net"
	"neb/src/utils/cnsl"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	cnsl.ClearConsole()

	app := client.NewApp()
	masterProxies := make([]*net.Proxy, app.InitialBotCt)
	slaveProxies := make([]*net.Proxy, app.InitialBotCt)

	if app.AuthAccount.IsPremium() {
		app.ProxyManager.FillBestSpeed()
		masterProxies, slaveProxies = app.ProxyManager.GetPerfectProxyList(app.InitialBotCt)
	}

	wg := new(sync.WaitGroup)
	wg.Add(int(app.InitialBotCt))
	for i := uint8(0); i < app.InitialBotCt; i++ {
		go func(i uint8) {
			var c *client.Client

			botName := ""
			if app.InitialBotName != nil {
				botName = fmt.Sprintf("%s%d", *app.InitialBotName, i)
			} else {
				botName = *app.UsernameManager.GetRandomUsername()
			}
			masterProxy := masterProxies[i]
			slaveProxy := slaveProxies[i]
			if app.AuthAccount.IsPremium() {
				c = client.NewClient(app, uint8(i), botName, ",-", app.InitialServer, slaveProxy, masterProxy)
			} else {
				c = client.NewClient(app, uint8(i), botName, ",-", app.InitialServer, nil, nil)
			}
			app.AddClient(c)
			c.RunBackgroundTasks()
			c.Connect()
			wg.Done()
		}(i)
	}
	// wait for all bots to connect
	wg.Wait()

	for {
		in := cnsl.GetUserInput()
		commandArgs := strings.Split(*in, " ")
		var selectedClient *client.Client = nil
		if len(commandArgs) > 1 && common.IsDigit(commandArgs[0]) {
			clientIndex, _ := strconv.Atoi(commandArgs[0])
			if clientIndex >= 0 && clientIndex < len(app.Clients) {
				selectedClient = app.Clients[clientIndex]
			}
			commandArgs = commandArgs[1:]
		}
		command := app.CommandHandler.GetCommand(commands.CommandName(commandArgs[0]))
		if command == nil {
			cnsl.Error("unknown command. are you sure you typed it correctly? example: jo goodapple123")
			continue
		}
		commandArgs = commandArgs[1:]

		doLog := true

		if command.IsOnCooldown() {
			doLog = false
			cnsl.Error(fmt.Sprintf("command %s on cooldown, please wait %d second(s)", command.Name, command.RemainingCooldownS))
			continue
		}

		for i, c := range app.Clients {
			if selectedClient != nil && c != selectedClient {
				continue
			}
			switch command.Name {

			case commands.CONNECT: // connect
				c.Connect()

			case commands.JOIN: // join
				if len(commandArgs) == 0 {
					goto END
				}
				if common.IsDigit(commandArgs[0]) {
					accountID, _ := strconv.Atoi(commandArgs[0])
					c.JoinGame(int32(accountID), "")
				} else {
					c.JoinGame(-1, strings.Join(commandArgs, " "))
				}

			case commands.ENTER: // enter
				c.EnterGame(false)

			case commands.REJOIN: // rejoin
				c.Rejoin()

			case commands.DISCONNECT_MULTI: // disconnect
				c.Disconnect(true)

			case commands.DISCONNECT_SINGLE: // disconnect (single bot)
				if len(app.Clients) == 0 {
					break
				}
				app.Clients[len(app.Clients)-1].Disconnect(true)

			case commands.FEED: // feed
				if len(commandArgs) == 0 {
					c.FeedPlayer(nil, nil)
					continue
				}
				if common.IsDigit(commandArgs[0]) {
					accountID, _ := strconv.Atoi(commandArgs[0])
					actualAccountID := int32(accountID)
					c.FeedPlayer(&actualAccountID, nil)
				} else {
					playerName := strings.Join(commandArgs, " ")
					c.FeedPlayer(nil, &playerName)
				}

			case commands.SETPLASMAFOR: // set plasma for
				if len(commandArgs) == 0 {
					c.SetPlasmaFarmingTarget(-1, nil)
					continue
				}
				if common.IsDigit(commandArgs[0]) {
					accountID, _ := strconv.Atoi(commandArgs[0])
					actualAccountID := int32(accountID)
					c.SetPlasmaFarmingTarget(actualAccountID, nil)
				} else {
					playerName := strings.Join(commandArgs, " ")
					c.SetPlasmaFarmingTarget(-1, &playerName)
				}

			case commands.EXIT: // exit
				os.Exit(0)

			case commands.CHAT: // chat
				message := strings.Join(commandArgs, " ")
				c.Chat(&message)

			case commands.EMOTE: // emote
				if !common.IsDigit(commandArgs[0]) {
					goto END
				}
				emoteIndex, _ := strconv.Atoi(commandArgs[0])
				c.Emote(uint8(emoteIndex))

			case commands.RENAME: // rename
				name := strings.Join(commandArgs, " ")
				newName := fmt.Sprintf("%s%d", name, i)
				c.Rename(&newName)

			case commands.SETEMOTELOOP: // set emote loop
				if len(commandArgs) != 2 {
					goto END
				}

				if commandArgs[0] == "y" {
					c.Statery.IsEmoteLooping = true
					emoteID, _ := strconv.Atoi(commandArgs[1])
					c.Statery.EmoteLoopID = uint8(emoteID)
				} else {
					c.Statery.IsEmoteLooping = false
					c.Statery.EmoteLoopID = 0
				}

			case commands.SETMASSTHRESHOLD: // set mass threshold
				if len(commandArgs) != 1 {
					goto END
				}
				massThreshold, _ := strconv.Atoi(commandArgs[0])
				c.Statery.MassThreshold = int32(massThreshold)

			case commands.SETAUTOREJOIN: // set autorejoin
				if len(commandArgs) != 1 {
					goto END
				}
				if commandArgs[0] == "y" {
					c.Statery.IsAutoRejoining = true
				} else {
					c.Statery.IsAutoRejoining = false
				}

			case commands.SETLVLMETA: // set lvlmeta
				if len(commandArgs) != 1 {
					goto END
				}
				if commandArgs[0] == "y" {
					c.Statery.IsFarmingHoles = true
				} else {
					c.Statery.IsFarmingHoles = false
				}

			case commands.SETPLASMAFARM: // set plasma farm
				if len(commandArgs) != 1 {
					goto END
				}
				if commandArgs[0] == "y" {
					c.Statery.IsFarmingPlasma = true
				} else {
					c.Statery.IsFarmingPlasma = false
				}

			case commands.CLEAR: // console clear
				cnsl.ClearConsole()

			case commands.INJECTTOKEN: // inject token
				if len(commandArgs) != 1 {
					goto END
				}
				clientIndex, _ := strconv.Atoi(commandArgs[0])
				if clientIndex >= len(app.Clients) {
					goto END
				}
				client := app.Clients[clientIndex]
				client.InjectToken()

			case commands.GETTOKEN: // get token
				token := client.GetSharedPrefsToken()
				if token != nil && *token != "nada" {
					cnsl.Log("put this into your tokens.txt on pc or mobile (do not use \"sign out all accounts\" feature on nebulous anymore, or token will be reset) -> " + *token)
					goto END
				}
				cnsl.Error("no token found. make sure you are running this script on mobile and that you have root. if done correctly, you will get the token of your last logged in account.")
				goto END

			case commands.SPLASMAEXPLOIT: // set plasma exploit
				cnsl.Error("command not supported yet.")
				goto END
				/*
					if len(commandArgs) != 2 {
						break
					}
					app.PlasmaExploitShipType = client.PLASMA_EXPLOIT_SHIP_TYPE(commandArgs[1])
					for _, client := range app.Clients {
						go client.SetPlasmaExploit(commandArgs[0] == "y")
					}
				*/

			case commands.NETSTATS:
				if len(commandArgs) < 1 {
					c.GetNetStats()
				} else {
					clientIndex, _ := strconv.Atoi(commandArgs[0])
					if clientIndex >= len(app.Clients) {
						goto END
					}
					client := app.Clients[clientIndex]
					client.GetNetStats()
					goto END
				}

			case commands.PROXYSTATS:
				if len(commandArgs) < 1 {
					c.GetProxyStats()
				} else {
					clientIndex, _ := strconv.Atoi(commandArgs[0])
					if clientIndex >= len(app.Clients) {
						goto END
					}
					client := app.Clients[clientIndex]
					client.GetProxyStats()
					goto END
				}

			case commands.PACKETSEARCH:
				if len(commandArgs) != 1 {
					goto END
				}
				// searches for a hex string in the packet log
				client := app.Clients[0]
				client.PacketSearch(&commandArgs[0])
				goto END

			case commands.DND:
				if len(commandArgs) != 1 {
					c.SetDND(nil)
					goto END
				}
				for _, direction := range common.GetAllCornerDirections() {
					// check if command arg is a valid direction and if not, print error
					if direction == commandArgs[0] {
						c.SetDND(&direction)
						break
					}
				}
			case commands.TEST: // ! test (DEBUG)
				c.Test()
			}
		}
	END:

		command.RunCooldown()

		if doLog {
			basicCommandStr := "**Who**: `%s`\n" +
				"**Command**: `%s`\n" +
				"**Args**: `%s`\n" +
				"**Session ID**: `%s`"
			who := "N/A"
			var args string
			if len(commandArgs) > 0 {
				args = strings.Join(commandArgs, " ")
			} else {
				args = "N/A"
			}
			if app.AuthAccount != nil {
				who = *app.AuthAccount.Name
			}
			logg := fmt.Sprintf(basicCommandStr, who, command.Name, args, strconv.FormatInt(int64(app.SessionID), 16))
			go auth.LogWebhookCommand(&logg)
		}
	}
}
