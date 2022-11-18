package config

import (
	"fmt"
	"log"
	"os"

	"github.com/b4cktr4ck5r3/rpl-discordbot/models"
	"github.com/joho/godotenv"
)

func loadDotEnv() {
	var err error = godotenv.Load(".env")

	if err != nil {
		log.Println("WARNING! Error loading .env file: " + err.Error())
	}
}

func getConfigValue(key string) string {
	var value string = os.Getenv(key)

	if value == "" {
		log.Fatalf(fmt.Sprintf("Error loading %s from env", key))
	}

	return value
}

var Cfg models.Config

func InitializeConfig() {
	loadDotEnv()

	Cfg = models.Config{}

	var botToken string = getConfigValue("BOT_TOKEN")
	Cfg.BotToken = botToken

	var dbUser string = getConfigValue("DB_USER")
	Cfg.DbUser = dbUser

	var dbPassword string = getConfigValue("DB_PASSWORD")
	Cfg.DbPassword = dbPassword

	var dbHost string = getConfigValue("DB_HOST")
	Cfg.DbHost = dbHost

	var dbPort string = getConfigValue("DB_PORT")
	Cfg.DbPort = dbPort

	var dbLinkName string = getConfigValue("DB_LINK_NAME")
	Cfg.DbLinkName = dbLinkName

	var dbSkinsName string = getConfigValue("DB_SKINS_NAME")
	Cfg.DbSkinsName = dbSkinsName

	var guildId string = getConfigValue("GUILD_ID")
	Cfg.GuildID = guildId

	var statsChannelId string = getConfigValue("STATS_CHANNEL_ID")
	Cfg.StatsChannelID = statsChannelId

	var linkChannelID string = getConfigValue("LINK_CHANNEL_ID")
	Cfg.LinkChannelID = linkChannelID

	var steamApiKey string = getConfigValue("STEAM_API_KEY")
	Cfg.SteamApiKey = steamApiKey

	var statsApiUrl string = getConfigValue("STATS_API_URL")
	Cfg.StatsApiUrl = statsApiUrl
}
