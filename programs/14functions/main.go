package main

import "fmt"

func main() {
	fmt.Println("functions")
	greet()

	// func greet2(){
	// 	fmt.Println("hello hello")
	// } can't define func inside a function

	result,_ := adder(3,5)

	fmt.Println("result is:",result)
	
	result = proAdder(1,2,3,4,5,6,7,8,9,10)
	fmt.Println("result is:",result)

}

func adder(x,y int) (int,string){ // when returning more than 1 value then wrap them in ()
	return x+y ,"we can return 2 datatypes like in comma,Ok"
}
// variadic function , ellipsis ->  (...) 
func proAdder(vals ...int) int{ // in the case where we rnt sure how much values to pass the function we use "...T" this makes a slice of all the passed values to the func
	total :=0
	for _,val :=range vals{ // vals is a slice
		total+=val
	}
	return total
}

func greet2(){
	fmt.Println("hello hello")
} 

func greet(){
	fmt.Println("hello there")
}