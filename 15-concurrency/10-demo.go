package main

import (
	"fmt"
	"time"
)

func main() {
	count := 10
	ch := genFib(count)
	for i := 0; i < count; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("Done")
}

func genFib(count int) chan int {
	ch := make(chan int)
	go func() {
		x := 0
		y := 1
		for i := 0; i < count; i++ {
			ch <- x
			x, y = y, x+y
			time.Sleep(500 * time.Millisecond)
		}
	}()
	return ch
}
