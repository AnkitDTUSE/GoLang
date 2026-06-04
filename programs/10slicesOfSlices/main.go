package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("slices of slices")

	board := [][]string{ 
		{"_", "_", "_"}, // []string{} can also be used but it gives a warning of redundancy
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s \n", strings.Join(board[i], " "))
	}

	// using make function

	sliceOfSlice := make([][]int, 3) // number of rows

	for i := 0; i < len(sliceOfSlice); i++ {
		sliceOfSlice[i] = make([]int, 4) //number of columns
	}

	for i := 0; i < len(sliceOfSlice); i++ {
		for j := 0; j < len(sliceOfSlice[0]); j++ {
			sliceOfSlice[i][j] = i*j + 1
		}
	}

	for i := 0; i < len(sliceOfSlice); i++ {
		for j := 0; j < len(sliceOfSlice[0]); j++ {
			fmt.Printf("%d ", sliceOfSlice[i][j])
		}
		fmt.Println("")
	}

}
