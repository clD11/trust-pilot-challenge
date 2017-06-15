package solve

import (
	"log"
	"testing"
)

func TestCreateLetter(t *testing.T) {
	slice := []string {"a", "aa", "aa", "aaa"}
	actual := CreateLetter(slice)
	log.Println(actual)
}