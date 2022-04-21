package main

import "fmt"

func main() {
	num := 100
	intChan := make(chan int, num)

	for i := 0; i < num; i++ {
		intChan <- i
	}

	//  for range 時，如果沒有 close 掉 channel，那當他取了 N+1 次時將會產生 Deadlock
	close(intChan)
	for i := range intChan {
		fmt.Println(i)
	}
}
