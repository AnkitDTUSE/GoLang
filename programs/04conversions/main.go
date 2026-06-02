package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("hello there")

	reader:=bufio.NewReader(os.Stdin)
	input,_ := reader.ReadString('\n')

	fmt.Println("input is:",input)

	// changedInput := input+1  this throw error as "input" is string and 1 is int

	changedInput, err := strconv.ParseFloat(strings.TrimSpace(input),64);

	if err!=nil{
		fmt.Println(err)
		// or use panic(err) to exit the program
	}else{
		fmt.Println("Added 1 to Input:",changedInput+1)
	}

}
