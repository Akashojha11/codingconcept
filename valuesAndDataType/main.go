package main

import "fmt"

func main() {
	var number int
	number = 12
	number = 5685

	var addresofNumber *int
	addresofNumber=&number

	var decNum float32
	decNum = 123.55
	decNum =635.6565
	var addressofdec *float32
	addressofdec =&decNum

	numb := 5665
	fmt.Println(numb)

	num4 :=number // it is use when number type not known then use short hand operator

	var flags bool
	flags = true
	flags=false

	var name string
	name = "Coding concepts"
	name ="akash ojha"

	fmt.Printf("number = %d,decNum = %f,flags = %v,name = %s \n", number, decNum, flags, name)
	fmt.Printf("addressofNumber = %d,addressofdecNum = %f,flags = %v,name = %s \n", *addresofNumber, *addressofdec, flags, name)
	



}
