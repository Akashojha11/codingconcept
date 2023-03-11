package main

import "fmt"

func main() {

	daysofweek := 10

	switch daysofweek {
	case 1: //10== 1-->false
		{
			fmt.Println("Monday")
		}

	case 2:
		{
			fmt.Println("Tuesday")
		}
	case 3:
		{
			fmt.Println("Wednesday")
		}
	case 4:
		{
			fmt.Println("Thrusday")
		}
	case 5:
		{
			fmt.Println("Friday")
		}
	case 6:
		{
			fmt.Println("Saturday")
		}
	case 7:
		{
			fmt.Println("Sunday")
		}
		default :{
			fmt.Println("Not A Week Day")
		}
		



	}
}