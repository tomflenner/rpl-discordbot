package models

import "database/sql"

type Player struct {
	SteamID   string         `sql:"steam_id"`
	Name      string         `sql:"name"`
	PermsLvl  int16          `sql:"perms_lvl"`
	DiscordID sql.NullString `sql:"discord_id"`
	LinkCode  sql.NullString `sql:"link_code"`
}
