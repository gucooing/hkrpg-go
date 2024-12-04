package discord

import (
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/suppl/discord/command"
)

var Once sync.Once

type Bot struct {
	*discordgo.Session
}

func Start() {
	logger.Info("DiscordBot Start")
	err := loadConfig("discord.json")
	if err != nil {
		logger.Error("Load Discord config error: %s", err)
		return
	}
	b := new(Bot)
	cfg := getConfig()
	database.NewDiscord(cfg.Db.Type, cfg.Db.Dns)
	b.Session, err = discordgo.New(cfg.Type + " " + cfg.Token)
	if err != nil {
		logger.Error("Discord Bot error: %s", err)
		return
	}
	b.LogLevel = discordgo.LogDebug
	c := command.InitCommands(b.Session)
	b.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := c.Commands[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
	err = b.Open()
	if err != nil {
		logger.Error("Open Discord Bot error: %s", err)
		return
	}
	c.Application()

	select {}
}
