package main

import "fmt"

func main(){

	var arr[4]string
	arr[0] = "ankit"
	arr[1]= "panchal"
	arr[3] = "BUSY"

	fmt.Println("arr is: ",arr)
	fmt.Println("length of arr is: ",len(arr))
	
	var arr2 = [5]string{"hello","there","how are you?"}
	
	fmt.Println("arr2 is: ",arr2)
	fmt.Println("length of arr2 is: ",len(arr2))
	fmt.Printf("type of arr2 is %T : ",arr2)	

	arr3 := [5]int{} // all zeros
	
	fmt.Println("arr3 is: ",arr3)

}
