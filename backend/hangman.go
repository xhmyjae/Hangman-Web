package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
)

type Hangman struct {
	Difficulty   string
	MaxLives     int
	Lives        int
	PlayerWord   string
	TriedLetters string
	TriedWords   []string
	Word         string
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
	Hangman.TriedWords = []string{}
	Hangman.TriedLetters = ""
	Hangman.PlayerWord = strings.Repeat("_", len(word))
}

func (Hangman *Hangman) InitRandomWord() {
	var file string
	if Hangman.Difficulty == "ez" {
		file = "words.txt"
	} else if Hangman.Difficulty == "moyen" {
		file = "words2.txt"
	} else {
		file = "words3.txt"
	}
	Hangman.Init(GetRandomWord("./backend/resources/" + file))
}

func (Hangman *Hangman) GetLives() int {
	return Hangman.MaxLives - (len(Hangman.TriedLetters) + len(Hangman.TriedWords)*2)
}

func (Hangman *Hangman) TryWord(word string) bool {
	if word == Hangman.Word {
		return true
	}
	Hangman.TriedWords = append(Hangman.TriedWords, word)
	Hangman.Lives -= 2
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
		Hangman.Lives--
		return false
	}
}

func (Hangman *Hangman) IsFinished() bool {
	return Hangman.PlayerWord == Hangman.Word
}
