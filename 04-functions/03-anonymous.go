package main

import "fmt"

func main() {

	func() {
		fmt.Println("Hi there!")
	}()

	func(userName string) {
		fmt.Printf("Hi %s, Have a nice day\n", userName)
	}("Magesh")

	msg := func(userName string) string {
		return fmt.Sprintf("Hi %s, Have a nice day\n", userName)
	}("Suresh")
	fmt.Print(msg)

	q, _ := func(x, y int) (quotient, remainder int) {
		quotient = x / y
		remainder = x % y
		return
	}(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d\n", q)
}
