package main

import (
	"fmt"
	"os"

	"golang.org/x/text/message"
)

func main() {
	/*b := []int{1, 2, 3, 4,}

	fmt.Println(b)
	//fmt.Println(len(b)
	fmt.Println(b[4])*/

	

	// fatal error if stops execution
	dbconnection := true
	if !dbconnection{
		fatalerror("Can not connect with db")
	}

	
	pass := "1254687923212"

	fmt.Println(len(pass))
	if len(pass)<9{
		fmt.Println(fmt.Errorf("passward not match"))
	}else{
		savetodb(pass)
	}

}

func fatalerror(message string){
	fmt.Printf("some error which is causeing fatal error, messagwe =%v \n",message)
	os.Exit(1)
}

func savetodb(data string){
	fmt.Println("Recoard success  =",data)
}