package commands

import (
	"log"
	"strconv"
	"strings"

	"github.com/MrWaggel/gosteamconv"
	"github.com/b4cktr4ck5r3/rpl-discordbot/config"
	"github.com/b4cktr4ck5r3/rpl-discordbot/database"
	"github.com/b4cktr4ck5r3/rpl-discordbot/http"
	"github.com/b4cktr4ck5r3/rpl-discordbot/utils"
	"github.com/bwmarrin/discordgo"
)

// Command definition
const StatsSteamProfileCommandName = "stats-steam-profil"

var (
	StatsSteamStatsSteamProfile = &discordgo.ApplicationCommand{
		Name:        StatsSteamProfileCommandName,
		Description: "Récupérer les stats d'un joueur à partir d'une URL steam.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "url",
				Description: "Url steam du joueur à rechercher",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
		},
	}
)

// Messages
var (
	playerNotFoundWithSteamProfileErrorMsg = &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Ce profile steam n'est pas dans le classement.",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	}

	badUrlParameterErrorMsg = &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "L'URL que vous avez fournit n'est pas correcte.",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	}
)

func StatsSteamCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if canExecuteRestrictedCommand(i, config.Cfg.StatsChannelID) {
		options := i.ApplicationCommandData().Options

		var steam64 int64
		var err error

		url := options[0].StringValue()
		urlType := utils.GetUrlType(url)

		steam64 = getSteam64(urlType, s, i, url, steam64)

		steamId, err := gosteamconv.SteamInt64ToString(steam64)
		if err != nil {
			log.Println("Erreur lors de la conversion du Steam64 en SteamID")
			s.InteractionRespond(i.Interaction, errorMsg)
		}

		steamId = strings.ReplaceAll(steamId, "STEAM_0", "STEAM_1")

		player, err := database.SelectPlayerBySteamId(steamId)
		if err != nil {
			s.InteractionRespond(i.Interaction, playerNotFoundWithSteamProfileErrorMsg)
		}

		createAndSendStatsEmbed(i, s, steam64, player)
	} else {
		s.InteractionRespond(i.Interaction, notAuthorizedMsg)
	}
}

func getSteam64(urlType int, s *discordgo.Session, i *discordgo.InteractionCreate, url string, steam64 int64) int64 {
	var err error

	if urlType == utils.BadUrl {
		s.InteractionRespond(i.Interaction, badUrlParameterErrorMsg)
	} else if urlType == utils.Steam64Url {
		steam64str := utils.GetSteam64FromUrl(url)

		steam64, err = strconv.ParseInt(steam64str, 10, 64)

		if err != nil {
			log.Println("1 - Erreur lors de la conversion de la string en int64 pour le steam64: ", err.Error())
			s.InteractionRespond(i.Interaction, errorMsg)
			return 0;
		}
	} else {
		customId := utils.GetCustomIdFromUrl(url)

		steam64, err = http.GetPlayerSteam64FromCustomId(customId)

		if err != nil {
			log.Println("2 - Erreur lors de getSteam64: ", err.Error())
			s.InteractionRespond(i.Interaction, errorMsg)
			return 0;
		}
	}
	return steam64
}
