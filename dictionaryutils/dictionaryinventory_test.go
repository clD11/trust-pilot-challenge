package dictionaryutils

import (
	"testing"
	"trust-pilot-challenge/letterinventory"
)

func TestRefineDictionary(t *testing.T) {
	words := []string{"aa", "bb", "cc"}
	var letterInventory = letterinventory.CreateLetterInventory("aabb")

	var testee = CreateDictionaryInventory(words)
	testee.RefineInventory(letterInventory)

	if key, ok := testee.dictionary["cc"]; ok {
		t.Error("Test failed: Key ", key, " should have been removed")
	}
}
