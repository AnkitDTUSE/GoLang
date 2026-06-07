package main

import "fmt"

type any interface{}

func main(){
	fmt.Println("emptyInterface")

	var x any // this variable can take any value passed to it
	x = "hello"
	fmt.Printf("x: %v x :%T \n", x,x)
	
	x = 16
	fmt.Printf("x: %v x :%T \n", x,x)
	
	x = 1.6
	fmt.Printf("x: %v x :%T \n", x,x)
	
	x = true
	fmt.Printf("x: %v x :%T \n", x,x)

	var y interface{} //another way to use empty interface
	y = "GOoooo"
	fmt.Printf("y: %v\n", y)

	var z any
	fmt.Printf("z: %v\n", z) 

	// type ASSERTIONS
	// A type assertion provides access to an interface value's underlying concrete value.

	var str any = "hello"
	s,ok:=str.(string)
	fmt.Printf("s: %v ok: %v\n ", s,ok)

	TypeSwitch()
}

//3. Collections of Mixed Types

// Suppose you want a slice containing different types:

// var values []any

// values = append(values, 10)
// values = append(values, "hello")
// values = append(values, true)