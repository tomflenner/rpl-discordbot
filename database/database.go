package database

import (
	"database/sql"
	"log"

	"github.com/b4cktr4ck5r3/rpl-discordbot/config"
	"github.com/b4cktr4ck5r3/rpl-discordbot/utils"
	_ "github.com/go-sql-driver/mysql"
)

var DbLink *sql.DB
var DbSkins *sql.DB

func InitializeDatabaseConnection() {
	initializeLinkDatabaseConn()
	initializeSkinsDatabaseConn()
}

func initializeSkinsDatabaseConn() {
	var err error

	dbSkinsConnString := utils.GetDatabaseDNS(
		config.Cfg.DbUser,
		config.Cfg.DbPassword,
		config.Cfg.DbHost,
		config.Cfg.DbPort,
		config.Cfg.DbSkinsName,
	)

	DbSkins, err = sql.Open("mysql", dbSkinsConnString)

	if err != nil {
		log.Fatal("Error on skins database open connection: " + err.Error())
	}

	if err := DbSkins.Ping(); err != nil {
		log.Fatal("Error on skins database ping: " + err.Error())
	}
}

func initializeLinkDatabaseConn() {
	var err error

	dbLinkConnString := utils.GetDatabaseDNS(
		config.Cfg.DbUser,
		config.Cfg.DbPassword,
		config.Cfg.DbHost,
		config.Cfg.DbPort,
		config.Cfg.DbLinkName,
	)

	DbLink, err = sql.Open("mysql", dbLinkConnString)

	if err != nil {
		log.Fatal("Error on link database open connection: " + err.Error())
	}

	if err := DbLink.Ping(); err != nil {
		log.Fatal("Error on link database ping: " + err.Error())
	}
}
