package commands

import (
	"github.com/b4cktr4ck5r3/rpl-discordbot/models"
	"github.com/bwmarrin/discordgo"
)

var (
	Commands = []*discordgo.ApplicationCommand{
		{
			Name:        "test-command",
			Description: "Test command",
		},
		{
			Name:        "test-embed-command",
			Description: "Test embed command",
		},
	}

	CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"test-command":       testCommand,
		"test-embed-command": testEmbedCommand,
	}
)

func testCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Premier test slash commands",
		},
	})
}

func testEmbedCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	embed := models.NewEmbed().
		SetTitle("Title test embed command").
		SetDescription("Description test embed command").
		AddField("I am a field", "I am a value").
		AddField("I am a second field", "I am a value").
		SetImage("https://cdn.discordapp.com/avatars/119249192806776836/cc32c5c3ee602e1fe252f9f595f9010e.jpg?size=2048").
		SetThumbnail("https://cdn.discordapp.com/avatars/119249192806776836/cc32c5c3ee602e1fe252f9f595f9010e.jpg?size=2048").
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
