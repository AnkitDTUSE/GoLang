package main

import (
	"fmt"
	"math"
)

func wishBday() string{return "happy bday"}

func MyMethods() {
	a := vertex{3, 4}
	fmt.Println(a.absDistance()) // as absDistance is a method assiciated to sturct vertex we can call it via (.) operator on an instant of vertex

	printMyVal(a) // here printMyVal func has a parameter of vertex type, but there is no type defined in its receiver arguements i.e we cant call printMyVal via (.) , it is just a normal func 

}

type vertex struct {
	x, y float64
}

func (v vertex) absDistance() float64 {
	return (math.Sqrt(v.x*v.x + v.y*v.y))
}

func printMyVal(v vertex) {
	fmt.Println(v.x)
	fmt.Println(v.y)
}