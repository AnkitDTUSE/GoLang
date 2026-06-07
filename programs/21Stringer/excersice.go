package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

// -> can do this but string in immutable in Go so doing str+= creates a new string every time which is inefficinet
// func (ipS IPAddr) String() string{
// 	str:=""
// 	for _,ip := range ipS{
// 		str+=fmt.Sprintf("%v.",ip)
// 	}
// 	return str
// }

// -> more efficient way to build a string
// func (ip IPAddr) String() string {
//     var b strings.Builder

//     for _, part := range ip {
//         b.WriteString(fmt.Sprintf("%v.", part))
//     }

//     return b.String()
// }

func (ipS IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ipS[0], ipS[1], ipS[2], ipS[3])
}

func StringerApplication() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {

		fmt.Printf("%v: %v\n", name, ip)
	}
}
