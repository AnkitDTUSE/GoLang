package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader:= bufio.NewReader(os.Stdin)
	fmt.Println("enter the price: ")
	
	// comma OK || err OK

	input ,_ := reader.ReadString('\n')
	fmt.Println("thanks for rating,",input)
	fmt.Printf("type of input %T ,",input)

}