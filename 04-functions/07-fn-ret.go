/* higher order functions - functions as return values */

package main

import "fmt"

func main() {
	fn := getFn()
	fn()
	op := getOperation(1)
	fmt.Println(op(100, 200))

	op = getOperation(2)
	fmt.Println(op(100, 200))
}

func getFn() func() {
	return func() {
		fmt.Println("fn invoked")
	}
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func getOperation(opId int) func(int, int) int {
	if opId == 1 {
		return add
	}
	return subtract
}
