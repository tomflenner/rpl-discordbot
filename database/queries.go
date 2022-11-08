package database

import (
	_ "embed"

	"github.com/b4cktr4ck5r3/rpl-discordbot/models"
)

var (
	//go:embed sql/selectplayerwherelinkcode.sql
	queryPlayerWhereLinkCode string

	//go:embed sql/updateplayer.sql
	queryUpdatePlayer string
)

func SelectPlayerByLinkCode(linkCode string) (models.Player, error) {
	row := Db.QueryRow(queryPlayerWhereLinkCode, linkCode)
	player := models.Player{}

	err := row.Scan(
		&player.SteamID,
		&player.Name,
		&player.PermsLvl,
		&player.DiscordID,
		&player.LinkCode,
	)

	return player, err
}

func UpdatePlayer(player models.Player) (bool, error) {
	_, err := Db.Exec(queryUpdatePlayer, player.DiscordID, player.SteamID, player.LinkCode.String)

	if err != nil {
		return false, err
	} else {
		return true, err
	}
}
