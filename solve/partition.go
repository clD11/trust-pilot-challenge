package solve

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
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

// read in parts
func ReadPartitions(filepath string, len int) {
	var partitions [3]int

	file := openFile(filepath)
	defer file.Close()

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}

		//partitions = append(partitions, record...) // record has the type []string
	}
}

func openFile(filePath string) *os.File {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return file
}
