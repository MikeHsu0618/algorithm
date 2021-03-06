package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	errChan1 := make(chan error, 1)
	errChan2 := make(chan error, 1)
	go cronjob(time.Second, process1, errChan1)
	go cronjob(time.Second, process2, errChan2)

	for {
		select {
		case err := <-errChan1:
			fmt.Println(err.Error())
			return
		case err := <-errChan2:
			fmt.Println(err.Error())
			return
		}
	}
}

func process1() (err error) {
	fmt.Println("process1 start")
	time.Sleep(time.Second)
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(10) + 1
	if number == 3 {
		err = errors.New("process1 get error")
	}
	fmt.Println("process1 finished")
	fmt.Println("=================")
	return err
}

func process2() (err error) {
	fmt.Println("process2 start")
	time.Sleep(time.Second)
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(10) + 1
	if number == 3 {
		err = errors.New("process1 get error")
	}
	fmt.Println("process2 finished")
	fmt.Println("=================")
	return err
}

func cronjob(d time.Duration, f func() error, errChan chan error) {
	ticker := time.NewTicker(d)
	for {
		select {
		case <-ticker.C:
			if err := f(); err != nil {
				errChan <- err
			}
		}
	}
}
