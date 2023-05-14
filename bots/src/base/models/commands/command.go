package commands

import "time"

type CommandName string

const (
	CONNECT           CommandName = "co"
	JOIN              CommandName = "jo"
	ENTER             CommandName = "en"
	REJOIN            CommandName = "rj"
	DISCONNECT_MULTI  CommandName = "dc"
	DISCONNECT_SINGLE CommandName = "d"
	FEED              CommandName = "fe"
	SETPLASMAFOR      CommandName = "splasmafor"
	EXIT              CommandName = "ex"
	CHAT              CommandName = "ch"
	EMOTE             CommandName = "em"
	RENAME            CommandName = "rn"
	SETEMOTELOOP      CommandName = "semloop"
	SETMASSTHRESHOLD  CommandName = "smassthreshold"
	SETAUTOREJOIN     CommandName = "sautorejoin"
	SETLVLMETA        CommandName = "slvlmeta"
	SETPLASMAFARM     CommandName = "splasmafarm"
	CLEAR             CommandName = "clear"
	INJECTTOKEN       CommandName = "injecttoken"
	GETTOKEN          CommandName = "gettoken"
	SPLASMAEXPLOIT    CommandName = "splasmaexploit"
	NETSTATS          CommandName = "netstats"
	PROXYSTATS        CommandName = "proxystats"
	PACKETSEARCH      CommandName = "packetsearch"
	TEST              CommandName = "test"
	DND               CommandName = "dnd"
)

type Command struct {
	Name               CommandName
	Description        string
	cooldownS          uint8
	RemainingCooldownS uint8
	PaidOnly           bool
	Deprecated         bool
}

func NewCommand(name CommandName, cooldownS uint8, paidOnly bool, deprecated bool) *Command {
	return &Command{
		Name:               name,
		cooldownS:          cooldownS,
		RemainingCooldownS: 0,
		PaidOnly:           paidOnly,
		Deprecated:         deprecated,
	}
}

func (c *Command) RunCooldown() {
	c.RemainingCooldownS = c.cooldownS
	go func() {
		for c.RemainingCooldownS > 0 {
			time.Sleep(time.Second)
			c.RemainingCooldownS--
		}
	}()
}

func (c *Command) IsOnCooldown() bool {
	return c.RemainingCooldownS > 0
}
