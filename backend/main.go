package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type State struct {
	Menu string
	Clay Hangman
}

var state = State{Menu: "main", Clay: Hangman{}}
var ch = make(chan struct{}, 2)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	ch <- struct{}{}
	defer func() {
		<-ch
	}()

	http.DefaultServeMux.ServeHTTP(w, r)
}

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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		queries := r.URL.Query()
		if queries.Get("reload") == "true" {
			state.Clay.InitRandomWord()
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
		state.Menu = "main"
		tmpl.ExecuteTemplate(w, "main", state)
	})

	http.HandleFunc("/hangman", func(w http.ResponseWriter, r *http.Request) {
		queries := r.URL.Query()
		if queries.Has("difficulty") {
			state.Clay.Difficulty = queries.Get("difficulty")
			state.Clay.InitRandomWord()
			http.Redirect(w, r, "/hangman", http.StatusSeeOther)
		}
		letter := r.FormValue("word_text")
		if letter == "" {
			letter = r.URL.Query().Get("word_text")
		}
		state.Menu = "game"

		if letter != "" {
			state.Clay.loadGame(r, strings.ToLower(letter))
		}
		tmpl.ExecuteTemplate(w, "main", state)
	})

	http.HandleFunc("/rules", func(w http.ResponseWriter, r *http.Request) {
		state.Menu = "rules"
		tmpl.ExecuteTemplate(w, "main", state)
	})

	http.HandleFunc("/reload", func(w http.ResponseWriter, r *http.Request) {
		state.Clay.InitRandomWord()
		state.Menu = "game"
		http.Redirect(w, r, "/hangman", http.StatusSeeOther)
	})

	http.ListenAndServe(":8999", http.HandlerFunc(mainHandler))
}
