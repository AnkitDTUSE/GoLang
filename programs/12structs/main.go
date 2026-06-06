package main

import "fmt"

func main(){
	// no inheritance in Go lang ; no super no parent 
	fmt.Println("structs")

	ankit := User{"ankit","ap@.com",true,19}
	fmt.Println(ankit)
	fmt.Printf("dets are: %+v \n ",ankit) // %+v give the value of the variables as well as the fields of the Struct
	fmt.Printf("Name is %v, %+v \n ",ankit.Name,ankit) // %+v give the value of the variables as well as the fields of the Struct

}

type User struct{
	Name string
	Email string
	Status bool
	Age int
}