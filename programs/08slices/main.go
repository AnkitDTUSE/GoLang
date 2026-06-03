package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("slices")
	
	var newSlice = []int{}
	fmt.Println(newSlice)

	mySlice := []string{"apple","google","BUSY"} // can add as many value we want
	fmt.Printf("type of myslice is %T:\n ",mySlice)
	
	mySlice = append(mySlice, "amazon") // add element at the end of slice
	mySlice = append(mySlice, "amazon") // add element at the end of slice
	mySlice = append(mySlice, "amazon") // add element at the end of slice
	fmt.Println("myslice is: ",mySlice)

	a := mySlice[0:4:4] // low : high : max (max is the exclusive upper bound in the backing array that the resulting slice is allowed to use for its capacity)
	mySlice = append(mySlice[0:5]) // low: high  ("high" element is excluded)
	fmt.Println(mySlice)
	fmt.Println(a)

	scores := make([]int,4)
	scores[0]=100
	scores[1]=101
	scores[2]=102
	scores[3]=103

	// scores[4] = 104 //this will throw error

	scores = append(scores,100,200,300) // automatically realloc the memory in the case of append func

	//Sort 
	sort.Ints(scores) //only availabe for slices	
	fmt.Println(scores)
	fmt.Println(scores)
	fmt.Println(sort.IntsAreSorted(scores))

}