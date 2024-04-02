package wordprocessor_test

import (
	"lettersquared/wordprocessor"
	"testing"
)

func TestIsAValidWord(t *testing.T) {
	board := [][]string{{"A", "B", "C"}, {"D", "E", "F"}, {"G", "H", "I"}, {"J", "K", "L"}}
	word := "ADGJA"
	expected := false
	actual, err := wordprocessor.IsAValidWord(word, board)
	print(actual)
	if err != nil {
		t.Fatal("Test failed")
	}
	if expected != actual {
		t.Fatalf("Test failed")
	}
}
