package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/b4cktr4ck5r3/rpl-discordbot/models"
	"github.com/bwmarrin/discordgo"
)

func CreateStatsEmbed(playerStats models.StatsApiPlayerResponse, player models.Player, playerSummaries models.SteamAccountSummaries) *discordgo.MessageEmbed {
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
	return embed
}
