package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

var clay Hangman

func main() {
	rand.Seed(time.Now().UnixNano())

	tmpl := template.Must(template.ParseFiles("../client/pages/game.html"))

	fs := http.FileServer(http.Dir("../client/style"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {


		println("test")
		clay.Init(GetRandomWord("resources/words.txt"))

		tmpl.Execute(w, clay)
	})

	http.ListenAndServe(":8999", nil)

}
