package main

import (
	"fmt"
)

func main() {
<<<<<<< HEAD

	fmt.Println("Welcome to Go Maps")

	//This is the declaration for a map in 
	languages := make(map[string]string)

	//This is the initialization Eg: ["JS"], ["RB"] they are the keys
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"


	fmt.Println("This is the languages map", languages)
	fmt.Println("JS shorts for" , languages["JS"])

	//THis is how you can delete a key value pair from a map
	delete(languages,"RB")

	fmt.Println("This is the languages map", languages)

	//Loops 

	for key,value := range languages {
		fmt.Printf("For key %v, the value is %v\n", key,value)
	}




	

=======
fmt.Println("Welcome to Go Maps")
>>>>>>> 225ab33c6c6e72094b766dd028c86b14578e6051
}
