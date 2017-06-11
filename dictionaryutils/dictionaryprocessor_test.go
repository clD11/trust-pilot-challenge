package dictionaryutils

import (
	"log"
	"testing"
)

func TestCreateRelevantWordInventory(t *testing.T) {
	dict := CreateRelevantWordInventory("D:/Code/Go/Workspace/src/trust-pilot-challenge/wordlist.txt", "poultryoutwitsants", 0)

	count := 0
	for key, value := range dict {
		log.Println(key, value)
		count++
	}

	log.Println("count", count)
}
