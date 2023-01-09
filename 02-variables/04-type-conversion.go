package main

import "fmt"

func main() {
	var i int = 100
	var f float64
	f = float64(i) // using the typename like a function for type conversion
	fmt.Println(f)
}
