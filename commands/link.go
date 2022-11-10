package commands

import (
	"database/sql"
	"log"

	"github.com/b4cktr4ck5r3/rpl-discordbot/config"
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
	playerNotFoundWithLinkCodeErrorMsg = &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Code d'activation introuvable.",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	}

	playerAlreadyRegisteredErrorMsg = &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Le compte Steam est déjà associé à un compte Discord.",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	}

	playerRegisteredErrorMsg = &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Une erreur s'est produite lors de la liaison, veuillez contacter un administrateur.",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	}

	playerRegisteredSuccessMsg = &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Compte lié avec succès !",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	}
)

func LinkCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if canExecuteRestrictedCommand(i, config.Cfg.LinkChannelID) {
		options := i.ApplicationCommandData().Options

		player, err := database.SelectPlayerByLinkCode(options[0].StringValue())

		if err != nil {
			log.Println("1 - Erreur du link: ", err.Error())
			s.InteractionRespond(i.Interaction, playerNotFoundWithLinkCodeErrorMsg)
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
			log.Println("2 - Erreur du link: ", err.Error())
			s.InteractionRespond(i.Interaction, playerRegisteredErrorMsg)
		}

		s.InteractionRespond(i.Interaction, playerRegisteredSuccessMsg)
	} else {
		s.InteractionRespond(i.Interaction, notAuthorizedMsg)
	}
}
