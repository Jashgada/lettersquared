package databaseinitializer_test

import (
	"lettersquared/databaseinitializer"
	"testing"
)

func TestIsValidWordReturnsFalse(t *testing.T) {
	word := "ADGJA"
	wordMap := make(map[string]int)
	wordMap["A"] = 1
	wordMap["B"] = 1
	expected := false
	actual := databaseinitializer.IsValidWord(word, wordMap)
	if expected != actual {
		t.Fatalf("Test failed")
	}
}

func TestIsValidWordReturnsTrue(t *testing.T) {
	word := "AB"
	wordMap := make(map[string]int)
	wordMap["A"] = 1
	wordMap["B"] = 1
	wordMap["C"] = 1
	expected := true
	actual := databaseinitializer.IsValidWord(word, wordMap)
	if expected != actual {
		t.Fatalf("Test failed")
	}
}
