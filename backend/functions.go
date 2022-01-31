package main

import "net/http"

func (Hangman *Hangman) loadGame(r *http.Request) {
	letter := r.FormValue("word_tried")
	if len(letter) > 1 {
		if Hangman.TryWord(letter) {
			if Hangman.IsFinished() {

			}
		} else {

		}

	} else {
		if Hangman.TryLetter(letter) {
			if Hangman.IsFinished() {

			}
		}

	}

}
