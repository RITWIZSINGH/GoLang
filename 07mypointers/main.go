package main

import "fmt"

func main() {

	fmt.Println("Welcome to the class on pointers")

	var ptr *int // this is the type of declaration for pointers where u r creating a pointer in GO. It says ptr is a pointer to memory address which stores integer value. *string etc is also valid.

	fmt.Println("Value of the pointer is", ptr) 
	fmt.Println("*****************************************************************************************************")
	myNumber := 23

	var myptr = &myNumber // Here notice that '*' sign is not used but instead '&' is used to reference to the var myNumber

	fmt.Println("Value of the pointer myptr is ", myptr)  // So here we are printing the value of the pointer which has the memory address of the variable myNumber
	fmt.Println("****************************************************************************************************")
	fmt.Println("Value the variable at the address myptr is ", *myptr)// Here by using '*' we can access the value at the memory location stored inside myptr
	fmt.Println("*****************************************************************************************************")
	*myptr = *myptr + 2 // Here we are saying that whatever is at the memory address of myptr , add '2' in it. the actual value of myNumber has changed not a copy of it .
	fmt.Println("New value is ", myNumber)
}