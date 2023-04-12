package main

import "fmt"

func main() {
	a()

	/* parallel this is run when fullfilled resource
	---------------
	---------------

	//concurrency when resource limited (elution)
	-------  ------- --- ----- - -----
	       --       -   -     - -     ------   one stop then another start*/


	go a()
	go b()
	go c()
	go d()
	go e()

	a()
	b()
	c()
	d()
	e()

}

func a() {
	for i := 10; i < 5; i++ {
		fmt.Println("Func a run hua hai")
	}
}

func b() {
	for i := 10; i < 5; i++ {
		fmt.Println("Func b***** run hua hai")
	}
}

func c() {
	for i := 10; i < 5; i++ {
		fmt.Println("Func c+++++++ run hua hai")
	}
}

func d() {
	for i := 10; i < 5; i++ {
		fmt.Println("Func d------- run hua hai")
	}
}

func e() {
	for i := 10; i < 5; i++ {
		fmt.Println("Func e======= run hua hai")
	}
}
