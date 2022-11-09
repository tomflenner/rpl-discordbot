package models

type StatsApiPlayerResponse struct {
	ID               int64   `json:"id"`
	SteamID          string  `json:"steam_id"`
	Name             string  `json:"name"`
	Score            int64   `json:"score"`
	Rank             int64   `json:"rank"`
	MVP              int64   `json:"mvp"`
	Kills            int64   `json:"kills"`
	Deaths           int64   `json:"deaths"`
	Ratio            float64 `json:"ratio"`
	Headshots        int64   `json:"headshots"`
	HeadshotsPercent int64   `json:"headshots_percent"`
	Assists          int64   `json:"assists"`
	FlashAssists     int64   `json:"flash_assists"`
	NoScope          int64   `json:"no_scope"`
	ThruSmoke        int64   `json:"thru_smoke"`
	Blind            int64   `json:"blind"`
	Wallbang         int64   `json:"wallbang"`
}
