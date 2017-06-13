package wordlistprocessor

import (
	"bufio"
	"log"
	"os"
	"strings"
	"trust-pilot-challenge/letterinventory"
)

func ProcessWordlist(filepath, anagram string, min int) []string {
	file := openFile(filepath)
	defer file.Close()
	return processWordFile(file, anagram, min)
}

func openFile(filePath string) *os.File {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func processWordFile(file *os.File, anagram string, min int) []string {
	relevantWords := []string{}
	letterInv := letterinventory.CreateLetterInventory(anagram)

	var previousLine string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		currentLine := scanner.Text()
		letterInventory := letterinventory.CreateLetterInventory(currentLine)

		if !isDuplicate(currentLine, previousLine) && !containsSpecialCharacter(currentLine) &&
			isMinimumLength(currentLine, min) && letterInv.Contains(letterInventory) {
			relevantWords = append(relevantWords, currentLine)
		}

		previousLine = currentLine
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return relevantWords
}

func isDuplicate(currentLine string, previousLine string) bool {
	return currentLine == previousLine
}

func containsSpecialCharacter(line string) bool {
	return strings.ContainsAny(line, "'")
}

func isMinimumLength(word string, min int) bool {
	return len(word) >= min
}
