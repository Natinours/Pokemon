package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Type personnalisé pour gérer les dégâts (string ou int)
type DamageType struct {
	Value string
}

// Méthode pour désérialiser (unmarshal) Damage depuis JSON
func (d *DamageType) UnmarshalJSON(data []byte) error {
	// Vérifie si la donnée est un nombre
	var intValue int
	if err := json.Unmarshal(data, &intValue); err == nil {
		d.Value = strconv.Itoa(intValue) // Convertit le nombre en string
		return nil
	}

	// Sinon, considère que c'est une chaîne
	var stringValue string
	if err := json.Unmarshal(data, &stringValue); err == nil {
		d.Value = stringValue
		return nil
	}

	// Retourne une erreur si aucun des cas n'est valide
	return fmt.Errorf("DamageType : impossible de convertir les données %s", string(data))
}

// Structure principale
type Carte struct {
	Category    string   `json:"category"`
	ID          string   `json:"id"`
	Illustrator string   `json:"illustrator"`
	Image       string   `json:"image"`
	Name        string   `json:"name"`
	Rarity      string   `json:"rarity"`
	DexID       []int    `json:"dexId"`
	Hp          int      `json:"hp"`
	Types       []string `json:"types"`
	Stage       string   `json:"stage"`
	Attacks     []struct {
		Cost   []string   `json:"cost"`
		Name   string     `json:"name"`
		Effect string     `json:"effect"`
		Damage DamageType `json:"damage,omitempty"`
	} `json:"attacks"`
	Weaknesses []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"weaknesses"`
	Retreat int `json:"retreat"`
}

func main() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 150; j++ {
			url := fmt.Sprintf("https://api.tcgdex.net/v2/en/cards/swsh%d-%d", i, j)

			req, errReq := http.NewRequest("GET", url, nil)
			if errReq != nil {
				fmt.Printf("Requête - Erreur lors de l'initialisation de la requête : %s\n", errReq.Error())
				return
			}

			httpClient := &http.Client{}
			res, errRes := httpClient.Do(req)
			if errRes != nil {
				fmt.Printf("Requête - Erreur lors de l'exécution de la requête : %s\n", errRes.Error())
				return
			}
			defer res.Body.Close()

			if res.StatusCode != http.StatusOK {
				fmt.Printf("Réponse - Erreur, code HTTP : %d, message : %s\n", res.StatusCode, res.Status)
				return
			}

			var decodeData Carte
			errDecode := json.NewDecoder(res.Body).Decode(&decodeData)
			if errDecode != nil {
				fmt.Printf("Decode - Erreur lors du décodage des données : %s\n", errDecode.Error())
				return
			}

			fmt.Printf("Carte récupérée :\n")
			fmt.Printf("Category : %s\n", decodeData.Category)
			fmt.Printf("ID : %s\n", decodeData.ID)
			fmt.Printf("Nom : %s\n", decodeData.Name)
			fmt.Printf("Illustrateur : %s\n", decodeData.Illustrator)
			fmt.Printf("Image : %s/high.webp\n", decodeData.Image)
			fmt.Printf("Rareté : %s\n", decodeData.Rarity)
			fmt.Printf("DexID : %v\n", decodeData.DexID)
			fmt.Printf("HP : %d\n", decodeData.Hp)
			fmt.Printf("Types : %v\n", decodeData.Types)
			fmt.Printf("Stage : %s\n", decodeData.Stage)
			fmt.Printf("Attaques :\n")
			for _, attack := range decodeData.Attacks {
				fmt.Printf("\tNom : %s\n", attack.Name)
				fmt.Printf("\tCoût : %v\n", attack.Cost)
				fmt.Printf("\tEffet : %s\n", attack.Effect)
				fmt.Printf("\tDégâts : %s\n", attack.Damage.Value)
			}
			fmt.Printf("Faiblesses :\n")
			for _, weakness := range decodeData.Weaknesses {
				fmt.Printf("\tType : %s\n", weakness.Type)
				fmt.Printf("\tValeur : %s\n", weakness.Value)
			}
			fmt.Printf("Retraite : %d\n", decodeData.Retreat)
			fmt.Printf("\n")
		}
	}
}
