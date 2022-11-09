package http

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/b4cktr4ck5r3/rpl-discordbot/models"
)

func GetPlayerStats(steamId string) (models.StatsApiPlayerResponse, error) {
	var playerStats models.StatsApiPlayerResponse

	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/api/players/%s", steamId))

	if err != nil {
		log.Println("Erreur sur la requête GetPlayerStat: ", err.Error())
		return playerStats, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &playerStats); err != nil {
		log.Println("Erreur sur l'unmarshal de la réponse à la requête GetPlayerStats: ", err.Error())
		return playerStats, err
	}

	return playerStats, err
}
