package main

import "fmt"

// jwtToken := 30000000.0 we cant use walrus Op here it is not allowed in global scope
// we are allowed to use := in a method only

const LoginToken string = "eghj2345"
// note using captial "L" or any capital letter in a var/method declaration 
// we are declaring that var/method as public 
// same as some -> public const LoginToken =" blahblah"
// public declaration mean we can even access this var in any file under same dir

func main() {
	var username string = "ankit panchal"
	fmt.Printf("Variable is of Type : %T \n", username)
	fmt.Printf("here %s \n",username)

	var isLoggedIn bool = true;
	fmt.Printf("variable is of type %T \n",isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	
	var smallFloat float64 = 255.44444444
 	fmt.Println(smallFloat)

	var anotherVar int
	// anotherVar =10
	fmt.Println(anotherVar)
	fmt.Printf("varibale type is : %T \n",anotherVar)
	
	var emptyStr string
	fmt.Print(emptyStr)
	fmt.Printf("variable type is %T \n",emptyStr)

	var website = "https://www.google.com" //implicit coversion
	fmt.Println(website)

	noOfUser := 300000
	fmt.Println(noOfUser)

	fmt.Println(LoginToken)
}
