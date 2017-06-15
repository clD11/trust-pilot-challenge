package solve

import (
	li "trust-pilot-challenge/letterinventory"
)

// create length dictionary words
func CreateLetter(wordlist []string) map[int][]li.LetterInventory {
	letterInventory := make(map[int][]li.LetterInventory)

	for _, word := range wordlist {
		letterInventory[len(word)] = append(letterInventory[len(word)], li.CreateLetterInventory(word))
	}

	return letterInventory
}