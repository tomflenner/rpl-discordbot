package http

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/b4cktr4ck5r3/rpl-discordbot/config"
	"github.com/b4cktr4ck5r3/rpl-discordbot/models"
)

func GetPlayerSteamAccountSummaries(steamId int64) (models.SteamAccountSummaries, error) {
	var payloadResponse models.SteamApiPlayerSummariesResponse
	var playerSummaries models.SteamAccountSummaries

	url := fmt.Sprintf("https://api.steampowered.com/ISteamUser/GetPlayerSummaries/v2/?key=%s&format=json&steamids=%d", config.Cfg.SteamApiKey, steamId)
	res, err := http.Get(url)

	if err != nil {
		log.Println("Erreur lors de la requête vers les serveur steam dans GetPlayerSteamAccountSummaries: ", err.Error())
		return playerSummaries, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err := json.Unmarshal(body, &payloadResponse); err != nil {
		log.Println("Erreur sur l'unmarshal de la réponse à la requête GetPlayerSteamAccountSummaries: ", err.Error())
		return playerSummaries, err
	}

	if len(payloadResponse.Response.SteamAccounts) != 1 {
		log.Println("Erreur sur le payload de la réponse fournit par l'API steam.")
		return playerSummaries, err
	}

	return payloadResponse.Response.SteamAccounts[0], err
}

func GetPlayerSteam64FromCustomId(customId string) (int64, error) {
	var steam64 int64
	var err error
	var payloadResponse models.SteamApiCustomIDResolverResponse

	url := fmt.Sprintf("http://api.steampowered.com/ISteamUser/ResolveVanityURL/v0001/?key=%s&vanityurl=%s", config.Cfg.SteamApiKey, customId)
	res, err := http.Get(url)

	if err != nil {
		log.Println("Erreur lors de la requête vers les serveur steam dans GetPlayerSteam64FromCustomId: ", err.Error())
		return steam64, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err := json.Unmarshal(body, &payloadResponse); err != nil {
		log.Println("Erreur sur l'unmarshal de la réponse à la requête GetPlayerSteam64FromCustomId: ", err.Error())
		return steam64, err
	}

	//See https://wiki.teamfortress.com/wiki/WebAPI/ResolveVanityURL
	if payloadResponse.Response.Success == 1 {
		steam64, err = strconv.ParseInt(payloadResponse.Response.Steamid, 10, 64)
	} else {
		log.Println("No match from SteamAPI")
	}

	return steam64, err
}
