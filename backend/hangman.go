package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
)

type Hangman struct {
	ActualWord string `json:"actual_word"`
	Letters    string `json:"letters"`
	Lives      int    `json:"lives"`
	Word       string `json:"word"`
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
	Hangman.Lives = 10
	Hangman.Word = word
	Hangman.ActualWord = strings.Repeat("_", len(word))
}

func (Hangman *Hangman) TryLetter(letter string) bool {
	if strings.Contains(Hangman.Word, letter) {
		for i, l := range Hangman.Word {
			if string(l) == letter {
				Hangman.ActualWord = Hangman.ActualWord[:i] + letter + Hangman.ActualWord[i+1:]
			}
		}
		return true
	} else {
		Hangman.Letters += letter
		return false
	}
}

func (Hangman *Hangman) IsFinished() bool {
	return Hangman.ActualWord == Hangman.Word
}
