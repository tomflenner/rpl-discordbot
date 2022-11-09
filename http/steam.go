package http

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/MrWaggel/gosteamconv"
	"github.com/b4cktr4ck5r3/rpl-discordbot/config"
	"github.com/b4cktr4ck5r3/rpl-discordbot/models"
)

func GetPlayerSteamAccountSummaries(steamId string) (models.SteamAccountSummaries, error) {
	var resPayload models.SteamApiPlayerSummariesResponse
	var playerSummaries models.SteamAccountSummaries

	steam64, err := gosteamconv.SteamStringToInt64(steamId)

	if err != nil {
		log.Println("Erreur lors de la conversion du SteamID en Steam64: ", err.Error())
		return playerSummaries, err
	}

	url := fmt.Sprintf("https://api.steampowered.com/ISteamUser/GetPlayerSummaries/v2/?key=%s&format=json&steamids=%d", config.Cfg.SteamApiKey, steam64)
	res, err := http.Get(url)

	if err != nil {
		log.Println("Erreur lors de la requête vers les serveur steam: ", err.Error())
		return playerSummaries, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err := json.Unmarshal(body, &resPayload); err != nil {
		log.Println("Erreur sur l'unmarshal de la réponse à la requête GetPlayerSteamAccountSummaries: ", err.Error())
		return playerSummaries, err
	}

	if len(resPayload.Response.SteamAccounts) != 1 {
		log.Println("Erreur sur le payload de la réponse fournit par l'API steam.")
		return playerSummaries, err
	}

	return resPayload.Response.SteamAccounts[0], err
}
