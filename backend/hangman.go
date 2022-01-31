package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
)

type Hangman struct {
	MaxLives     int      `json:"max_lives"`
	PlayerWord   string   `json:"player_word"`
	TriedLetters string   `json:"tried_letters"`
	TriedWords   []string `json:"tried_words"`
	Word         string   `json:"word"`
}

func GetRandomWord(filename string) string {
	words := strings.Split(ReadFile(filename), "\n")
	return strings.ReplaceAll(words[rand.Intn(len(words)-1)], "\r", "")
}

func ReadFile(filename string) string {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(file)
}

func (Hangman *Hangman) Init(word string) {
	Hangman.MaxLives = 10
	Hangman.Word = word
	Hangman.PlayerWord = strings.Repeat("_", len(word))
}

func (Hangman *Hangman) GetLives() int {
	return Hangman.MaxLives - (len(Hangman.TriedLetters) + len(Hangman.TriedWords)*2)
}

func (Hangman *Hangman) TryWord(word string) bool {
	if word == Hangman.Word {
		Hangman.TriedWords = append(Hangman.TriedWords, word)
		return true
	}
	return false
}

func (Hangman *Hangman) TryLetter(letter string) bool {
	if strings.Contains(Hangman.Word, letter) {
		for i, l := range Hangman.Word {
			if string(l) == letter {
				Hangman.PlayerWord = Hangman.PlayerWord[:i] + letter + Hangman.PlayerWord[i+1:]
			}
		}
		return true
	} else {
		Hangman.TriedLetters += letter
		return false
	}
}

func (Hangman *Hangman) IsFinished() bool {
	return Hangman.PlayerWord == Hangman.Word
}
