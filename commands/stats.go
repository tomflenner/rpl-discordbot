// Common stuff between stats_discord and stats_steam
package commands

import (
	"log"

	"github.com/b4cktr4ck5r3/rpl-discordbot/http"
	"github.com/b4cktr4ck5r3/rpl-discordbot/models"
	"github.com/b4cktr4ck5r3/rpl-discordbot/utils"
	"github.com/bwmarrin/discordgo"
)

// Messages
var (
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

func getPlayerStatsAndSummaries(i *discordgo.InteractionCreate, s *discordgo.Session, steam64 int64, player models.Player) (models.StatsApiPlayerResponse, models.SteamAccountSummaries) {
	playerStats, err := http.GetPlayerStats(player.SteamID)

	if err != nil {
		log.Println("1 - Erreur lors de getPlayerStatsAndSummaries: ", err.Error())
		s.InteractionRespond(i.Interaction, playerStatRequestErrorMsg)
	}

	playerSummaries, err := http.GetPlayerSteamAccountSummaries(steam64)

	if err != nil {
		log.Println("2 - Erreur lors de getPlayerStatsAndSummaries: ", err.Error())
		s.InteractionRespond(i.Interaction, playerSummariesRequestErrorMsg)
	}

	return playerStats, playerSummaries
}

func createAndSendStatsEmbed(i *discordgo.InteractionCreate, s *discordgo.Session, steam64 int64, player models.Player) {
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
}
