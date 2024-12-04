package suppl

import (
	"flag"

	"github.com/gucooing/hkrpg-go/suppl/discord"
)

func Start() {
	discordFlag := flag.Bool("discord", false, "Start Discord Bot")
	flag.Parse()
	if *discordFlag {
		discord.Once.Do(func() {
			go discord.Start()
		})
	}
	// go pushc.NewPushClient()
}
