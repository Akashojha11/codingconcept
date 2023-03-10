package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int
	a = 10

	var b int16
	b = int16(a)

	fmt.Println(b)

	var c string
	c = fmt.Sprintf("%d", b)

	fmt.Println(c)

	num := "12584"
	numint, err := strconv.Atoi(num)
	fmt.Println("numint = ", numint, "err = ", err)

	numb := "Akash"
	numinti, err := strconv.Atoi(numb)
	fmt.Println("numint = ", numinti, "err = ", err)
}
