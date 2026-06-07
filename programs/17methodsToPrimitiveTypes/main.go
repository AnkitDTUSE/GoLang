package main

import "fmt"

type celcius float64 // we can different different types of a underlying primitve type to add methods on it

type meter float64 //this type of declaration helps to prevent mixing different units accidently
type kilometer float64

func main() {
	temp := celcius(45)
	fmt.Println(temp.toFarhen())

	// now above we define meter and kilometer, both are underlying float64 but as we declared them seperately we cannot directly perfrom any operation between them

	m := meter(5)
	km := kilometer(10)

	// dist = m+km  // this will throw a error 
	fmt.Printf("m: %v\n", m)
	fmt.Printf("km: %v\n", km)
}

func (c celcius) toFarhen() float64 { // but note any changes happen here to "c" not change the temp varibale in the main func  
	return float64(c)*9/5 + 32
}