package main

import (
	"flag"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(1) //how may CPUs to be used by the scheduler
	rand.Seed(7)
	count := flag.Int("count", 0, "Number of goroutines to spawn")
	flag.Parse()
	fmt.Printf("main started. Starting %d goroutines... Hit ENTER to start\n", *count)
	fmt.Scanln()
	for i := 0; i < *count; i++ {
		wg.Add(1)
		go fn(i + 1)
	}
	wg.Wait()
	fmt.Println("main completed.. Hit ENTER to shutdown...")
	fmt.Scanln()
}

func fn(id int) {
	fmt.Printf("fn[%d] started\n", id)
	time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
	fmt.Printf("fn[%d] completed\n", id)
	wg.Done() // decrement the counter by 1
}
