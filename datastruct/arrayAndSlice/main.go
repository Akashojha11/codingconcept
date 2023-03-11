package main

import "fmt"

func main() {
	/*a:=1
	b:=2
	c:=3
	//or
	var p,q,r int*/

	var list [10]int
	list[0] = 10
	list[1] = 11
	list[2] = 10
	list[3] = 11
	list[4] = 10
	list[5] = 11
	list[6] = 10
	list[7] = 11
	list[8] = 10
	list[9] = 10

	fmt.Println("list = ",list)

	list[5]=2054
	fmt.Println("list = ",list)

	var name [10]string
	name[0] = "10"
	name[1] = "11"
	name[2] = "10"
	name[3] = "11"
	name[4] = "10"
	name[5] = "11"
	name[6] = "10"
	name[7] = "11"
	name[8] = "10"
	name[9] = "10"

	fmt.Println("name = ",name)

	var akash [10]interface{}   //interface collect all type data
	akash[0] = "atharv"
	akash[1] = "ojha"
	akash[2] = 10           // like this int value
	akash[3] = false        // bool valuec
	akash[4] = "aditya"
	akash[5] = "Shaya"
	akash[6] = "ram"

	fmt.Println("akash = ", akash)
	// Slice array and slice main and one differance 
	//arrrey size[10] but Slice [] no size 
	//we are free to increase and decrease a/q to us

	var students []string

	students = append(students, "akash")  // this is method to add value in slice
	students = append(students, "atharv")

	fmt.Println("students = ", students[1],students[0])
	fmt.Println("students = ", students)
	


}