package main

import ("fmt")

func main(){
	fmt.Println("Pointers")
	
	// var one int;

	// var ptr *int = &one

	// *ptr = 2
	// fmt.Println("value of ptr is: ",*ptr)

	n := 2
	ptr := &n //using walrus operator to hold address

	fmt.Println("address and value are: ",ptr,*ptr) //*ptr is the deRefrencing

	fmt.Printf("type of ptr and *ptr are %T %T",ptr,*ptr)
	
}