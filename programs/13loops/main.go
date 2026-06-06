package main

import "fmt"

func main() {
	fmt.Println("functions")

	days:=[]string{"monday","tuesday","wednesday","thrusday","friday","saturday","sunday"}

	fmt.Println(days)

	// for d:=0;d<len(days);d++{
	// 	fmt.Println(days[d])	
	// }

	for i,d:= range days{
		fmt.Println(i,d)
	}
	
	// while Loop of GO is "for"
	val:=1

	for val<10{
		fmt.Println(val)
		val++
	}

 
}