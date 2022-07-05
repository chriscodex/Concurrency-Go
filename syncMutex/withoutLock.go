package main

import (
	"fmt"
	"sync"
)

var (
	BALANCE int = 100
)

func Deposit(amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	b := BALANCE
	BALANCE = b + amount
}

func withoutLock() {
	var wg sync.WaitGroup
	for i := 1; i <= 5000; i++ {
		wg.Add(1)
		go Deposit(i*100, &wg)
	}
	wg.Wait()
	fmt.Println(BALANCE)
}

func main() {
	// Results differents if we dont lock
	// If we check the the binary file, we have a warning
	for i := 0; i < 5; i++ {
		withoutLock()
		BALANCE = 100
	}
}
