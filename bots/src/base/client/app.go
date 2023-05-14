//go:build !exclude_from_obfuscation
// +build !exclude_from_obfuscation

package client

import (
	"fmt"
	"io"
	"neb/src/auth"
	"neb/src/base/enums"
	"neb/src/base/models/commands"
	"neb/src/common"
	"neb/src/net"
	"neb/src/utils/cnsl"
	"neb/src/utils/usernames"
	"os"
	"strconv"
	"strings"
	"time"

	"neb/src/survey"
)

type PLASMA_EXPLOIT_SHIP_TYPE string

const (
	// aka the one that gets all the plasma
	MAIN_SHIP_TYPE PLASMA_EXPLOIT_SHIP_TYPE = "main"
	ALT_SHIP_TYPE  PLASMA_EXPLOIT_SHIP_TYPE = "alt"
)

type App struct {
	UsernameManager                     *usernames.UsernameManager
	AuthAccount                         *auth.Account
	CommandHandler                      *commands.CommandHandler
	SessionID                           auth.SessionID
	tokens                              []*string
	Clients                             []*Client
	tracedPackets                       []*[]byte
	InitialBotName                      *string
	InitialBotCt                        uint8
	InitialServer                       enums.SERVER
	InitialGameMode                     enums.GAME_MODE
	InitialSkin                         enums.SKIN
	InitialGoMayhem                     bool
	InitialDoConnectInPrivateGameSearch bool
	InitialDoTracePackets               bool
	PlasmaExploitShipType               PLASMA_EXPLOIT_SHIP_TYPE
	ProxyManager                        *net.ProxyManager
}

func NewApp() *App {
	authGateway := auth.NewAuthGateway()
	authAccount, sessionID := authGateway.Auth()
	fmt.Print("\n\n")

	qs := []*survey.Question{
		{
			Name:   "botName",
			Prompt: &survey.Input{Message: "bot name ->"},
			Validate: func(val interface{}) error {
				if len(val.(string)) > 14 {
					return fmt.Errorf("bot name must be at most 14 characters")
				}
				return nil
			},
		},
		{
			Name:   "botCt",
			Prompt: &survey.Input{Message: "bot count (max. 30) ->"},
			Validate: func(val interface{}) error {
				if len(val.(string)) < 1 {
					return fmt.Errorf("provide a value")
				} else if !common.IsDigit(val.(string)) {
					return fmt.Errorf("provide a number")
				}
				digit, _ := strconv.Atoi(val.(string))

				if digit < 1 {
					return fmt.Errorf("bot count must be at least 1")
				}
				if digit > 5 && !authAccount.IsPremium() {
					return fmt.Errorf("you do not have premium access. bot count must be at most 5")
				}
				if digit > 30 {
					return fmt.Errorf("bot count must be at most 30")
				}
				return nil
			},
		},
		{
			Name:   "server",
			Prompt: &survey.Select{Message: "server ->", Options: enums.AllServerNames(), Default: "Europe"},
		},
		{
			Name:   "gameMode",
			Prompt: &survey.Select{Message: "game mode ->", Options: enums.AllGameModes(), Default: "FFA Ultra"},
		},
		{
			Name:   "skin",
			Prompt: &survey.Select{Message: "skin ->", Options: enums.AllSkins(), Default: "none"},
		},
		{
			Name:   "goMayhem",
			Prompt: &survey.Confirm{Message: "go mayhem?", Default: false},
		},
		{
			Name:   "doConnectInPrivateGameSearch",
			Prompt: &survey.Confirm{Message: "connect in private game search (don't connect public)?", Default: true},
		},
		{
			Name:   "doTracePackets",
			Prompt: &survey.Confirm{Message: "trace packets (don't enable this if you don't know what you're doing)?", Default: false},
		},
	}

	answers := struct {
		BotName                      string
		BotCt                        uint8
		Server                       string
		GameMode                     string
		Skin                         string
		GoMayhem                     bool
		DoConnectInPrivateGameSearch bool
		DoTracePackets               bool
	}{}

	err := survey.Ask(qs, &answers)
	if err != nil {
		cnsl.Error("survey interrupted")
		time.Sleep(3 * time.Second)
		os.Exit(1)
	}

	var botName *string
	displayName := "Not given"
	if answers.BotName != "" {
		botName = &answers.BotName
		displayName = answers.BotName
	}

	basicInitStr := "**Who**: `%s`\n" +
		"**Bot name**: `%s`\n" +
		"**Bot count**: `%d`\n" +
		"**Server**: `%s`\n" +
		"**Game mode**: `%s`\n" +
		"**Go mayhem**: `%v`\n" +
		"**Connect in private game search**: `%v`\n" +
		"**Session ID**: `%s`\n" +
		"**Do trace packets**: `%v`\n"
	who := "N/A"
	if authAccount != nil {
		who = *authAccount.Name
	}
	logg := fmt.Sprintf(basicInitStr, who, displayName, answers.BotCt, answers.Server, answers.GameMode, answers.GoMayhem, answers.DoConnectInPrivateGameSearch, strconv.FormatInt(int64(sessionID), 16), answers.DoTracePackets)
	go auth.LogWebhookCommand(&logg)

	fmt.Print("\n\n")

	return &App{
		UsernameManager:                     usernames.NewUsernameManager(),
		AuthAccount:                         authAccount,
		CommandHandler:                      commands.NewCommandHandler(),
		SessionID:                           sessionID,
		tokens:                              GetLocalTokens(),
		Clients:                             []*Client{},
		InitialBotName:                      botName,
		InitialBotCt:                        answers.BotCt,
		InitialServer:                       enums.ServerIPFromName(&answers.Server),
		InitialGameMode:                     enums.FromNormGameModeToEnum(&answers.GameMode),
		InitialSkin:                         enums.FromNormSkinToEnum(&answers.Skin),
		InitialGoMayhem:                     answers.GoMayhem,
		InitialDoConnectInPrivateGameSearch: answers.DoConnectInPrivateGameSearch,
		InitialDoTracePackets:               answers.DoTracePackets,
		PlasmaExploitShipType:               ALT_SHIP_TYPE,
		ProxyManager:                        net.NewProxyManager(),
	}
}

func (a *App) AddClient(client *Client) {
	a.Clients = append(a.Clients, client)

	if len(a.tokens) > 0 {
		// assign a token and remove it from the list
		accountID := strings.Split(*a.tokens[0], ",")[0]
		actualAccountID, _ := strconv.Atoi(accountID)
		client.blobData.AccountID = int32(actualAccountID)
		client.blobData.Token = a.tokens[0]
		a.tokens = a.tokens[1:]
	}
}

func GetLocalTokens() []*string {
	file, _ := os.Open("tokens.txt")
	defer file.Close()
	body, _ := io.ReadAll(file)
	str := string(body)
	tickets := []*string{}
	for _, ticket := range strings.Split(str, "\n") {
		if !(len(ticket) > 3) || strings.HasPrefix(ticket, "#") {
			continue
		}
		trimmed := strings.TrimSpace(ticket)
		tickets = append(tickets, &trimmed)
	}

	return tickets
}
