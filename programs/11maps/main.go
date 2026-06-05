package main

import (
	"fmt"
	"golang.org/x/tour/wc"
)

func main() {
	fmt.Println("maps")

	states:=make(map[int]string) //map[key]value

	states[1] ="delhi"
	states[2] ="UP"
	states[3] ="UK"
	states[4] ="himachal"

	fmt.Println(states)
	fmt.Println("value at 1: ",states[1])
	
	delete(states,1) //to delete a key
	fmt.Println(states)

	// range -> The range form of the for loop iterates over a slice or map.

	for key,value:= range states{ // can use _,value pr key,_ to avoid the underscored variable
		fmt.Printf("key %v value %s \n",key,value)
	}

	for key:= range states{ // can totally omit the other var
		fmt.Printf("key %v ",key)
	}
	fmt.Println()

	m :=map[int]string{
		1:"j&K",
		2:"haryana",
	}
	fmt.Println(m)

	//retrving value from map

	value:= m[1]
	fmt.Println(value)

	// Test that a key is present with a two-value assignment:

	_,ok :=m[5]
	fmt.Println(ok)

	wc.Test(WordCount) // need to run -> go run . (this will trigger all files under the same package) or go run main.go excersice.go

	// note WordCount can be named as wordCount as main and wordCount are now under same dir , but if we need to use wordCount outside the current package then name it WordCount
	 
}