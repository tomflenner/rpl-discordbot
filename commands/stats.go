package commands

import (
	"fmt"
	"strconv"
	"time"

	"github.com/b4cktr4ck5r3/rpl-discordbot/config"
	"github.com/b4cktr4ck5r3/rpl-discordbot/database"
	"github.com/b4cktr4ck5r3/rpl-discordbot/http"
	"github.com/b4cktr4ck5r3/rpl-discordbot/models"
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

	playerStatRequestErrorMsg = &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Une erreur s'est produite lors de la récupération des statistiques, veuillez contacter un administrateur.",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	}

	playerSummariesRequestErrorMsg = &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Une erreur s'est produite lors de la récupération des informations steam liées à l'utilisateur, veuillez contacter un administrateur.",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	}
)

func StatsCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if canExecuteRestrictedCommand(i, config.Cfg.StatsChannelID) {
		var discordId string

		options := i.ApplicationCommandData().Options

		if len(options) > 0 {
			discordId = options[0].UserValue(s).ID
		} else {
			discordId = i.Member.User.ID
		}

		player, err := database.SelectPlayerByDiscordId(discordId)

		if err != nil {
			s.InteractionRespond(i.Interaction, playerNotFoundWithDiscordIdErrorMsg)
		}

		playerStats, err := http.GetPlayerStats(player.SteamID)

		if err != nil {
			s.InteractionRespond(i.Interaction, playerStatRequestErrorMsg)
		}

		playerSummaries, err := http.GetPlayerSteamAccountSummaries(player.SteamID)

		if err != nil {
			s.InteractionRespond(i.Interaction, playerSummariesRequestErrorMsg)
		}

		var color int
		var ratio string

		if playerStats.Ratio > 1 {
			color = 0x00FF00
			ratio = "📈 Ratio"
		} else {
			color = 0xFF0000
			ratio = "📉 Ratio"
		}

		currentTime := time.Now()
		embed := models.NewEmbed().
			SetTitle(fmt.Sprintf("📊 Statistiques du joueur %s", player.Name)).
			SetDescription("Données officielles du classement des serveurs Retake Pro League.").
			AddField("🏆 Rang", strconv.Itoa(int(playerStats.Rank))).
			AddField("🔫 Kills", strconv.Itoa(int(playerStats.Kills))).
			AddField("💀 Morts", strconv.Itoa(int(playerStats.Kills))).
			AddField(ratio, fmt.Sprintf("%.2f", playerStats.Ratio)).
			AddField("🤯 Headshots", strconv.Itoa(int(playerStats.Headshots))).
			AddField("💥 Headshots %", strconv.Itoa(int(playerStats.HeadshotsPercent))).
			InlineAllFields().
			SetFooter(fmt.Sprintf("Généré le %s", currentTime.Local().Format("02-Jan-2006 15:04:05"))).
			SetThumbnail(playerSummaries.Avatarmedium).
			SetColor(color).MessageEmbed

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
		s.InteractionRespond(i.Interaction, notAuthorized)
	}
}
