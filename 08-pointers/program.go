package main

import "fmt"

func main() {
	var no int = 100
	// noPtr := &no
	var noPtr *int
	noPtr = &no

	/*
		var f float32 = 200
		noPtr = &f
	*/
	fmt.Println(no, noPtr)

	//deferencing - accessing the value using the address (pointer)
	v := *noPtr
	fmt.Println(v)

	fmt.Println(no == *(&no))
}
