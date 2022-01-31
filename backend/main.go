package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type State struct {
	Menu string
	Clay Hangman
}

var state = State{Menu: "main", Clay: Hangman{}}

func main() {
	rand.Seed(time.Now().UnixNano())

	tmpl, err := template.ParseGlob("./templates/*.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tmpl.DefinedTemplates())

	css := http.FileServer(http.Dir("./client/style"))
	js := http.FileServer(http.Dir("./client/scripts"))
	resources := http.FileServer(http.Dir("./backend/resources/"))
	http.Handle("/static/", http.StripPrefix("/static/", css))
	http.Handle("/js/", http.StripPrefix("/js/", js))
	http.Handle("/resources/", http.StripPrefix("/resources/", resources))


	state.Clay.Init(GetRandomWord("./backend/resources/words.txt"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		state.Menu = "main"
		tmpl.ExecuteTemplate(w, "main", state)
	})

	http.HandleFunc("/hangman", func(w http.ResponseWriter, r *http.Request) {
		state.Menu = "game"

		state.Clay.loadGame(r)
		tmpl.ExecuteTemplate(w, "main", state)
	})

	http.HandleFunc("/rules", func(w http.ResponseWriter, r *http.Request) {
		state.Menu = "rules"
		tmpl.ExecuteTemplate(w, "main", state)
	})

	http.ListenAndServe(":8999", nil)

}
