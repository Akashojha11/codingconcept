package main

import "fmt"

type User struct {
	Name    string
	Address string
	Contact int
}

type student struct {
	Name   string
	class  int
	RollNo int
}

type UserOperator interface{
	addUser()
}


func main() {

	var akash User
	akash.Name = "Akash Ojha"
	akash.Address = "Basti"
	akash.Contact = 9161199921
	/*var aditya student
	aditya.Name = "Aditya Ojha"
	aditya.class = 12
	aditya.RollNo = 10*/
	aditya := student{
		Name:   "Aditya Ojha",
		class:  12,
		RollNo: 10,
	}

	
	//akash.addUser(5)
	//aditya.addUser(6)
	//add()

	var userOper UserOperator
	userOper = akash
	userOper.addUser()

	userOper = aditya
	userOper.addUser()
	
}
//func add() {}
func (user User) addUser() {
	user.Contact =8299747213
	fmt.Println("User Wale method char rha hai", user)
}
func (s student) addUser() {
	s.Name ="Atharv"  // its is use to change the value 

	fmt.Println("Student wale hai             ", s)
}

