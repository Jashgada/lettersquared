package wordprocessor

func IsAValidWord(word string, board [][]string) (bool, error) {
	for pos, letter := range word {
		for row := range board {
			is_letter_in_word, err := letter_in_row(string(letter), board[row])
			if err != nil {
				return false, err
			}
			if is_letter_in_word {
				if pos != len(word)-1 {
					is_next_letter_in_same_row, err := letter_in_row(string(word[pos+1]), board[row])
					if err != nil {
						return false, err
					}
					if is_next_letter_in_same_row {
						return false, nil
					}
				}
			}
		}
	}
	return true, nil
}

func letter_in_row(letter string, row []string) (bool, error) {
	for i := range row {
		if row[i] == letter {
			return true, nil
		}
	}
	return false, nil
}
