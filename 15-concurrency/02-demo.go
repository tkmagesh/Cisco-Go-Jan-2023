package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main started")
	go f1() // scheduling the function to be executed through the scheduler
	f2()
	// time.Sleep(500 * time.Millisecond) // blocking the execution of the main function and there by allowing the scheduler to look for other scheduled goroutines and execute them (NOT ADVISABLE)

	fmt.Scanln() // DO NOT DO THIS AS WELL
	fmt.Println("main completed")
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(2 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
