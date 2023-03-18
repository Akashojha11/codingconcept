package main

import "fmt"

func main() {

	f := 9
	g := 6

	addbyone(&f, g)
	fmt.Println("f  ",f, "g  ",g)
}

func addbyone(a *int, b int) {

	*a = *a + 1
	 b = b +1
}