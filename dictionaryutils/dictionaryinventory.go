package dictionaryutils

import (
	"sort"
	"strings"
	li "trust-pilot-challenge/letterinventory"
)

type DictionaryInventory struct {
	dictionary map[string][]string
}

func CreateDictionaryInventory(words []string) DictionaryInventory {
	di := DictionaryInventory{}
	di.dictionary = make(map[string][]string)
	for _, word := range words {
		str := sortString(word)
		di.dictionary[str] = append(di.dictionary[str], word)
	}
	return di
}

func (d *DictionaryInventory) RefineInventory(letterInv li.LetterInventory) DictionaryInventory {
	words := []string{}
	for k, v := range d.dictionary {
		if letterInv.Contains(k) {
			words = append(words, v...)
		}
	}
	return CreateDictionaryInventory(words)
}

func (d *DictionaryInventory) Values() []string {
	words := []string{}
	for _, v := range d.dictionary {
		words = append(words, v...)
	}
	return words
}

func (d *DictionaryInventory) Size() int {
	return len(d.dictionary)
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
