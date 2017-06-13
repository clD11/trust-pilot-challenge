package main

import (
	"log"
	"trust-pilot-challenge/letterinventory"
	"trust-pilot-challenge/letterinventorymap"
	"trust-pilot-challenge/solve"
	"trust-pilot-challenge/wordlistprocessor"
)

const (
	TEST    = "8b35bbd7ff2f5dd7c94fffbb1a3512bc"
	EASY    = "e4820b45d2277f3844eac66c903e84be"
	MEDIUM  = "23170acc097c24edb98fc5488ab033fe"
	HARD    = "665e5bcb0c20062fe8abaaf4628bb154"
	ANAGRAM = "ailnooprssttttuuwy"
)

func main() {
	log.Println("Running")

	wordInventory := wordlistprocessor.ProcessWordlist(
		"D:/Code/Go/Workspace/src/trust-pilot-challenge/wordlist.txt", ANAGRAM, 2)

	log.Println("solving with", len(wordInventory), "words")

	findUsingBacktrack(ANAGRAM, EASY, wordInventory)

	log.Println("Finished")
}

func findUsingBacktrack(anagram, targetHash string, wordlist []string) {
	letterInventory := letterinventory.CreateLetterInventory(anagram)
	letterInventoryMap := letterinventorymap.CreateLetterInventoryMap(wordlist)

	backtrack := solve.CreateBacktrack(letterInventoryMap)
	backtrack.Solve(letterInventory, letterInventoryMap, targetHash)
}

func findUsingPermutations(permSize int, targetHash string, wordlist []string) {
	perm := solve.CreatePermutation(permSize, wordlist)
	perm.Solve(targetHash)
}
