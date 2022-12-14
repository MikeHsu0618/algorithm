package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	words := []string{
		"A",
		"B",
		"C",
		"D",
	}
	wg.Add(len(words))
	for _, word := range words {
		go PrintSomething(word, &wg)
	}
	wg.Wait()
	time.Sleep(2 * time.Second)
	fmt.Println("Done")
}

func PrintSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}
