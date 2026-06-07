package main

import "fmt"

type User struct{
	Name string
	Age int
}

func (u User) String() string{ // this method implements Stringer Interface as it has String()
	// fmt.Sprintf is a function
	// used to format a string and 
	// return it as a value rather than printing it directly to the console
	return fmt.Sprintf("Name is %v Age is %v\n",u.Name,u.Age)
}

func main(){
	fmt.Println("STRINGERS")

	a:= User{"ankit",19}
	b:= User{"ansh",19}
	fmt.Println(a,b)

	StringerApplication()
}