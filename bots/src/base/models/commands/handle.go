package commands

type CommandHandler struct {
	Commands []*Command
}

func NewCommandHandler() *CommandHandler {
	commands := []*Command{
		NewCommand(CONNECT, 60, false, false),
		NewCommand(JOIN, 10, false, false),
		NewCommand(ENTER, 3, false, false),
		NewCommand(REJOIN, 3, false, false),
		NewCommand(DISCONNECT_MULTI, 3, false, false),
		NewCommand(DISCONNECT_SINGLE, 3, false, false),
		NewCommand(FEED, 3, false, false),
		NewCommand(SETPLASMAFOR, 3, false, false),
		NewCommand(EXIT, 3, false, false),
		NewCommand(CHAT, 3, false, false),
		NewCommand(EMOTE, 3, false, false),
		NewCommand(RENAME, 3, false, false),
		NewCommand(SETEMOTELOOP, 3, false, false),
		NewCommand(SETMASSTHRESHOLD, 3, false, false),
		NewCommand(SETAUTOREJOIN, 3, false, false),
		NewCommand(SETLVLMETA, 3, false, false),
		NewCommand(SETPLASMAFARM, 3, false, false),
		NewCommand(CLEAR, 3, false, false),
		NewCommand(INJECTTOKEN, 3, false, false),
		NewCommand(GETTOKEN, 3, false, false),
		NewCommand(SPLASMAEXPLOIT, 3, true, true),
		NewCommand(NETSTATS, 3, false, false),
		NewCommand(PROXYSTATS, 3, false, false),
		NewCommand(PACKETSEARCH, 3, false, false),
		NewCommand(TEST, 3, false, false),
		NewCommand(DND, 3, false, false),
	}
	return &CommandHandler{
		Commands: commands,
	}
}

func (c *CommandHandler) GetCommand(name CommandName) *Command {
	for _, command := range c.Commands {
		if command.Name == name {
			return command
		}
	}
	return nil
}
