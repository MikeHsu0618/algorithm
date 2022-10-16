package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 題目：實作一個每十秒或每十筆存入一次資料的程式
// 1. 做出一個阻塞 channel 來防止主線程執行完畢
// 2. 聲明一個 bufferChannel 來堆放資料
// 3. 產生一個 producer 的 goroutine，隨機時間插入資料
// 4. 聲明一個 timer 才定時或重置時間
// 5. 產生一個 consumer 來使用 timer 對資料消化
// 6. 優化：加上互斥鎖確保消化以及插入的動作不會同時執行，並補上 time.Sleep 幫助查看

var bufferChan = make(chan string, 10)
var timer = time.NewTimer(10 * time.Second)
var mux = sync.Mutex{}

func main() {
	forever := make(chan struct{})
	go producer()
	go consumer()
	<-forever
}

func producer() {
	fmt.Println("start producer")
	for {
		mux.Lock()
		bufferChan <- fmt.Sprintf("%v", time.Now())
		fmt.Printf("插入資料，現在總共 %v 筆 \n", len(bufferChan))
		mux.Unlock()
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	}
}

func consumer() {
	fmt.Println("start consumer")
	for {
		select {
		case <-timer.C:
			fmt.Println("到達預定時閒：執行動作")
			handleBuffer()
			timer.Reset(10 * time.Second)
		default:
			if len(bufferChan) != cap(bufferChan) {
				continue
			}
			fmt.Println("到達預定數目：執行動作")
			if !timer.Stop() {
				<-timer.C
			}
			handleBuffer()
			timer.Reset(10 * time.Second)
		}
		time.Sleep(50 * time.Millisecond)
	}
}

func handleBuffer() {
	mux.Lock()
	defer mux.Unlock()
	fmt.Println("開始處理資料...")
	time.Sleep(2 * time.Second)
	current := len(bufferChan)
	for i := 0; i < current; i++ {
		fmt.Printf("處理資料：%v \n", <-bufferChan)
	}
	fmt.Println("結束處理資料...")
	time.Sleep(2 * time.Second)
}
