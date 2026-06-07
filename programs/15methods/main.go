package main

import "fmt"

func main() {
	ankit:=User{"ankit",20,"ap@.com",true}
	ankit.GetStatus()
	fmt.Println(ankit.Email)
	ankit.NewEmail()
	fmt.Println(ankit.Email)

	// calling func from another file
	MyMethods()
}

type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
}

// methods are functions associated with a Struct

func (u User) GetStatus() { // here we are getting only the copy of the calling Struct variable not the actual address thus any changes happen inside the method to the variable doesnt effect the original var
	fmt.Println(u.Status)
}

func (u User) NewEmail(){
	u.Email = "test@.com"
	fmt.Println("new email of user is:",u.Email)
}