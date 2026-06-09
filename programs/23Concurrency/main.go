package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	ID     int
	Status string
	mu     sync.Mutex
}

var (
	totalUpdates int
	updateMutex  sync.Mutex
)

func generateOrder(count int) []*Order {
	orders := make([]*Order, count)

	for i := 0; i < count; i++ {
		orders[i] = &Order{ID: i + 1, Status: "pending"}
	}

	return orders
}

func processOrders(orders []*Order) {
	for _, order := range orders {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		fmt.Printf("Processing Order %d\n", order.ID)
	}
}

func updateOrderStatus(order *Order) {
	order.mu.Lock()
	time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
	order.Status = []string{"processing", "shipped", "delivered"}[rand.Intn(3)]
	fmt.Printf("Updated Order %d status %s\n", order.ID, order.Status)
	order.mu.Unlock()

	updateMutex.Lock()
	defer updateMutex.Unlock()
	currupdates := totalUpdates
	time.Sleep(5 * time.Millisecond)
	totalUpdates = currupdates + 1

}

func reportOrderStatus(orders []*Order) {
	time.Sleep(1 * time.Second)

	fmt.Println("------Order Status Report--------")
	for _, order := range orders {
		fmt.Printf("Order %d %s\n", order.ID, order.Status)
	}
	fmt.Println("------Order Status Report--------")

}

func main() {
	fmt.Println("Concurrency")

	orders := generateOrder(20)

	var wg sync.WaitGroup // creating a wait group

	// wg:=sync.WaitGroup{}

	t1 := time.Now()
	// create Go routine by "go" keyword

	wg.Add(5) // as we have 3 go routines below

	// go processOrders(orders)

	// go updateOrderStatus(orders)

	// go reportOrderStatus(orders)

	// func(){}() //IIF in GO

	go func() { // making this IIF a go Routine
		defer wg.Done() // this decerements the wg.Add() by 1
		processOrders(orders)
	}()

	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()
			for _, order := range orders {
				updateOrderStatus(order)
			}
		}()
	}

	go func() {
		defer wg.Done()
		reportOrderStatus(orders)
	}()

	wg.Wait() // this blocks the execution until the wg become zero

	// _,err := fmt.Scanln()

	// if err !=nil{
	// 	fmt.Print(err)
	// }

	elapsed := time.Since(t1)
	fmt.Printf("Program termination time and total updates => %v %v", elapsed, totalUpdates)

}
