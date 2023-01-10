package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError error = errors.New("divisor cannot be zero")

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic occurred - ", err)
			return
		}
		fmt.Println("Thank you!")
	}()
	q, r, err := divideClient(100, 0)
	if err != nil {
		fmt.Println("error occured. do something else - ", err)
		return
	}
	fmt.Println(q, r)

}

func divideClient(x, y int) (quotient, remainder int, err error) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("panic occured in divide()")
			err = e.(error)
			return
		}
	}()
	quotient, remainder = divide(x, y)
	return
}

// 3rd party libary
func divide(x, y int) (quotient, remainder int) {
	defer func() {
		fmt.Println("[divide] - deferred")
	}()
	fmt.Println("calculating the quotient")
	if y == 0 {
		panic(DivideByZeroError)
	}
	quotient = x / y
	fmt.Println("calculating the remainder")
	remainder = x % y
	return
}
