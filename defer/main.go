package main

import "fmt"

func main() {
	for i := 0; i < 4; i++ {
		//fmt.Println(i)
		//defer fmt.Println(i)
		//fmt.Println(i)

		defer fmt.Println("Hello")
		defer fmt.Println("Akash")
	}
}