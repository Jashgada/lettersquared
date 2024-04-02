package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world.")
	input()
}

func input() {
	rows := make([][]string, 4)
	for i := range rows {
		rows[i] = make([]string, 3)
	}
	for i := range rows {
		for j := range rows[i] {
			fmt.Println("Enter a value for row ", i, " column ", j)
			fmt.Scan(&rows[i][j])
		}
	}
	fmt.Println(rows)
}
