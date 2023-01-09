package main

import (
	"fmt"
	"log"
	"time"
)

type OperationFn func(int, int)

func main() {
	// logging
	/*
		logAdd := getLogOperation(add)
		logAdd(100, 200)

		logSubtract := getLogOperation(subtract)
		logSubtract(100, 200)
	*/

	// profiling
	/*
		profiledAdd := getProfileOperation(add)
		profiledAdd(100, 200)

		profiledSubtract := getProfileOperation(subtract)
		profiledSubtract(100, 200)
	*/

	// logging & profiling
	/*
		logAdd := getLogOperation(add)
		profiledLogAdd := getProfileOperation(logAdd)
		profiledLogAdd(100, 200)

		logSubtract := getLogOperation(subtract)
		profiledLogSubtract := getProfileOperation(logSubtract)
		profiledLogSubtract(100, 200)
	*/

	/*
		profiledLogAdd := getProfileOperation(getLogOperation(add))
		profiledLogAdd(100, 200)

		profiledLogSubtract := getProfileOperation(getLogOperation(subtract))
		profiledLogSubtract(100, 200)
	*/

	getProfileOperation(getLogOperation(add))(100, 200)
	getProfileOperation(getLogOperation(subtract))(100, 200)

}

func getProfileOperation(oper OperationFn) OperationFn {
	return func(x, y int) {
		start := time.Now()
		oper(x, y)
		elapsed := time.Since(start)
		fmt.Println("Elapsed :", elapsed)
	}
}

func getLogOperation(oper OperationFn) OperationFn {
	return func(x, y int) {
		log.Println("Operation started")
		oper(x, y)
		log.Println("Operation completed")
	}
}

//3rd party library
func add(x, y int) {
	fmt.Println("Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Result :", x-y)
}
