package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Welcome to class on slices")
	//slices are flexible, and more dynamic view of an array. They can grow and shrink in size but it still points to an underlying array
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

	highScores := make([]int, 4) // here we are declaring and initializing a slice using the make function. 'make' func not only declares the slice but it also initializes the slice with non - zeroed storage. Non zeroed storage refers that you can store data in the memory allocated.

	highScores[0] = 234 
	highScores[1] = 454
	highScores[2] = 323
	highScores[3] = 164

	fmt.Println("The highscores without sorting are", highScores) // this will simply showcase the element in the slice highScores
	fmt.Println("*********************************************************************")

	highScores = append(highScores, 555,846,743) // Here we are just appending some values in the slice

	fmt.Println("After appending some values the new highscores without sorting are", highScores)
	fmt.Println("*********************************************************************")

	sort.Ints(highScores) // 'Ints' func of the 'sort' package sorts the element of int slice in increasing order.

	fmt.Println("After sorthing the elements are ", highScores)
	fmt.Println("*********************************************************************")

	fmt.Println("Are the elements in highScores slice sorted?", sort.IntsAreSorted(highScores)) // IntsAreSorted func returns true if the elements are sorted

	

	



}
