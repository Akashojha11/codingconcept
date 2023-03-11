package main

import "fmt"

func main() {

	event := "Ruk ja Bhai Aram Kar le"

	if event =="Open Door" { //13==10 
		fmt.Println("Hello Team.... Gate Khlo = ", event)

	}else if event =="Close Door" { //11 == 13
		fmt.Println("Gate Band Karo ", event)
	}else{
		fmt.Println("Pta nhi ky kar rha hai = ", event)
	}

	num := 10

	if num == 0|| num>0 { //13==10 
		fmt.Println("Hello Team.... Gate Khlo = ", num)

	}else if num < 0 { //11 == 13
		fmt.Println("Gate Band Karo ", num)
	}else{
		fmt.Println("Pta nhi ky kar rha hai = ", num)
	}
} 