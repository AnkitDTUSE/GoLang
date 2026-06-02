package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println("current time: ",presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	createdDate := time.Date(2006,time.November,4,12,10,56,0,time.UTC)
	fmt.Println("created Date is: ",createdDate.Format("02-01-2006 Monday 15:04:05"))
}
