package letterinventorymap

import (
	li "trust-pilot-challenge/letterinventory"
)

type LetterInventoryMap struct {
	dictionary map[string]li.LetterInventory
}

func CreateLetterInventoryMap(wordlist []string) LetterInventoryMap {
	dictionary := make(map[string]li.LetterInventory)
	for _, word := range wordlist {
		dictionary[word] = li.CreateLetterInventory(word)
	}
	return LetterInventoryMap{dictionary: dictionary}
}

func (lim *LetterInventoryMap) Dictionary() map[string]li.LetterInventory {
	return lim.dictionary
}

func (lim *LetterInventoryMap) Size() int {
	return len(lim.dictionary)
}
