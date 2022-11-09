package commands

import (
	"github.com/b4cktr4ck5r3/rpl-discordbot/config"
	"github.com/bwmarrin/discordgo"
)

var (
	Commands = []*discordgo.ApplicationCommand{
		LinkCommand,
		StatsCommand,
		StatsDiscordCommand,
	}

	CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		LinkCommandName:         LinkCommandHandler,
		StatsCommandName:        StatsCommandHandler,
		StatsDiscordCommandName: StatsCommandHandler,
	}
)

var (
	notAuthorized = &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Cette commande n'est pas autorisée ici.",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	}
)

func canExecuteRestrictedCommand(i *discordgo.InteractionCreate, channelId string) bool {
	return i.GuildID == config.Cfg.GuildID && i.ChannelID == channelId
}
