package events

import (
	"log"

	"github.com/b4cktr4ck5r3/rpl-discordbot/database"
	"github.com/bwmarrin/discordgo"
)

func RegisterEvents(session *discordgo.Session) {
	log.Println("Adding events...")
	session.AddHandler(ReadyEventHandler)
	session.AddHandler(GuildMemberRemoveEventHandler)
	log.Println("Events added !")
}

func ReadyEventHandler(s *discordgo.Session, r *discordgo.Ready) {
	log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
}

func GuildMemberRemoveEventHandler(s *discordgo.Session, r *discordgo.GuildMemberRemove) {
	discordId := r.Member.User.ID

	player, err := database.SelectPlayerByDiscordId(discordId)

	if err != nil {
		log.Printf("Erreur lors de la récupération du joueur qui vient de quitter le discord : %s", err.Error())
		return
	}

	ok, err := database.RemoveLinkFromUserWithDiscordId(player.SteamID)

	if !ok || err != nil {
		log.Printf("Erreur lors de la suppression du lien pour le joueur qui vient de quitter le discord: %s", err.Error())
		return
	}

	ok, err = database.DeleteSkinsWhereSteamId(player.SteamID)

	if !ok || err != nil {
		log.Printf("Erreur lors de la suppression des skins pour le joueur qui vient de quitter le discord: %s", err.Error())
		return
	}

	log.Printf("Le joueur Discord=%s et Steam=%s (DiscordID=%s et SteamID=%s) a bien été unlink et reset skin", r.Member.User.Username, player.Name, r.Member.User.ID, player.SteamID)
}
