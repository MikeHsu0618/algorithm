package main

import (
	"fmt"
	"sync"
	"time"
)

var msg string

func updateMessage(s string) {
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func main() {

	// challenge: modify this code so that the calls to updateMessage() on lines
	// 28, 30, and 33 run as goroutines, and implement wait groups so that
	// the program runs properly, and prints out three different messages.
	// Then, write a test for all three functions in this program: updateMessage(),
	// printMessage(), and main().

	msg = "Hello, world!"

	updateMessage("Hello, universe!")
	printMessage()

	updateMessage("Hello, cosmos!")
	printMessage()

	updateMessage("Hello, world!")

	printMessage()

	var wg sync.WaitGroup

	msgs := []string{
		"Hello, universe!",
		"Hello, cosmos!",
		"Hello, world!",
	}

	wg.Add(len(msgs))

	for _, msg := range msgs {
		go func(msg string, wg *sync.WaitGroup) {
			defer wg.Done()
			updateMessage(msg)
			printMessage()
		}(msg, &wg)
	}

	wg.Wait()

	time.Sleep(2 * time.Second)
	fmt.Println("done")
}
