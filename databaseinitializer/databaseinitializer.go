package databaseinitializer

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func FetchDataFromCSV(words chan string, filename string, wordMap map[string]int) {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error reading records")
	}

	for _, eachrecord := range records {
		if IsValidWord(eachrecord[0], wordMap) {
			words <- eachrecord[0]
		}
	}
}

func CleanDatabase(wordMap map[string]int) {
	words := make(chan string)
	FetchDataFromCSV(words, "database.csv", wordMap)
	close(words)
}

func IsValidWord(word string, wordMap map[string]int) bool {
	for _, letter := range word {
		if _, ok := wordMap[string(letter)]; !ok {
			return false
		}
	}
	return true
}
