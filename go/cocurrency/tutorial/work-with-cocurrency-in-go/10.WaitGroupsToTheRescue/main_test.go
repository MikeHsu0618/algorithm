package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func TestPrintSomething(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)
	go PrintSomething("aa", &wg)

	wg.Wait()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "aa") {
		t.Errorf("Error, it's not there")
	}
}
