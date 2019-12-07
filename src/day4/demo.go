package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "process job", j)
		time.Sleep(time.Second)
		jobs <-j
	}
}
func main() {
	jobs := make(chan int, 100)
	results := make(chan int,100)
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	// <-results
	close(jobs)
	for a := 1; a <= 10; a++ {
		<-results
	}
	
}
