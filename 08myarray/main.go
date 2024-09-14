package main

import "fmt"

func main() {
	fmt.Println("Welcome to array")

	var fruitList [4]string // this is the declaration of an array in go. You have to give a value for size to initialize it.
	
	fruitList[0] = "Apple"
	fruitList[1] = "Peach"
	fruitList[3] = "Orange" // Initiazlization of the array and leaving one empty

	fmt.Println("Fruit List is:", fruitList ) //Notice the output 
	fmt.Println("Size of fruit List is:", len(fruitList) ) // len used to calculate the size of the array 
	fmt.Println("*****************************************************************************************************")

	var veggieList = [3]string {"Potato", "beans", "mushroom"} // declaration and initialization of the array in the same line

	fmt.Println("Veggie List is:", veggieList)
	fmt.Println("Size of Veggie List is:", len(veggieList)) //That's pretty much it that u can do with arrays

}
