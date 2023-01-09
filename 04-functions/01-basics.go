package main

import "fmt"

func main() {
	sayHi()
	greet("Magesh")
	fmt.Print(getGreetMsg("Suresh"))
	fmt.Println(add(100, 200))
	// fmt.Println(divide(100, 7))
	/*
		q, r := divide(100, 7)
		fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)
	*/

	// printing ONLY the quotient
	// "_" => placeholder for a variable that we dont want to use
	q, _ := divide(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d\n", q)
}

/* no parameters, no results */
func sayHi() {
	fmt.Println("Hi there!")
}

/* 1 parameter */
func greet(userName string) {
	fmt.Printf("Hi %s, Have a nice day\n", userName)
}

/* 1 parameter, 1 result */
func getGreetMsg(userName string) string {
	return fmt.Sprintf("Hi %s, Have a nice day\n", userName)
}

/* 2 parameters, 1 result */
/*
func add(x int, y int) int {
	return x + y
}
*/

func add(x, y int) int {
	return x + y
}

/* 2 parameter, 2 results */
/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

func divide(x, y int) (quotient, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
