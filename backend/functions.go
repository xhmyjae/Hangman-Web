package main

import "net/http"

func (Hangman *Hangman) loadGame(r *http.Request, letter string) {
	if len(letter) > 1 {
		if Hangman.TryWord(letter) {
			state.Menu = "win"
		}
	} else {
		if Hangman.TryLetter(letter) {
			if Hangman.IsFinished() {
				state.Menu = "win"
			}
		}
	}
	if Hangman.Lives <= 0 {
		state.Menu = "lose"
	}
}
