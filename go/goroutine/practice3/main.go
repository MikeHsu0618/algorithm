package main

import (
	"fmt"
	"time"
)

// 實現異步接口
// 1. 啟動服務之前啟動多個 goroutine
// 2. 啟動 server 服務
// 3. 等待傳參數到通道，如果通道滿了，會主動丟棄
// 4. 如果通道有值，會觸發 goroutine 去做處理

var buffChan = make(chan string, 100)

func producer() {
	for {
		select {
		case buffChan <- "插入一條消息":
			fmt.Println("插入成功")
			time.Sleep(10 * time.Millisecond)
		default:
			fmt.Println("丟棄訊息")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func consumer() {
	for {
		fmt.Println("處理消息：", <-buffChan, "現在有", len(buffChan))
		time.Sleep(300 * time.Millisecond)
	}
}

func workerpool(n int) {
	for i := 0; i < n; i++ {
		go consumer()
	}
}

func main() {
	forever := make(chan struct{})
	go producer()
	go workerpool(3)
	<-forever
}
