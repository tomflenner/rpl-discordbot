package database

import (
	_ "embed"

	"github.com/b4cktr4ck5r3/rpl-discordbot/models"
)

var (
	//go:embed sql/selectplayerwherelinkcode.sql
	queryPlayerWhereLinkCode string

	//go:embed sql/selectplayerwherediscordid.sql
	queryPlayerWhereDiscordId string

	//go:embed sql/selectplayerwheresteamid.sql
	queryPlayerWhereSteamId string

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
		&player.CountUnlink,
	)

	return player, err
}

func SelectPlayerByDiscordId(discordId string) (models.Player, error) {
	row := Db.QueryRow(queryPlayerWhereDiscordId, discordId)
	player := models.Player{}

	err := row.Scan(
		&player.SteamID,
		&player.Name,
		&player.PermsLvl,
		&player.DiscordID,
		&player.LinkCode,
		&player.CountUnlink,
	)

	return player, err
}

func SelectPlayerBySteamId(steamId string) (models.Player, error) {
	row := Db.QueryRow(queryPlayerWhereSteamId, steamId)
	player := models.Player{}

	err := row.Scan(
		&player.SteamID,
		&player.Name,
		&player.PermsLvl,
		&player.DiscordID,
		&player.LinkCode,
		&player.CountUnlink,
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
