package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("func as value and closure")

	// func can be treated as variable as well by passing the refernce of the actually fucn to a identifier

	hypo := func(x, y float64) float64 {
		return (math.Sqrt(x*x + y*y))
	}

	fmt.Println(hypo(12, 5))

	fmt.Println(compute(hypo)) // func passed as an arguement to another func

	pos, neg := adder(),adder() // adder in itself doesnt need any arguements, but it returns an anonymus function thats why we are catching the refernce of that returning func in POS and NEG
	fmt.Println(pos(2), neg(-2)) // now as pos has the refernce of the outer scope var SUM of adder , whenever we call the pos func we got the same SUM var updated at every call
	fmt.Println(pos(2), neg(-2))
	fmt.Println(pos(2), neg(-2))

	// note we called adder() only 2 times, 1 for pos and 1 for neg i.e only 2 sum variables are formed and get included in their respective closures
	// calling pos and neg over and over dont create new sum vars, it is same as if we defined a global var VAR1 and access it from a function, the value of VAR1 remains same and accessable within the whole program.
	// Simply, calling pos and neg over and over dont create new SUM  
}

func adder() func(int) int { // adder is a function which has no parameters but it returns a function which takes 1 int arguement and returns 1 int
	sum := 0
	fmt.Println("outer ",sum)
	return func(x int) int {
		fmt.Println("inner ",sum)
		sum += x // due to lexical scoping , this anonymus func has the access of "SUM" var of the outer scope
		return sum
	}
}

func compute(myFunc func(float64, float64) float64) float64 { // this compute func accepts a func myFunc with infact accpets 2 float64 values and return a float64 value and compute also return a float64 value
	// we can also write myFunc func(x,y float64) as well
	result := myFunc(3, 4)
	return result
}

