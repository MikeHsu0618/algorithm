package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		for {
			select {
			case <-ctx.Done():
				println("hold up hold up")
				time.Sleep(2 * time.Second)
				fmt.Println("Break the loop")
				wg.Done()
				return
			case <-time.After(1 * time.Second):
				fmt.Println("Hello in a loop")
			}
		}
	}()

	wg.Add(1)
	go func() {
		for {
			select {
			case <-ctx.Done():
				println("hold up hold up2222")
				time.Sleep(2 * time.Second)
				fmt.Println("Break the loop222")
				wg.Done()
				return
			case <-time.After(1 * time.Second):
				fmt.Println("Ciao in a loop")
			}
		}
	}()

	wg.Wait()
	fmt.Println("Main done")
}
