package main

import (
	"fmt"
	"sync"
	"time"
)

// 尋找一到一個數字間的質數
// 進階使用 wait group 來等待協程結束
func main() {
	num := 800000
	wg := new(sync.WaitGroup)
	wg.Add(num)
	start := time.Now()
	for i := 1; i <= num; i++ {
		go findPrimesWaitGroup(i, wg)
	}
	end := time.Now()
	fmt.Println(end.Unix()-start.Unix(), "seconds")
}

func findPrimesWaitGroup(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	if num == 1 {
		return
	} else if num == 2 {
		fmt.Println(num)
	} else {
		for i := 2; i < num; i++ {
			if num%i == 0 {
				return
			}
		}
		fmt.Println(num)
	}
}
