package solve

import (
	"log"
	"strings"
	"trust-pilot-challenge/hashutils"
	li "trust-pilot-challenge/letterinventory"
	lim "trust-pilot-challenge/letterinventorymap"
)

var result = []string{}

type Backtrack struct {
	letterInventoryMap lim.LetterInventoryMap
}

func CreateBacktrack(letterInventoryMap lim.LetterInventoryMap) Backtrack {
	return Backtrack{letterInventoryMap: letterInventoryMap}
}

func (b Backtrack) Solve(letters li.LetterInventory, letterInventoryMap lim.LetterInventoryMap, targetHash string) bool {

	if b.isBadNode() {
		return false
	}

	if b.isCandidateNode(letters) {
		if b.isGoalNode(result, targetHash) {
			log.Println("FOUND", result)
			return true
		}
		return false
	}

	for key, value := range letterInventoryMap.Dictionary() {
		// explore option
		letters.Subtract(value)		

		if !letters.ContainsNegative() {

			result = append(result, key)

			if b.Solve(letters, letterInventoryMap, targetHash) {
				return true
			}

			result = append(result[0:len(result)-1], result[len(result):]...)
		}

		// fail add word back
		letters.Add(value)		
	}

	return false
}

func (b Backtrack) isBadNode() bool {
	return len(result) > 3
}

func (b Backtrack) isCandidateNode(letters li.LetterInventory) bool {
	return letters.Size() < 1
}

func (b Backtrack) isGoalNode(result []string, targetHash string) bool {
	return hashutils.IsTargetHash(strings.Join(result, " "), targetHash)
}
