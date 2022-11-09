package commands

import (
	"github.com/b4cktr4ck5r3/rpl-discordbot/config"
	"github.com/b4cktr4ck5r3/rpl-discordbot/models"
	"github.com/bwmarrin/discordgo"
)

var (
	Commands = []*discordgo.ApplicationCommand{
		LinkCommand,
		StatsDiscordCommand,
	}

	CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		LinkCommandName:         LinkCommandHandler,
		StatsDiscordCommandName: StatsDiscordCommandHandler,
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

func testEmbedCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	embed := models.NewEmbed().
		SetTitle("Title test embed command").
		SetDescription("Description test embed command").
		AddField("I am a field", "I am a value").
		AddField("I am a second field", "I am a value").
		SetImage("https://picsum.photos/500").
		SetThumbnail("https://picsum.photos/100").
		SetColor(0x00ff00).MessageEmbed

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				embed,
			},
		},
	})
}
