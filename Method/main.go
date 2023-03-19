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

type ojha struct{
	Last string
	Name string
	Address string
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

	Hunar := ojha{
		Last: "Tripathi",
		Name: "Hunar",
		Address: "Kahamajpur",
	}
	akash.addUser(5)
	aditya.addUser(6)
	add()
	Hunar.gol()
	fmt.Println("hunar ",Hunar)
	value := Hunar.gola()
	fmt.Println("Ya Value Wale Hai ", value)


}

func add() {}
func (user *User) addUser(a int) {
	user.Contact =8299747213
	fmt.Println("User Wale method char rha hai", user)
}
func (s student) addUser(b int) {
	s.Name ="Atharv"  // its is use to change the value 

	fmt.Println("Student wale hai", s)
}

func (a ojha)gol(){
	a.Name = "Rama"
	fmt.Println("Bhanjai  ",a)
}

func (a *ojha)gola() int{
	//a.Name = "Ram"
	fmt.Println("Bhanjai  ",a)
	return 0
}
