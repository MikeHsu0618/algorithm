package main

import (
	"fmt"
	"math/rand"
	"time"
)

type worker struct {
	id     int
	number int
}

// worker pool 模式
func main() {
	num := 10
	jobs := make(chan worker, num)
	results := make(chan worker, num)

	// 執行了五個 worker
	for i := 1; i <= 5; i++ {
		go work(jobs, results)
	}

	for i := 1; i <= num; i++ {
		worker := worker{id: i}
		jobs <- worker
	}
	close(jobs)

	for i := 1; i <= num; i++ {
		result := <-results
		fmt.Printf("workerId: %d, result: %v \n", result.id, result.number)
	}
}

func work(jobs <-chan worker, results chan<- worker) {
	for job := range jobs {
		fmt.Println("worker", job.id, "start job")
		time.Sleep(time.Duration(rand.Intn(10)+1) * time.Second)
		fmt.Println("worker", job.id, "finished job")
		job.number = rand.Intn(10) + 1
		results <- job
	}
}
