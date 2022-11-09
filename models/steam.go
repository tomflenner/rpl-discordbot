package models

type SteamApiPlayerSummariesResponse struct {
	Response PlayerSummariesResponse `json:"response"`
}

type PlayerSummariesResponse struct {
	SteamAccounts []SteamAccountSummaries `json:"players"`
}

type SteamAccountSummaries struct {
	Avatar                   string `json:"avatar"`
	Avatarfull               string `json:"avatarfull"`
	Avatarhash               string `json:"avatarhash"`
	Avatarmedium             string `json:"avatarmedium"`
	Commentpermission        int64  `json:"commentpermission"`
	Communityvisibilitystate int64  `json:"communityvisibilitystate"`
	Lastlogoff               int64  `json:"lastlogoff"`
	Personaname              string `json:"personaname"`
	Personastate             int64  `json:"personastate"`
	Personastateflags        int64  `json:"personastateflags"`
	Primaryclanid            string `json:"primaryclanid"`
	Profilestate             int64  `json:"profilestate"`
	Profileurl               string `json:"profileurl"`
	Realname                 string `json:"realname"`
	Steamid                  string `json:"steamid"`
	Timecreated              int64  `json:"timecreated"`
}

type SteamApiCustomIDResolverResponse struct {
	Response CustomIDResolverResponse `json:"response"`
}

type CustomIDResolverResponse struct {
	Steamid string `json:"steamid"`
	Success int64  `json:"success"`
}
