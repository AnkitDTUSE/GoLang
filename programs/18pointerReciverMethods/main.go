package main

import "fmt"

// This is a Go language convenience feature called automatic address-taking for method calls.

func main() {
	// when we use func calls with pass by value, there is no change in the data of passed variable
	// to make sure we passed the original variable to a function or a method (not its copy) we use pass by reference calling

	a := 10
	fmt.Printf("a: %v\n", a)

	changeValue(&a)
	fmt.Printf("a: %v\n", a) // here we can see the value of "a" is changed 

	user := User{"ankit",19} 
	userPtr := &user  // can use this too (pointer to struct)
	fmt.Printf("userPtr: %v\n", userPtr)

	fmt.Printf("user: %v\n", user)
	// userPtr.updateUser("panchal",19)
	user.updateUser("panchal",19)
	fmt.Printf("user: %v\n", user)

	user2:=User{"badarpur",20}
	fmt.Printf("user2: %v\n", user2)
	changeUserName(&user2,"aadrsh") // here we have to pass &user2 , as changeUserName func is not a method to the User struct , it cant automatically retrive the calling object's address
	fmt.Printf("user2: %v\n", user2)
	
}

type User struct{
	Name string
	Age int
}

func (u *User) updateUser(s string,n int){ //method reciving pointer to struct
	u.Name = s // this is a shorthand for the code written just below
	(*u).Name = s // this is also correct but above one is prefferd in GO
	u.Age = n
}

func changeValue(x *int) {
	*x = 15
}

func changeUserName(u *User,s string){
	u.Name = s
}