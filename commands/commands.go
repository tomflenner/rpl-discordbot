package commands

import "github.com/bwmarrin/discordgo"

var (
	Commands = []*discordgo.ApplicationCommand{
		{
			Name:        "test-command",
			Description: "Test command",
		},
	}

	CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"test-command": testCommand,
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
