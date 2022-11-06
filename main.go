package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/b4cktr4ck5r3/rpl-discordbot/config"
	"github.com/b4cktr4ck5r3/rpl-discordbot/session"
)

func main() {
	config.InitializeConfig()
	session.InitializeSession()
	session.RegisterCommands()

	defer session.S.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	session.RemoveCommands()
	log.Println("Gracefully shutting down.")
}
