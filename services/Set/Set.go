package SetPokemon

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Set struct {
	CardCount struct {
		Total int `json:"total"`
	} `json:"cardCount"`
	ID          string `json:"id"`
	Logo        string `json:"logo"`
	Name        string `json:"name"`
	ReleaseDate string `json:"releaseDate"`
}

func SetP() []Set {
	result := []Set{}

	for i := 1; i <= 10; i++ {
		url := fmt.Sprintf("https://api.tcgdex.net/v2/en/sets/swsh%d", i)

		req, errReq := http.NewRequest("GET", url, nil)
		if errReq != nil {
			fmt.Printf("Requête - Erreur lors de l'initialisation de la requête : %s\n", errReq.Error())
			return []Set{}
		}

		httpClient := &http.Client{}
		res, errRes := httpClient.Do(req)
		if errRes != nil {
			fmt.Printf("Requête - Erreur lors de l'exécution de la requête : %s\n", errRes.Error())
			return []Set{}
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			fmt.Printf("Réponse - Erreur, code HTTP : %d, message : %s\n", res.StatusCode, res.Status)
			return []Set{}
		}

		var decodeData Set
		errDecode := json.NewDecoder(res.Body).Decode(&decodeData)
		if errDecode != nil {
			fmt.Printf("Decode - Erreur lors du décodage des données : %s\n", errDecode.Error())
			return []Set{}
		}
		result = append(result, decodeData)
	}
	return result
}
