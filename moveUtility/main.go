package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	source:=args[1]
	dest:=args[2]

	err := os.Rename(source,dest)

	if err!=nil{
		fmt.Println("error during moving ",err)
	}else{
		fmt.Println("movement done")
	}
	
}