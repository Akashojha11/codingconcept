package main

import "fmt"

func main() {

	f := 44
	g :=6

	//fmt.Println("a/b =",a/b)
	fmt.Println("add =",add(f,g))
	fmt.Println("sub =",sub(f,g))
	fmt.Println("mul =",mul(f,g))

	divide ,err :=divi(f,g)
	if err== nil{
		fmt.Println("divide =",divide,"  ",err)
	}else{
		fmt.Println("Divide Nhi Hua     ",  err)
	}
	// faling case

	divide ,err =divi(f,0)
	if err== nil{
		fmt.Println("divide =",divide,err)
	}else{
		fmt.Println("Divide Nhi Hua     ",  err)
	}
 //passing case
	divide ,err =divi(f,2)
	if err== nil{
		fmt.Println("divide =",divide,"  ",err)
	}else{
		fmt.Println("Divide Nhi Hua     ",  err)
	}
	fmt.Println("f  ",f,";g  ",g)

}

//calcultor for adding two number
func add(a,b int) int{
	c:= a+ b
	return c

}

func sub(a,b int) int{
	return a-b

}

func mul(a,b int) int{
	c :=a*b
	return c

}

func divi(d,e int)(int, error){
	if e!= 0{
		return d/e , nil
	}else {
		return 0 ,fmt.Errorf("divinding is not possible with zero")
	}


}