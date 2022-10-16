package main

import (
	"fmt"
	"sync"
	"time"
)

// channel 練習
// 1. 啟動一個 goroutine，生成一百個隨機數送到 chan1
// 2. 啟動一個 goroutine，從 chan1 取出值，然後計算其平方，放到 chan2 中
// 3. 在 main 中，從 chan2 取出值打印

var dataChan = make(chan int, 100)
var resultChan = make(chan int, 100)
var wg = sync.WaitGroup{}
var once sync.Once // 確保某個動作只執行一次

func producer() {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		dataChan <- i
		fmt.Println("插入數值", i)
		time.Sleep(100 * time.Millisecond)
	}
	close(dataChan)
}

func consumer() {
	defer wg.Done()
	for data := range dataChan {
		resultChan <- data * data
		fmt.Println("處理數值", data*data)
	}
	// 如果要確保只有一個線程可以關閉 channel 需要用到 sync.Once
	once.Do(func() {
		close(resultChan)
		fmt.Println("我被關閉拉！")
		time.Sleep(2 * time.Second)
	})
}

func main() {
	wg.Add(3)
	go producer()
	go consumer()
	go consumer()
	wg.Wait()
	for result := range resultChan {
		fmt.Println(result)
	}
	fmt.Println("完成了哥")
}
