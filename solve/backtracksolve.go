package solve

import (
	"log"
	"strings"
	du "trust-pilot-challenge/dictionaryutils"
	"trust-pilot-challenge/hashutils"
	"trust-pilot-challenge/letterinventory"
)

var result = []string{}

type Backtrack struct {
	phrase []string
}

func CreateBacktrack() Backtrack {
	return Backtrack{phrase: []string{}}
}

func (b Backtrack) Solve(letters letterinventory.LetterInventory, targetHash string, wordlist du.DictionaryInventory) bool {

	if len(result) > 3 {
		return false
	}

	if b.isGoalNode(letters) {
		log.Println("testing", result)
		if hashutils.IsTargetHash(strings.Join(result, " "), targetHash) {
			log.Println(result)
			return true
		}
		return false
	}

	if b.isBadNode(letters, wordlist) {
		return false
	}

	for _, word := range wordlist.Values() {
		// remove chars from letInv add to result
		letters.Subtract(word)
		result = append(result, word)

		words := wordlist.RefineInventory(letters)

		if b.Solve(letters, targetHash, words) {
			return true
		}

		// fail add word back
		letters.Add(word)
		result = append(result[0:len(result)-1], result[len(result):]...)
	}

	return false
}

func (b Backtrack) isGoalNode(letters letterinventory.LetterInventory) bool {
	return letters.GetSize() < 1
}

func (b Backtrack) isBadNode(letters letterinventory.LetterInventory, wordlist du.DictionaryInventory) bool {
	return letters.GetSize() > 1 && wordlist.Size() < 1
}
