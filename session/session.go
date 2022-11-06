package session

import (
	"log"

	"github.com/b4cktr4ck5r3/rpl-discordbot/commands"
	"github.com/b4cktr4ck5r3/rpl-discordbot/config"
	"github.com/bwmarrin/discordgo"
)

var S *discordgo.Session
var RegisteredCommand []*discordgo.ApplicationCommand

func InitializeSession() {
	var err error
	S, err = discordgo.New("Bot " + config.Cfg.BotToken)

	//Display message on ready
	S.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	err = S.Open()

	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}
}

func RegisterCommands() {
	log.Println("Adding commands...")
	RegisteredCommand := make([]*discordgo.ApplicationCommand, len(commands.Commands))
	for i, v := range commands.Commands {
		cmd, err := S.ApplicationCommandCreate(S.State.User.ID, "", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		RegisteredCommand[i] = cmd
	}

	//Adding an event that trigger on registered command call
	S.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commands.CommandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

func RemoveCommands() {
	log.Println("Removing commands...")
	for _, v := range RegisteredCommand {
		err := S.ApplicationCommandDelete(S.State.User.ID, "", v.ID)
		if err != nil {
			log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
		}
	}
}
