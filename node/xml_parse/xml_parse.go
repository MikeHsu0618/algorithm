package main

import (
	"log"
	"os"
)

func main() {
	// open xml
	file, err := os.Open("example.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}
