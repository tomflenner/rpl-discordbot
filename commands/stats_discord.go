package commands

import (
	"log"

	"github.com/MrWaggel/gosteamconv"
	"github.com/b4cktr4ck5r3/rpl-discordbot/config"
	"github.com/b4cktr4ck5r3/rpl-discordbot/database"
	"github.com/b4cktr4ck5r3/rpl-discordbot/utils"
	"github.com/bwmarrin/discordgo"
)

// Command definition
const StatsCommandName = "stats"
const StatsDiscordCommandName = "stats-discord"

var (
	StatsCommand = &discordgo.ApplicationCommand{
		Name:        StatsCommandName,
		Description: "Récupérer vos stats.",
	}

	StatsDiscordCommand = &discordgo.ApplicationCommand{
		Name:        StatsDiscordCommandName,
		Description: "Récupérer les stats d'un joueur à partir de son compte discord.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "user",
				Description: "@ du joueur à rechercher",
				Type:        discordgo.ApplicationCommandOptionUser,
				Required:    true,
			},
		},
	}
)

// Messages
var (
	playerNotFoundWithDiscordIdErrorMsg = &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Le joueur n'a pas lié son compte Discord et son compte Steam.",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	}
)

func StatsDiscordCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if canExecuteRestrictedCommand(i, config.Cfg.StatsChannelID) {
		var discordId string

		options := i.ApplicationCommandData().Options

		if len(options) == 0 {
			discordId = i.Member.User.ID
		} else {
			discordId = options[0].UserValue(s).ID
		}

		player, err := database.SelectPlayerByDiscordId(discordId)

		if err != nil {
			log.Println("1 - Erreur lors du stats-discord: ", err.Error())
			s.InteractionRespond(i.Interaction, playerNotFoundWithDiscordIdErrorMsg)
		}

		steam64, err := gosteamconv.SteamStringToInt64(player.SteamID)

		if err != nil {
			log.Println("2 - Erreur lors de la conversion du SteamID en Steam64: ", err.Error())
			s.InteractionRespond(i.Interaction, errorMsg)
		}

		playerStats, playerSummaries := getPlayerStatsAndSummaries(i, s, steam64, player)
		embed := utils.CreateStatsEmbed(playerStats, player, playerSummaries)

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					embed,
				},
				Flags: discordgo.MessageFlagsEphemeral,
			},
		})
	} else {
		s.InteractionRespond(i.Interaction, notAuthorizedMsg)
	}
}
