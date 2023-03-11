package main

import (
	"fmt"
	
)

func main() {
	// map are pointer type , less complicity of map than slice
	// map decleration
	m := make(map[string]int)
	// save data in map
	m["xyz"] = 10
	m["pqr"] = 12

	// read data in map
	val , ok :=m["pqr"]   //val =12 and "ok"-->to say that val are exist yes or no
	fmt.Println("val = ", val, "ok = ",ok)

	val , ok =m["xyz"]
	if ok{
		fmt.Println("value found ,val = ",val)
	}else{
		fmt.Println("value not found")
	}
	fmt.Println("val = ", val, "ok = ",ok) 


}