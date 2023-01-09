package main

import "fmt"

// var msg string

// var msg string = "Hello World!"

// var msg = "Hello World!"

// msg := "Hello World!" // => NOT supported

var i = 100

func main() {
	/*
		var msg string
		msg = "Hello World!"
	*/

	/*
		var msg string = "Hello World!"
	*/

	//type inference
	// var msg = "Hello World!"
	// var msg = 100

	msg := "Hello World!"
	/*
		":=" syntax can be used ONLY in a "function" (NOT at package scope)
	*/
	fmt.Printf("msg=%v, type=%T\n", msg, msg)

	/*
		var i = 100
		i = 200
		fmt.Println(i) // reading the value of "i" and there by "using" it
	*/

}
