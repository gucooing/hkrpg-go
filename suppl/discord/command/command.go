package command

import (
	"github.com/bwmarrin/discordgo"
)

type CommandFunc func(s *discordgo.Session, i *discordgo.InteractionCreate)

type Command struct {
	*discordgo.Session
	Commands map[string]CommandFunc
	Cmds     []*discordgo.ApplicationCommand
}

func InitCommands(s *discordgo.Session) *Command {
	c := &Command{
		Session:  s,
		Commands: make(map[string]CommandFunc),
		Cmds:     make([]*discordgo.ApplicationCommand, 0),
	}
	c.infoCommand()
	c.bindingCommand()

	return c
}

func (c *Command) Application() {
	c.infoApplication()
	c.bindingApplication()
}
