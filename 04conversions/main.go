package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("welcome to our app")
	fmt.Println("Please rate us from 1 to 5")

	reader := bufio.NewReader(os.Stdin)//reader to read the input

	input, _ := reader.ReadString('\n') // storing in input and reading till the end '\n'

	fmt.Println("Thanks for the rating: ", input )//printing the statement

	numRating,err := strconv.ParseFloat(input,64)//strconv is a package you to convert string into different datatypes, it takes two arguments
	//first is the string that is needed to be parsed, second is the size in the above case is 64 due to float64
	//numRating,err := strconv.ParseFloat(strings.TrimSpace(input),64)
	if err != nil { //if there is an error
		fmt.Println(err)
	}else{
		fmt.Println("Added 1 to your rating", numRating+1)
	}
	

}
