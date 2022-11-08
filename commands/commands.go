package commands

import (
	"github.com/b4cktr4ck5r3/rpl-discordbot/models"
	"github.com/bwmarrin/discordgo"
)

const (
	testCommandName      string = "test-command"
	testEmbedCommandName string = "test-embed-command"
)

var (
	Commands = []*discordgo.ApplicationCommand{
		{
			Name:        testCommandName,
			Description: "Test command",
		},
		{
			Name:        testEmbedCommandName,
			Description: "Test embed command",
		},
		LinkCommand,
	}

	CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		testCommandName:      testCommand,
		testEmbedCommandName: testEmbedCommand,
		LinkCommandName:      LinkCommandHandler,
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
