package main

import (
	"fmt"
	"sync"
)

func main() {
	mutex := sync.Mutex{}
	account := 0
	wg := sync.WaitGroup{}

	for i := 1; i <= 10000000; i++ {
		wg.Add(1)
		go increment(&account, &wg, &mutex)

	}
	wg.Wait()
	fmt.Println(account)

}

func increment(bal *int, wg *sync.WaitGroup, tala *sync.Mutex) {
	defer wg.Done()
	tala.Lock()
	*bal++
	tala.Unlock()

}
