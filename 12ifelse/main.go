package main

import "fmt"

func main() {

	fmt.Println("Welcome to ifElse")

	loginCount := 10
	var result string

	//the '{' should be in the same line
	if loginCount > 10 {
		result = "Regular user"
	} else if loginCount < 10{
		result = "Less than 10 count"
	} else {
		result = "Exactly 10 count"
	}

	if 9%2 == 0 {
		fmt.Println("Number is even")
	}else {
		fmt.Println("number is ODD")
	}

	//Special syntax you can initiazlize variables on the go too
	if num:= 3;num<10 {
		fmt.Print("Number is less than 10")
	}else {
		fmt.Print("Number is NOT less than 10")
	}

	fmt.Println(result)
	
}