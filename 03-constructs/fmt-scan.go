package main

import "fmt"

func main() {
	/*
		var name string
		fmt.Println("Enter the name:")
		fmt.Scanln(&name)
		var msgTemplate string = "Hi %s, Have a nice day!"
		var greetMsg = fmt.Sprintf(msgTemplate, name)
		fmt.Println(greetMsg)
	*/

	var firstName, lastName string
	fmt.Println("Enter the name[firstName lastName]:")
	fmt.Scanln(&firstName, &lastName)
	var msgTemplate string = "Hi %s, %s, Have a nice day!"
	var greetMsg = fmt.Sprintf(msgTemplate, lastName, firstName)
	fmt.Println(greetMsg)

	var no int
	fmt.Println("Enter a no:")
	fmt.Scanln(&no)
	if no%2 == 0 {
		fmt.Printf("%d is an even number\n", no)
	} else {
		fmt.Printf("%d is an odd number\n", no)
	}
}
