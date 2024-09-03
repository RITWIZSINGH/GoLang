package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)  //'bufio' is used to read different things like string, slice etc
										//os.Stdin is used to tell the 'reader' where to read from
	fmt.Println("Enter the rating:")	// every println's argument has a '\n' at the end of the string

	//comma ok || error ok syntax
	input, _ := reader.ReadString('\n') // here ReadString will read until the last i.e., '\n' comes
	//Go doesn't have try or catch blocks , it directly stores the value in the input if there is to store somethin otherwise it will return 'err', if we dont want to use 'err' then just replace it with '_'.

	fmt.Println("Thanks for rating: ", input)


}
