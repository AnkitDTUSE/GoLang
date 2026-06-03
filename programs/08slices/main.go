package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("slices")

	var newSlice = []int{}
	fmt.Println(newSlice)

	mySlice := []string{"apple", "google", "BUSY"} // can add as many value we want
	fmt.Printf("type of myslice is %T:\n ", mySlice)

	mySlice = append(mySlice, "amazon") // add element at the end of slice
	mySlice = append(mySlice, "amazon") // add element at the end of slice
	mySlice = append(mySlice, "amazon") // add element at the end of slice
	fmt.Println("myslice is: ", mySlice)

	a := mySlice[0:4:4]            // low : high : max (max is the exclusive upper bound in the backing array that the resulting slice is allowed to use for its capacity)
	mySlice = append(mySlice[0:5]) // low: high  ("high" element is excluded)
	fmt.Println(mySlice)
	fmt.Println(a)

	//Sort
	sort.Ints(scores) //only availabe for slices
	fmt.Println(scores)
	fmt.Println(scores)
	fmt.Println(sort.IntsAreSorted(scores))

	//Slices as refernces to arrays

	myArr := [5]int{1, 2, 3, 4, 5}

	slice1 := myArr[:3]
	slice2 := myArr[3:]
	//len(s) The length of a slice is the number of elements it contains.
	//cap(s) The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
	fmt.Println(slice1, slice2)
	fmt.Println(len(slice2), cap(slice2))
	fmt.Println(len(slice1), cap(slice1))

	slice1[0] = 0  // this will the actual value in myArr
	slice2[1] = 56 // rule is low+index (low of slice intilization, index is currently changing index)

	fmt.Println(myArr)

	// Slices as literals -> text representations of fixed values written directly into the source code are literals

	// [3]T{} this creates an array BUT []T{} this also creates an arrya but then builds a slice that refernces it

	r := []bool{true, false, true}
	fmt.Println(r)

	s := []struct { // slice of struct
		i int
		x bool
	}{
		{1, true},
		{2, false},
	}

	fmt.Println(s)

	// nil slice
	// zero value of a slice is nil

	var c []int // this makes NIL slice; c:= []int{} this makes notnil empty slice
	fmt.Println(c, len(c), cap(c))

	if c == nil {
		fmt.Println("NIL")
	}

	// deleting an element from an slice using slice

	brr:=[]int{1,2,3,4,5}
	fmt.Println(brr) 

	indexToRemove := 2
	brr = append(brr[:indexToRemove],brr[indexToRemove+1:]...)
	fmt.Println(brr) 
}
