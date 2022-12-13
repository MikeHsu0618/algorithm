package main

import (
	"fmt"
	"math/rand"
)

// 製造 block 狀態, 使主進程等待協程
func main() {
	intChan := make(chan int, 1)
	go generateNumbers(intChan)

	for num := range intChan {
		fmt.Println("get num:", num)
	}
	fmt.Println("結束了 各位")
}

func generateNumbers(intChan chan int) {
	count := 0
	for {
		if count != 6 {
			num := rand.Intn(10) + 1
			intChan <- num
			fmt.Println("put num:", num)
			count++
		} else {
			close(intChan)
			break
		}
	}
}
