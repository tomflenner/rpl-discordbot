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
const StatsCommandName = "stats"
const StatsDiscordCommandName = "stats-discord"
const StatsSteamProfileCommandName = "stats-steam-profile"

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
	playerNotFoundWithDiscordIdErrorMsg = &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Le joueur n'a pas lié son compte Discord et son compte Steam.",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	}

	playerNotFoundWithSteamProfileErrorMsg = &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Ce profile steam n'est pas dans le classement.",
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

	badUrlParameterErrorMsg = &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "L'URL que vous avez fournit n'est pas correcte.",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	}
)

func StatsDiscordCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if canExecuteRestrictedCommand(i, config.Cfg.StatsChannelID) {
		options := i.ApplicationCommandData().Options
		player, err := database.SelectPlayerByDiscordId(options[0].UserValue(s).ID)

		if err != nil {
			s.InteractionRespond(i.Interaction, playerNotFoundWithDiscordIdErrorMsg)
		}

		playerStats, err := http.GetPlayerStats(player.SteamID)

		if err != nil {
			s.InteractionRespond(i.Interaction, playerStatRequestErrorMsg)
		}

		steam64, err := gosteamconv.SteamStringToInt64(player.SteamID)

		if err != nil {
			log.Println("Erreur lors de la conversion du SteamID en Steam64")
			s.InteractionRespond(i.Interaction, errorMsg)
		}

		playerSummaries, err := http.GetPlayerSteamAccountSummaries(steam64)

		if err != nil {
			s.InteractionRespond(i.Interaction, playerSummariesRequestErrorMsg)
		}

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

func StatsCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if canExecuteRestrictedCommand(i, config.Cfg.StatsChannelID) {
		player, err := database.SelectPlayerByDiscordId(i.Member.User.ID)

		if err != nil {
			s.InteractionRespond(i.Interaction, playerNotFoundWithDiscordIdErrorMsg)
		}

		playerStats, err := http.GetPlayerStats(player.SteamID)

		if err != nil {
			s.InteractionRespond(i.Interaction, playerStatRequestErrorMsg)
		}

		steam64, err := gosteamconv.SteamStringToInt64(player.SteamID)

		if err != nil {
			log.Println("Erreur lors de la conversion du SteamID en Steam64")
			s.InteractionRespond(i.Interaction, errorMsg)
		}

		playerSummaries, err := http.GetPlayerSteamAccountSummaries(steam64)

		if err != nil {
			s.InteractionRespond(i.Interaction, playerSummariesRequestErrorMsg)
		}

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

func StatsSteamCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if canExecuteRestrictedCommand(i, config.Cfg.StatsChannelID) {
		options := i.ApplicationCommandData().Options

		var steam64 int64
		var err error

		url := options[0].StringValue()
		urlType := utils.GetUrlType(url)

		if urlType == utils.BadUrl {
			s.InteractionRespond(i.Interaction, badUrlParameterErrorMsg)
		} else if urlType == utils.Steam64Url {
			steam64str := utils.GetSteam64FromUrl(url)

			steam64, err = strconv.ParseInt(steam64str, 10, 64)

			if err != nil {
				log.Println("Erreur lors de la conversion de la string en int64 pour le steam64")
				s.InteractionRespond(i.Interaction, errorMsg)
			}
		} else {
			customId := utils.GetCustomIdFromUrl(url)

			steam64, err = http.GetPlayerSteam64FromCustomId(customId)

			if err != nil {
				s.InteractionRespond(i.Interaction, errorMsg)
			}
		}

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

		playerStats, err := http.GetPlayerStats(player.SteamID)

		if err != nil {
			s.InteractionRespond(i.Interaction, playerStatRequestErrorMsg)
		}

		playerSummaries, err := http.GetPlayerSteamAccountSummaries(steam64)

		if err != nil {
			s.InteractionRespond(i.Interaction, playerSummariesRequestErrorMsg)
		}

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
