package command

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

func (c *Command) bindingCommand() {
	c.Commands["bind"] = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options
		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}
		response := "绑定玩家"
		if option, ok := optionMap["uid"]; ok {
			response += option.StringValue()
		}
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: response,
			},
		})
		if err != nil {
			logger.Error("绑定玩家错误:", err)
			return
		}
	}
	// c.bindingApplication()
}

func (c *Command) bindingApplication() {
	cmd, err := c.ApplicationCommandCreate(c.State.User.ID, "", &discordgo.ApplicationCommand{
		Name:        "bind",
		Description: "绑定玩家",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "uid",
				Required:    true,
				Description: "你游戏里面的uid",
			},
		},
	})
	if err != nil {

	}
	c.Cmds = append(c.Cmds, cmd)
}
