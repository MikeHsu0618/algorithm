package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

// classical race condition without sync.mutex lock
// check if race condition via go run  -race .
func main() {
	msg = "Hello World"

	var mutex sync.Mutex

	wg.Add(2)
	go updateMessage("Hello AAA", &mutex)
	go updateMessage("Hello BBB", &mutex)
	wg.Wait()

	fmt.Println(msg)
}

func updateMessage(s string, mutex *sync.Mutex) {
	defer wg.Done()
	defer mutex.Unlock()
	mutex.Lock()
	msg = s
}
