package main

import "fmt"

//Suppose I use the 'walrus operator' outside a method like
//token := 300 (But we can't) walrus operator can be only used inside a method or a main function
//The correct way of declaring a var outside the above cond will be the following:
// var token = 412 // or
// var token2 int = 23

const LoginToken string = "adiubachsbc" // Here the 'L' is capital that means the LoginToken can be accessed anywhere in this program 

func main() {
	 
	var username string = "Ritwiz"//normal string
	fmt.Println(username)//**Interesting Fact** Here the first letter of Println function of 'fmt' package is Capital that means in the package this function was exported publicly. You can directly use this of fmt in your code. There can be many functions starting with a small letter that are helping the Println to print a line.
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

	//Implicit type declaration

	var website = "yourname.in" //Here the lexer has assumed the value of the var 'website' as string
	fmt.Println(website) //This works fine because 
	//now , i cannot write a line like "webiste = 3" as it is now a string due to line34.

	//No var Style Declaration
	noOfUsers := 30000 // Here ":=" is called the 'WALRUS OPERATOR' it automatically assumes the type of the var 	
	fmt.Println(noOfUsers) //without using the keyword var
	//If I change noOfUser := 3000.0 in line39 it will consider a float

	fmt.Println(LoginToken)

}