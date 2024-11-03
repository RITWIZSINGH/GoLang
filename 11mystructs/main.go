package main

import "fmt"

func main() {

	fmt.Println("Welcome to myStructs")

	//This is a initialization of a struct
	ritwiz := User{"Ritwiz","ritwiz@go.dev",true,20}

	fmt.Printf("This is the value of the struct %+v\n", ritwiz)

	fmt.Printf("Name is %v and Email is %v\n",ritwiz.Name,ritwiz.Email)

}

//Structs and Interface are basically replacement of classes in Go they are very userful below is a declaration of a simple struct 
type User struct {
	Name string
	Email string
	Status bool
	Age int
}