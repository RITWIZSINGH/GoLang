package main

import "fmt"

func main() {

	fmt.Println("Welcome to class on slices")
	//slices are basically more powerful version of arrays
	var fruitList = []string{"Apple", "Tomato", "Peach"}// this is a slice declaration when u give it a size while initializing, it becomes an array
	fmt.Printf("The type of fruitlist is %T\n",fruitList ) // printing to see what is the type of this slice
	fmt.Println("*********************************************************************")

	//using append func with slicename as the first arg and elements that are to be added
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println("The final fruitlist is ", fruitList)
	fmt.Println("*********************************************************************")

	fruitList = fruitList[1:] // slicing from 1: means all elements after the 1st i.e., leave the 0th and include all . Here include the element at index 1
	fmt.Println("After slicing from 1: the fruitlist is ", fruitList)
	fmt.Println("*********************************************************************")

	fruitList = fruitList[1:3]// slicing from 1:3 means slice from first element till the third element i.e. , 1st and 2nd dont include 3rd
	fmt.Println("After slicing from 1:3 the fruitlist is ", fruitList)
	fmt.Println("*********************************************************************")

	fruitList = fruitList[:3] // slicing from :3 means slice from default that is 0 to 3rd element and dont include the 3rd index element
	fmt.Println("After slicing from :3 the fruitlist is ", fruitList)
	fmt.Println("*********************************************************************")



}
