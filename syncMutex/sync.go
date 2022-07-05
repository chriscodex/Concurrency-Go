package main

import (
	"fmt"
	"sync"
)

var (
	BALANCE int = 100
)

// Without lock
func depositWithoutLock(amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	b := BALANCE
	BALANCE = b + amount
}

func withoutLock() {
	var wg sync.WaitGroup
	for i := 1; i <= 5000; i++ {
		wg.Add(1)
		go depositWithoutLock(i*100, &wg)
	}
	wg.Wait()
	fmt.Println(BALANCE)
}

// With Lock
func depositWithLock(amount int, wg *sync.WaitGroup, lock *sync.Mutex) {
	defer wg.Done()
	lock.Lock()
	b := BALANCE
	BALANCE = b + amount
	lock.Unlock()
}
func withLock() {
	var wg sync.WaitGroup
	var lock sync.Mutex
	for i := 1; i <= 5000; i++ {
		wg.Add(1)
		go depositWithLock(i*100, &wg, &lock)
	}
	wg.Wait()
	fmt.Println(BALANCE)
}

func main() {
	// Results differents if we dont lock
	// If we check the the binary file, we have a warning
	fmt.Println("Results without lock")
	for i := 0; i < 5; i++ {
		withoutLock()
		BALANCE = 100
	}
	fmt.Println("Results with lock")
	// If we check the output, the values doesn't change anymore
	for i := 0; i < 5; i++ {
		withLock()
		BALANCE = 100
	}
}
