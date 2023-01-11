package main

import (
	"fmt"
	"time"
)

func main() {

	ch := genFib()

	for {
		if data, ok := <-ch; ok {
			fmt.Println(data)
			continue
		}
		break
	}
	fmt.Println("Done")
}

func genFib() chan int {
	count := 15
	ch := make(chan int)
	go func() {
		x := 0
		y := 1
		for i := 0; i < count; i++ {
			ch <- x
			x, y = y, x+y
			time.Sleep(500 * time.Millisecond)
		}
		close(ch)
	}()
	return ch
}
