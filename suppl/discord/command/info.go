package command

import (
	"github.com/bwmarrin/discordgo"
	hkrpg_go_pe "github.com/gucooing/hkrpg-go/hkrpg-go-pe"
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

func (c *Command) infoCommand() {
	c.Commands["info"] = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		response := hkrpg_go_pe.GetServerInfo()
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: response,
			},
		})
		if err != nil {
			logger.Error("获取服务器状态错误:", err)
			return
		}
	}
	// c.infoApplication()
}

func (c *Command) infoApplication() {
	cmd, err := c.ApplicationCommandCreate(c.State.User.ID, "", &discordgo.ApplicationCommand{
		Name:        "info",
		Description: "获取服务器状态",
	})
	if err != nil {

	}
	c.Cmds = append(c.Cmds, cmd)
}
