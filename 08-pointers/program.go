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

	x := 100
	fmt.Println("Before incrementing, x = ", x)
	increment(x)
	fmt.Println("After incrementing, x = ", x)

	n1, n2 := 100, 200
	fmt.Printf("Before swapping, n1 = %d and n2 = %d\n", n1, n2)
	swap(n1, n2)
	fmt.Printf("After swapping, n1 = %d and n2 = %d\n", n1, n2)
}

func increment(no int) /* do not return */ {
	no++
}

func swap(n1, n2 int) /* do not return */ {
	/*  */
}
