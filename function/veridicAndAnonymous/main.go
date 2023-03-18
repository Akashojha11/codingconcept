package main

import (
	"fmt"

	
)

func main() {


	func (a int)int{// anonymous function
		fmt.Println("Anonymous function is running")
		return 0

	}(5)
	fmt.Println("Sum = ",add(66,54 ,6))
	add(4558,55877,69771,7891,895210,)

}

func add( b ...int) int {
	sum := 0
	for i/*_ also use*/ , val :=range b{  // _ (under score) is use becouse slice and array two value return indes and their value
		fmt.Println("i ", i,";val ",val)

		//sum += val
		sum =sum +val        
	}
	return sum

}