package main

import "fmt"

//non primitive data type

type Student struct {
	Name       string
	class      int
	RollNumber int
	SADD       Address
}

type Address struct {
	Lane1 string
	Lane2 string
	Post  string
	Dist  string
}

type mobile struct {
	Model string
	Speci Feature
	IMEI  string
}

type Feature struct {
	CamOFdetail Camera
	Battery     string
	Ram         string
	Stroage     string
}

type Camera struct {
	Front string
	Back  string
}

func main() {

	var akash Student
	//var aditya Student

	akash.class = 12
	akash.Name = "Akash ojha"
	akash.SADD.Lane1 = "NH27"
	akash.SADD.Lane2 = "Rajauli"
	akash.SADD.Post = "MAHARAJGANJ"
	akash.SADD.Dist = "BASTI"

	aditya := Student{
		Name:       "Aditya Ojha",
		class:      10,
		RollNumber: 14,
		SADD: Address{
			Dist:  "BASTI",
			Lane1: "NH27",
			Lane2: "RAJAULI",
			Post:  "MAHARAJGANJ",
		},
	}
	var poco mobile
	poco.Model = "F1"
	poco.IMEI = "EMI125S32D1256ED12"
	poco.Speci.Battery = " Battery=4500mah"
	poco.Speci.Ram = "8GB"
	poco.Speci.Stroage = "128GB"
	poco.Speci.CamOFdetail.Back = "108MP"
	poco.Speci.CamOFdetail.Front = "55MP"

	realme := mobile{
		Model: "PRO MAX",
		IMEI:  "EMI587223121SJUHD5",
		Speci: Feature{
			Battery: "50000mah",
			Ram:     "12GB",
			Stroage: "256GB",
			CamOFdetail: Camera{
				Back:  "120MP",
				Front: "50MP",
			},
		},
	}

	//INTERFACE IS ACCEPTED ANY VALUES
	val := 122
	val2 :="65562"
	//val3 := false

	var akashAddress *Student                                    // if any change in pointer variable then 
	akashAddress = &akash                                         //change in original values

	adityaAddress := &aditya

	

	fmt.Println("Akash pointer address = ",*akashAddress)           //* mean give the value of this address
	fmt.Println("Aditya pointer address = ", adityaAddress)         // & mean give the address 


	var interfaceExample interface{}
		interfaceExample =val
		fmt.Println("Interface value =",interfaceExample)

		interfaceExample =val2
		fmt.Println("Interface value =",interfaceExample)

		interfaceExample =false                                 // to save one variable,becouse 
	                                                             //when varible not use in future	
		fmt.Println("Interface value =",interfaceExample)
	



	fmt.Println("Hello akash Team...", akash)
	fmt.Println("Hello akash Team...", aditya)

	fmt.Println("Hello akash Team... POCO F1", poco)
	fmt.Println("Hello akash Team...REALME PRO", realme)
}
