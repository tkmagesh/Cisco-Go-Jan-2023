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

	/* More than 1 varibale */
	/*
		var x int
		var y int
		var result int
		var str string
		x = 100
		y = 200
		result = x + y
		str = "Sum of %d and %d is %d\n"
		fmt.Printf(str, x, y, result)
	*/

	/*
		var x, y, result int
		var str string
		x = 100
		y = 200
		result = x + y
		str = "Sum of %d and %d is %d\n"
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y, result int
			str          string
		)
		x = 100
		y = 200
		result = x + y
		str = "Sum of %d and %d is %d\n"
		fmt.Printf(str, x, y, result)
	*/

	//Declaration & initialization
	/*
		var x int = 100
		var y int = 200
		var result int = x + y
		var str string = "Sum of %d and %d is %d\n"
		fmt.Printf(str, x, y, result)
	*/

	/*
		var x, y int = 100, 200
		var result int = x + y
		var str string = "Sum of %d and %d is %d\n"
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y   int    = 100, 200
			result int    = x + y
			str    string = "Sum of %d and %d is %d\n"
		)
		fmt.Printf(str, x, y, result)
	*/

	//type inference
	/*
		var x, y = 100, 200
		var result = x + y
		var str = "Sum of %d and %d is %d\n"
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y   = 100, 200
			result = x + y
			str    = "Sum of %d and %d is %d\n"
		)
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y, str = 100, 200, "Sum of %d and %d is %d\n"
			result    = x + y
		)
		fmt.Printf(str, x, y, result)
	*/

	x, y, str := 100, 200, "Sum of %d and %d is %d\n"
	result := x + y
	fmt.Printf(str, x, y, result)

}
