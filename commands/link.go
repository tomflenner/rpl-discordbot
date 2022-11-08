package commands

import (
	"database/sql"

	"github.com/b4cktr4ck5r3/rpl-discordbot/database"
	"github.com/bwmarrin/discordgo"
)

// Command definition
const LinkCommandName = "link"

var (
	LinkCommand = &discordgo.ApplicationCommand{
		Name:        LinkCommandName,
		Description: "Lier son compte Discord à son compte Steam",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "activation-code",
				Description: "Code d'activation fournit par le serveur CSGO",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
		},
	}
)

// Messages
var (
	playerNotFoundErrorMsg = &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Code d'activation introuvable.",
		},
	}

	playerAlreadyRegisteredErrorMsg = &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Le compte Steam est déjà associé à un compte Discord.",
		},
	}

	playerRegisteredErrorMsg = &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Une erreur s'est produite lors de la liaison, veuillez contacter un administrateur.",
		},
	}

	playerRegisteredSuccessMsg = &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Compte lié avec succès !",
		},
	}
)

func LinkCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	player, err := database.SelectPlayerByLinkCode(options[0].StringValue())

	if err != nil {
		s.InteractionRespond(i.Interaction, playerNotFoundErrorMsg)
	}

	if player.DiscordID.Valid {
		s.InteractionRespond(i.Interaction, playerAlreadyRegisteredErrorMsg)
	}

	player.DiscordID = sql.NullString{
		String: i.Interaction.Member.User.ID,
		Valid:  true,
	}

	ok, err := database.UpdatePlayer(player)

	if err != nil || !ok {
		s.InteractionRespond(i.Interaction, playerRegisteredErrorMsg)
	}

	s.InteractionRespond(i.Interaction, playerRegisteredSuccessMsg)

}
