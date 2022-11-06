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
		log.Fatal("Error loading .env file: " + err.Error())
	}
}

func getConfigValue(key string) string {
	var value string
	value = os.Getenv(key)

	if value == "" {
		log.Fatal(fmt.Sprintf("Error loading %s from env", key))
	}

	return value
}

var Cfg models.Config

func InitializeConfig() {
	loadDotEnv()

	Cfg = models.Config{}

	var botToken string
	botToken = getConfigValue("BOT_TOKEN")
	Cfg.BotToken = botToken
}
