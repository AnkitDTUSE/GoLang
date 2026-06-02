package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to GO"
	fmt.Println(welcome)

	reader:= bufio.NewReader(os.Stdin)
	fmt.Println("enter the price: ")
	
	// comma OK || err OK

	input ,_ := reader.ReadString('\n')
	fmt.Println("price is,",input)
	fmt.Printf("type of input %T \n",input)

	_,err := reader.ReadString('\n')
	fmt.Println("err is:",err)
	fmt.Printf("err type is %T:",err)

}