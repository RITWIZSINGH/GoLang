package main

import "fmt"

func main() {

	fmt.Println("Hello Pointers")

	var ptr *int // this is the type of declaration for pointers where u r creating a pointer in GO. It says ptr is a pointer to memory address which stores integer value. *string etc is also valid.

	fmt.Println("Value of the pointer is", ptr) 

	myNumber := 23

	var myptr = &myNumber // Here notice that '*' sign is not used but instead '&' is used to reference to the var myNumber

	fmt.Println("Value of the pointer myptr is ", myptr) 

}