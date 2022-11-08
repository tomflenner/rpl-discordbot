package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/b4cktr4ck5r3/rpl-discordbot/config"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func InitializeDatabaseConnection() {
	var err error

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.Cfg.DbUser,
		config.Cfg.DbPassword,
		config.Cfg.DbHost,
		config.Cfg.DbPort,
		config.Cfg.DbName)

	Db, err = sql.Open("mysql", connectionString)

	if err != nil {
		log.Fatal("Error on database open connection: " + err.Error())
	}

	if err := Db.Ping(); err != nil {
		log.Fatal("Error on database ping: " + err.Error())
	}
}
