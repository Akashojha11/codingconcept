package main

import "fmt"

func main() {

	fmt.Println(1)

	var i int
	// start point ;end point ; change piont

	for i := -100  ; i<=100    ; i+=1{
			fmt.Println("i = ",i)
	}

	i =10
	condition  := true
	for condition{
		
		if i >=101{
			condition =false
		}else{
			i++
			fmt.Println("dusra second wale hai",i)
		}

		//range
		var list [10]int
		list[0] =10
		list[1] =13
		list[2] =15
		list[3] =106
		list[4] =104
		list[5] =108
		list[6] =109
		list[7] =102
		list[8] =106
		list[9] =107

		for i:= 0; i< len(list); i++{
			fmt.Println(list[i])
		}
		for _ ,j :=range list{
			fmt.Println( "value is " ,j)
		}
	}

	m := make(map[string]int)
	// save data in map
	m["xyz"] = 10
	m["pqr"] = 12

	for key , val :=range m{
		fmt.Println("key ",key,"value ",val)
	}

}