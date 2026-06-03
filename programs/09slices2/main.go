package main

import "fmt"

func main() {
	// creating slice using make

	// The make function allocates a zeroed array and returns a slice that refers to that array:
	scores := make([]int, 4)
	scores[0] = 100
	scores[1] = 101
	scores[2] = 102
	scores[3] = 103

	// scores[4] = 104 //this will throw error

	scores = append(scores, 100, 200, 300) // automatically realloc the memory in the case of append func
	fmt.Println(scores)

	// To specify a capacity, pass a third argument to make

	b:=make([]int,0,5) // to ask

	b[0]=100
	b[1]=101
	b[2]=102
	b[3]=103
	b[4]=104

	fmt.Println(b)
}