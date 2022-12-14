package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup

func Test_printMessage(t *testing.T) {
	wg.Add(1)
	go updateMessage("AAA")
	wg.Done()
	time.Sleep(1 * time.Second)
	wg.Wait()
	if msg != "AAA" {
		t.Errorf("Error")
	}
}

func Test_updateMessage(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	msg = "AAA"
	printMessage()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "AAA") {
		t.Errorf("Error")
	}
}
