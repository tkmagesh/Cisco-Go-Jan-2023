package main

import (
	"fmt"
	"log"
)

func main() {
	/*
		log.Println("Operation started")
		add(100, 200)
		log.Println("Operation completed")

		log.Println("Operation started")
		subtract(100, 200)
		log.Println("Operation completed")
	*/
	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/
	logOperation(100, 200, add)
	logOperation(100, 200, subtract)
}

func logOperation(x, y int, oper func(int, int)) {
	log.Println("Operation started")
	oper(x, y)
	log.Println("Operation completed")
}

func logAdd(x, y int) {
	log.Println("Operation started")
	add(x, y)
	log.Println("Operation completed")
}

func logSubtract(x, y int) {
	log.Println("Operation started")
	subtract(x, y)
	log.Println("Operation completed")
}

//3rd party library
func add(x, y int) {
	fmt.Println("Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Result :", x-y)
}
