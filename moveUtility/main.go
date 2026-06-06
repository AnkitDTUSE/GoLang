package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	source:=args[1]
	// dest:=args[2]
	
	directory,err:= os.ReadDir(source)
	if err!=nil{
		fmt.Print("some error is there")
	}
	
	paths:=[]string{}

	for i,val :=range directory{
		fmt.Println(i)
		paths = append(paths, val.Name())

	}
	fmt.Println(paths)

	

	// err := os.Rename(source,dest)

	// if err!=nil{
	// 	fmt.Println("error during moving ",err)
	// }else{
	// 	fmt.Println("movement done")
	// }
	
}