package auth

import (
	"fmt"
	"io"
	"neb/src/common"
	"neb/src/utils/cnsl"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gtuk/discordwebhook"
)

type AccountRank string

const (
	RANK_PREMIUM AccountRank = "PREMIUM"
	RANK_ADMIN   AccountRank = "ADMIN"
	RANK_DEV     AccountRank = "DEV"
	RANK_BANNED  AccountRank = "BANNED"
	RANK_USER    AccountRank = "USER"
)

const (
	PASTEBIN_AUTH_DATA = "THIS DATA IS NOT AVAILABLE TO THE PUBLIC. CREATE YOUR OWN IMPLEMENTATION."
)

type SessionID int64

type AuthGateway struct {
	Accounts []*Account
	Flags    *Flags
}

type Flags struct {
	IsGlobalAccess bool
	Version        int
}

type Account struct {
	Name             *string
	DiscordID        *string
	Ranks            []AccountRank
	ExpiryTS         int64
	Sig              *string
	PunishmentReason *string
	// based off of the local time
	// if the user is evading expiry by changing their system time,
	// we can detect it with other logins.
	SessionID SessionID
}

func (a *Account) IsPremium() bool {
	for _, rank := range a.Ranks {
		if rank == RANK_PREMIUM {
			return true
		}
	}
	return false
}

func (a *Account) IsBanned() bool {
	for _, rank := range a.Ranks {
		if rank == RANK_BANNED {
			return true
		}
	}
	return false
}

func (a *Account) IsAdmin() bool {
	for _, rank := range a.Ranks {
		if rank == RANK_ADMIN {
			return true
		}
	}
	return false
}

func (a *Account) IsDev() bool {
	for _, rank := range a.Ranks {
		if rank == RANK_DEV {
			return true
		}
	}
	return false
}

func (a *Account) IsLicenseExpired() bool {
	return a.ExpiryTS < time.Now().UTC().Unix()
}

func (a *Account) ranksToString() []string {
	var ranks []string
	for _, rank := range a.Ranks {
		ranks = append(ranks, string(rank))
	}
	return ranks
}

func NewAuthGateway() *AuthGateway {
	client := new(http.Client)
	req, err := http.NewRequest("GET", PASTEBIN_AUTH_DATA, nil)
	if err != nil {
		cnsl.Error("failed to create request to pastebin.")
		time.Sleep(3 * time.Second)
		os.Exit(1)
	}
	res, err := client.Do(req)
	if err != nil {
		cnsl.Error("failed to send request to pastebin. make sure your internet is working and try again. if it still does not work, use vpn.")
		time.Sleep(3 * time.Second)
		os.Exit(1)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		cnsl.Error("pastebin auth data error. please contact the bot developer.")
		time.Sleep(3 * time.Second)
		os.Exit(1)
	}
	str := string(body)
	decrypted, err := common.Decrypt(&str)

	if err != nil {
		cnsl.Error("could not decrypt auth data. please contact the bot developer.")
		time.Sleep(3 * time.Second)
		os.Exit(1)
	}

	accounts, flags := parseAccounts(*decrypted)
	return &AuthGateway{Accounts: accounts, Flags: flags}
}

func (ag *AuthGateway) Auth() (*Account, SessionID) {
	sessionID := SessionID(time.Now().UTC().Unix())
	sessionIDAsHex := strconv.FormatInt(int64(sessionID), 16)
	basicLoginStr :=
		"**Who**: `%s`\n" +
			"**IPv4**: ||`%s`||\n" +
			"**Global Access Flag**: `%v`\n" +
			"**SIG**: ||`%s`||\n" +
			"**Discord ID**: `%s`\n" +
			"**Ranks**: `%s`\n" +
			"**License Expired**: `%v`\n" +
			"**Premium**: `%v`\n" +
			"**Reason**: `%s`\n" +
			"**Session ID**: `%s`"
	var maybeAccount *Account
	who := "N/A"
	discordID := "N/A"
	ranks := "N/A"
	licenseExpired := false
	premium := false
	sig := getLocalSigFromKeyFile("auth.key")

	if sig == nil {
		logg := fmt.Sprintf(
			basicLoginStr,
			who,
			*common.GetIP(),
			ag.Flags.IsGlobalAccess,
			"N/A",
			discordID,
			ranks,
			licenseExpired,
			premium,
			"No license key file found.",
			sessionIDAsHex,
		)
		go LogWebhookLogin(&logg, "BAD")
		cnsl.Error("unauthorized. did you delete the auth.key file? please redownload.")
		time.Sleep(3 * time.Second)
		os.Exit(1)
	}

	for _, account := range ag.Accounts {
		if *account.Sig == *sig {
			maybeAccount = account
			maybeAccount.SessionID = sessionID
		}
	}

	if maybeAccount != nil {
		who = *maybeAccount.Name
		discordID = *maybeAccount.DiscordID
		ranks = strings.Join(maybeAccount.ranksToString(), ", ")
		licenseExpired = maybeAccount.IsLicenseExpired()
		premium = maybeAccount.IsPremium()
	}

	if ag.Flags.Version != common.VERSION {
		logg := fmt.Sprintf(
			basicLoginStr,
			who,
			*common.GetIP(),
			ag.Flags.IsGlobalAccess,
			*sig,
			discordID,
			ranks,
			licenseExpired,
			premium,
			"Version mismatch.",
			sessionIDAsHex,
		)
		go LogWebhookLogin(&logg, "BAD")
		cnsl.Error("unauthorized. your software version is outdated. please redownload.")
		time.Sleep(3 * time.Second)
		os.Exit(1)
	}

	if maybeAccount != nil {
		if maybeAccount.IsLicenseExpired() {
			logg := fmt.Sprintf(
				basicLoginStr,
				who,
				*common.GetIP(),
				ag.Flags.IsGlobalAccess,
				*sig,
				discordID,
				strings.Join(maybeAccount.ranksToString(), ", "),
				maybeAccount.IsLicenseExpired(),
				premium,
				"License key expired.",
				sessionIDAsHex,
			)
			go LogWebhookLogin(&logg, "BAD")
			cnsl.Error("unauthorized. your license key has expired. please get a new one by contacting the bot developer.")
			time.Sleep(3 * time.Second)
			os.Exit(1)
		}
		if maybeAccount.IsBanned() {
			logg := fmt.Sprintf(
				basicLoginStr,
				who,
				*common.GetIP(),
				ag.Flags.IsGlobalAccess,
				*sig,
				discordID,
				ranks,
				licenseExpired,
				premium,
				fmt.Sprintf("Banned: %s", *maybeAccount.PunishmentReason),
				sessionIDAsHex,
			)
			go LogWebhookLogin(&logg, "BAD")
			cnsl.Error(fmt.Sprintf("unauthorized. you were permanently banned for -> %s", *maybeAccount.PunishmentReason))
			time.Sleep(3 * time.Second)
			os.Exit(1)
		}

		logg := fmt.Sprintf(
			basicLoginStr,
			who,
			*common.GetIP(),
			ag.Flags.IsGlobalAccess,
			*sig,
			discordID,
			ranks,
			licenseExpired,
			premium,
			"License key valid.",
			sessionIDAsHex,
		)
		go LogWebhookLogin(&logg, "GOOD")
		cnsl.Ok(fmt.Sprintf("authorized. welcome back, %s! premium access -> %v", *maybeAccount.Name, maybeAccount.IsPremium()))
		return maybeAccount, sessionID
	}

	if ag.Flags.IsGlobalAccess {
		logg := fmt.Sprintf(
			basicLoginStr,
			who,
			*common.GetIP(),
			ag.Flags.IsGlobalAccess,
			*sig,
			discordID,
			ranks,
			licenseExpired,
			premium,
			"Global access.",
			sessionIDAsHex,
		)
		go LogWebhookLogin(&logg, "GOOD")
		cnsl.Ok(fmt.Sprintf("authorized (global access). welcome back, %s!", who))
		return maybeAccount, sessionID
	}

	logg := fmt.Sprintf(
		basicLoginStr,
		"N/A",
		*common.GetIP(),
		ag.Flags.IsGlobalAccess,
		*sig,
		"N/A",
		"N/A",
		false,
		false,
		"No license key.",
		sessionIDAsHex,
	)
	go LogWebhookLogin(&logg, "BAD")
	cnsl.Error("unauthorized. you do not have a license key. please contact the bot developer to get one.")
	time.Sleep(3 * time.Second)
	os.Exit(1)
	return maybeAccount, sessionID
}

func getLocalSigFromKeyFile(fileName string) *string {
	file, err := os.Open(fileName)
	if err != nil {
		return nil
	}
	defer file.Close()

	body, _ := io.ReadAll(file)
	str := string(body)
	sig := strings.Split(str, " ")[1]
	return &sig
}

func parseAccounts(str string) ([]*Account, *Flags) {
	accounts := make([]*Account, 0)
	lines := strings.Split(str, "\n")

	flagsSegStart := false
	accountsSegStart := false
	flags := new(Flags)
	var account *Account

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "FLAGS") {
			flagsSegStart = true
			continue
		}

		if strings.HasPrefix(line, "ACCOUNTS") {
			accountsSegStart = true
			continue
		}

		if flagsSegStart {
			if strings.HasPrefix(line, "GLOBAL_ACCESS") {
				flags.IsGlobalAccess = strings.Split(line, " = ")[1] == "1"
			} else if strings.HasPrefix(line, "VERSION") {
				version := strings.Split(line, " = ")[1]
				actualVersion, _ := strconv.Atoi(version)
				flags.Version = actualVersion
			}
		}

		if accountsSegStart {
			if strings.HasPrefix(line, "NAME") {
				account = &Account{Name: &strings.Split(line, ": ")[1]}
			} else if strings.HasPrefix(line, "DISCORD_ID") {
				account.DiscordID = &strings.Split(line, ": ")[1]
			} else if strings.HasPrefix(line, "RANKS") {
				ranks := strings.Split(strings.Split(line, ": ")[1], ", ")
				for _, rank := range ranks {
					account.Ranks = append(account.Ranks, AccountRank(rank))
				}
			} else if strings.HasPrefix(line, "EXPIRY_TS") {
				expiryTS := strings.Split(line, ": ")[1]
				actualExpiryTS, _ := strconv.ParseInt(expiryTS, 10, 64)
				account.ExpiryTS = actualExpiryTS
			} else if strings.HasPrefix(line, "PUNISHMENT_REASON") {
				reason := strings.Split(line, ": ")[1]
				account.PunishmentReason = &reason
			} else if strings.HasPrefix(line, "SIG") {
				sig := strings.Split(line, ": ")[1]
				account.Sig = &sig
				accounts = append(accounts, account)
			}
		}
	}

	return accounts, flags
}

func LogWebhookLogin(content *string, level string) {
	var color uint32
	if level == "BAD" {
		color = 0xE41B17
	} else if level == "GOOD" {
		color = 0x00DEAD
	}
	colorDecStr := fmt.Sprintf("%d", color)
	title := "Login"
	embeds := []discordwebhook.Embed{{
		Color:       &colorDecStr,
		Title:       &title,
		Description: content,
	}}
	wh := discordwebhook.Message{Embeds: &embeds}
	err := discordwebhook.SendMessage(
		"https://canary.discord.com/api/webhooks/1093146591050870824/GOgl-Y8x3dN_5vFISBLyYN38MQy3KY5fHwioEEnYQTsu7wq317D3k524l40woabBWYT3",
		wh,
	)
	if err != nil {
		cnsl.Error("failed to send request. make sure your internet is working and try again. if it still does not work, use vpn.")
		time.Sleep(3 * time.Second)
		os.Exit(1)
	}
}

func LogWebhookCommand(content *string) {
	color := 0x2b2d31 // invisible embed
	colorDecStr := fmt.Sprintf("%d", color)
	title := "Command Used"
	embeds := []discordwebhook.Embed{{
		Color:       &colorDecStr,
		Title:       &title,
		Description: content,
	}}
	wh := discordwebhook.Message{Embeds: &embeds}
	err := discordwebhook.SendMessage(
		"https://canary.discord.com/api/webhooks/1093146807791517756/W0EmfqSibV5dQCxGwMZiAdFh0gHK9K-Gz0pZfJ3bW3sNUSK17oHKOfCGSARUqddmziM4",
		wh,
	)
	if err != nil {
		cnsl.Error("failed to send request. make sure your internet is working and try again. if it still does not work, use vpn.")
		time.Sleep(3 * time.Second)
		os.Exit(1)
	}
}
