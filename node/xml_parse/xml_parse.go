package main

import (
	"log"
	"os"
)

func main() {
	// open xml
	file, err := os.Open("examples.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}
