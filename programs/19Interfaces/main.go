package main

import "fmt"

type paymenter interface{
	pay(float64) // this interface help us to directly use those TYPES who has pay(float64) as their method 
}

type payment struct{
	gateway paymenter
}

func (p payment) makePayment(amnt float64){
	p.gateway.pay(amnt)
}

type paytm struct{}

func (p *paytm) pay(amnt float64){
	
	fmt.Println(p == nil)

	fmt.Println("making payment using paytm of rupees:",amnt)

}
type googlePay struct{}

func (p googlePay) pay(amnt float64){
	fmt.Println("making payment using googlePay of rupees:",amnt)

}

func main(){
	fmt.Println("Interfaces")

	newPayment := payment{gateway: &paytm{},} // now here we have two gateways to choose between but we to hardcode the gateway to choose one
	newPayment.makePayment(100) 

	fmt.Printf("newPayment: %T\n", newPayment)

	newPayment = payment{gateway: googlePay{}}
	newPayment.makePayment(1000)


	// a:=paymenter  this is same as a:=int (totally invalid) as payementer is a type

	var a paymenter // here we first make a interface variable
	
	a = &paytm{} // then initialize it with a compatible TYPE (a type which has the method signature which is written in paymenter)

	a.pay(100) 

	//Interface Value

	fmt.Printf("value 'a' : %v type 'a' : %T\n", a,a)

}

// package main

// import (
// 	"fmt"
// 	"math"
// )

// type Abser interface {
// 	Abs() float64
// }

// func main() {
// 	var a Abser
// 	f := MyFloat(-math.Sqrt2)
// 	v := Vertex{3, 4}

// 	a = f  // a MyFloat implements Abser
// 	a = &v // a *Vertex implements Abser

// 	// In the following line, v is a Vertex (not *Vertex)
// 	// and does NOT implement Abser.
// 	a = v

// 	fmt.Println(a.Abs())
// }

// type MyFloat float64

// func (f MyFloat) Abs() float64 {
// 	if f < 0 {
// 		return float64(-f)
// 	}
// 	return float64(f)
// }

// type Vertex struct {
// 	X, Y float64
// }

// func (v *Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }
