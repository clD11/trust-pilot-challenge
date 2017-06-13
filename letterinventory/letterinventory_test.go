package letterinventory

import (
	"log"
	"testing"
)

func TestLetterInventoryAdd(t *testing.T) {
	str := "abc"
	testee := CreateLetterInventory(str)
	li := CreateLetterInventory(str)

	testee.Add(li)
	log.Println(testee)
}
