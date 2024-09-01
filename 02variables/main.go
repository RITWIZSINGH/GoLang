package main

import "fmt"

func main() {
	 
	var username string = "Ritwiz"//normal string
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n",username)

	var isLoggedIn bool = true//bool value can be true or false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n",isLoggedIn)

	var smallVal uint8 = 255//uint 8's range is (0 to 2^8) similarly unit16's range is (0 to 2^16)
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n",smallVal)

	var smallFloat float64 = 255.455548841622//more precision than float 32
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n",smallFloat)

	//default values and some aliases
	var anotherVar int
	fmt.Println(anotherVar)
	fmt.Printf("Variable is of type: %T \n",anotherVar)

	var anotherString string
	fmt.Println(anotherString)
	fmt.Printf("Variable is of type: %T \n",anotherString)

	


}