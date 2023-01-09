package main

import "fmt"

func main() {
	var c1 complex64 = 4 + 5i
	fmt.Println(c1)

	var c2 complex64 = 8 + 9i
	result := c1 + c2
	fmt.Println(result)

	fmt.Println("real(result) =", real(result))
	fmt.Println("imag(result) =", imag(result))
}
