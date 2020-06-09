package main

import (
	"fmt"
	"runtime"
)

func main() {
	n := 30

	jobs := make(chan int, n)
	results := make(chan int, n)

	for i := 0; i < runtime.NumCPU(); i++ {
		go worker(jobs, results)
	}

	for i := 0; i < n; i++ {
		jobs <- i
	}

	for j := 0; j < n; j++ {
		fmt.Println(<-results)
	}
}

func worker(in <-chan int, out chan<- int) {
	for n := range in {
		out <- fibonacci(n)
	}
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}
