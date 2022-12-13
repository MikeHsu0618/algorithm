package main

import (
	"fmt"
	"time"
)

// 生產者-消費者 練習

// 使用 goroutine 實現生產者消費者模型
// 1. 開啟一個 worker pool，啟動指定數量的 goroutine，讀取 channel 裡面的數據，處理完成後丟到另一個 channel 中
// 2. 開啟多個任務給 channel 裡面輸入數據
// 3. 打印處理完成的 channel 裡面的數據

var jobChan = make(chan int, 100)
var resultChan = make(chan int, 100)
var exitChan = make(chan struct{}, 5)

func producer(id int) {
	for i := 0; i < 5; i++ {
		jobChan <- int(time.Now().UnixNano())
		time.Sleep(1 * time.Second)
	}
	fmt.Println("producer:", id, " 執行完畢")
	exitChan <- struct{}{}
}

func consumer() {
	for job := range jobChan {
		resultChan <- job * 2
	}
}

func main() {
	for i := 0; i < 3; i++ {
		go producer(i)
	}
	go func() {
		for {
			if len(exitChan) != 3 {
				continue
			}
			fmt.Println("close")
			close(jobChan)
			break
		}
	}()
	consumer()
	//close(resultChan)

	for i := 0; i < len(resultChan); i++ {
		fmt.Println("結果就是", <-resultChan)
	}
}
