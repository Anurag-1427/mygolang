package main

import "fmt"

const LoginToken string = "ghabbhhjd" //Public keyword

func main() {
	var username string = "hitesh"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.45544511254451885
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)
	//exercise
	var anotherVariableString string
	fmt.Println(anotherVariableString)
	fmt.Printf("Variable is of type: %T \n", anotherVariableString)

	//implicit type of declaring the variable
	var website = "learncodeonline.in"
	fmt.Println(website)

	//no var style
	numberofUser := 300000
	fmt.Println(numberofUser)

	fmt.Println((LoginToken))
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
