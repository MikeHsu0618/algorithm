package main

import "fmt"

func main() {
	intChan := make(chan int, 3)
	intChan <- 1
	intChan <- 2
	close(intChan)

	// 關閉 channel 後，可以取出但不能插入

	// intChan <- 3 // 報錯
	a := <-intChan
	fmt.Println(a)
}
