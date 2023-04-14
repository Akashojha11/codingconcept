package main

import (
	"fmt"
	"sync"
)

func main() {

	// channel bufffering
	ch2 := make(chan int ,5)
	go func ()  {
	ch2 <- 1
	ch2 <- 2	
	ch2 <- 3
	ch2 <- 4
	ch2 <- 5
	//ch2 <- 6
	
	}()
	
	/*for val := range ch2{
		fmt.Println("buffering example wale output hai",val)
		if val==5{
			break
		}
	}*/
	

	var ch chan int = make(chan int)  // make use for memory alocation anymous func ->>func (){}()

	wg :=sync.WaitGroup{}
	 
	go SendNumber(ch)
	go ReciveNumber(ch , &wg)
	defer wg.Wait()
	wg.Add(1)

}

func SendNumber(ch chan int) {

	for i := 0; i <= 5; i++ {
		ch <- i+1

	}
	close(ch)
}

func ReciveNumber(ph chan int ,wg *sync.WaitGroup) {
defer wg.Done()
	for val:= range ph {
		fmt.Println(val)

	}
}
