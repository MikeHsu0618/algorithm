package main

import "fmt"

// 使用 channel ＋ gorutine 找質數
func main() {
	size := 300000
	intChan := make(chan int, size)
	primeChan := make(chan int, size)
	done := make(chan bool, size)

	go putNumbers(size, intChan)

	go findPrimes1(intChan, primeChan, done)

	for primeNum := range primeChan {
		fmt.Println(primeNum)
	}
}

func findPrimes1(intChan chan int, primeChan chan int, done chan bool) {
	for i := range intChan {
		go checkPrime(i, primeChan, done)
	}
	for i := 0; i < cap(done); i++ {
		<-done
	}
	close(primeChan)
}

func checkPrime(num int, primeChan chan int, done chan bool) {
	defer func() {
		done <- true
	}()
	if num == 1 {
		return
	} else if num == 2 {
		primeChan <- num
	} else {
		for i := 2; i < num; i++ {
			if num%i == 0 {
				return
			}
		}
		primeChan <- num
	}

}

func putNumbers(size int, intChan chan int) {
	for i := 1; i <= size; i++ {
		intChan <- i
	}
	close(intChan)
}
