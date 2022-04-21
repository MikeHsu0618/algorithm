package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 使用 value, ok := <- intChan 取值來確認還有沒有下一個數
func main() {
	intChan := make(chan int, 11)
	go generateNumbers1(intChan)

	for {
		if num, ok := <-intChan; ok {
			fmt.Printf("num : %v, isSucc : %v \n", num, ok)
			time.Sleep(time.Second)
		} else {
			fmt.Printf("num : %v, isSucc : %v \n", num, ok)
			break
		}
	}
	fmt.Println("結束了 各位")
}

func generateNumbers1(intChan chan int) {
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
