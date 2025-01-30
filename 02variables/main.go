package main

import "fmt"

const LoginToken string = "jsonwebtoken"

func main() {

	// string

	var username string = "harris"
	fmt.Println(username)
	fmt.Printf("Variable of the typee is  : %T\n", username)

	// boolean

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable of the type is : %T\n", isLoggedIn)

	// smaller values

	var smallvalue uint8 = 215
	fmt.Println(smallvalue)
	fmt.Printf("Variable of the type is : %T\n", smallvalue)

	// int takes values from (-127 to 127 ) whereas uint8 values are ranges from 0 to 255

	// float values

	var floatvalue float64 = 234455.4547589999969
	fmt.Println(floatvalue)
	fmt.Printf("Variable of the type is ; %T\n", floatvalue)

	// 32 bit does not give me the precise value so preferreed the 64 bit

	// default valus and some aliases ( in go lang if we do not allocate any data it does not store the garbage value collector )

	var annotherVariable int
	fmt.Println(annotherVariable)
	fmt.Printf("Variable of the type is ; %T\n", annotherVariable)

	//implicit type

	var website = "go.dev"
	fmt.Println(website)

	// no var style

	numberOfusers := 300000
	fmt.Println(numberOfusers)

}
