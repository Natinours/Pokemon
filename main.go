package main

import (
	CartePokemon "api/services/Carte"
	SetPokemon "api/services/Set"
	"fmt"
	"net/http"
	"text/template"
)

var temp, tempErr = template.ParseGlob("./templates/*.html")

func main() {
	// Récupération des templates
	if tempErr != nil {
		fmt.Printf("ERREUR => %s\n", tempErr.Error())
		return
	}

	// Route pour la page d'accueil
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/Set", http.StatusSeeOther)
	})

	http.HandleFunc("/Set", func(w http.ResponseWriter, r *http.Request) {
		type DataStruct struct {
			Data []SetPokemon.Set
		}
		DataToSend := DataStruct{
			Data: SetPokemon.SetP(),
		}
		err := temp.ExecuteTemplate(w, "Set", DataToSend)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/Carte", func(w http.ResponseWriter, r *http.Request) {
		type DataStruct struct {
			Data []CartePokemon.Carte
		}
		DataToSend := DataStruct{
			Data: CartePokemon.CarteP(),
		}
		err := temp.ExecuteTemplate(w, "Carte", DataToSend)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	fmt.Println("Serveur démarré sur http://localhost:6969/")
	err := http.ListenAndServe(":6969", nil)
	if err != nil {
		fmt.Printf("Erreur lors du démarrage du serveur : %s\n", err.Error())
	}
}
