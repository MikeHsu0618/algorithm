package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const HandleDataPeriod = 3 * time.Second

var (
	mux        = sync.Mutex{}
	timer      = time.NewTimer(HandleDataPeriod)
	bufferChan = make(chan string, 10)
)

func main() {
	forever := make(chan struct{})
	go producer()
	go consumer()
	<-forever
}

func producer() {
	fmt.Println("start producer")
	for {
		func() {
			mux.Lock()
			defer mux.Unlock()
			bufferChan <- "我是一筆資料"
			fmt.Println(fmt.Sprintf("插入完成，目前總共 %v 個", len(bufferChan)))
		}()
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func consumer() {
	fmt.Println("start consumer")
	for {
		select {
		case <-timer.C:
			fmt.Println("定時：執行動作")
			timer.Stop()
			handleData()
			timer.Reset(HandleDataPeriod)
		default:
			if len(bufferChan) != cap(bufferChan) {
				continue
			}
			fmt.Println("緩存已滿：執行動作")
			if !timer.Stop() {
				<-timer.C
			}
			handleData()
			timer.Reset(HandleDataPeriod)
		}
		time.Sleep(50 * time.Millisecond)
	}
}

func handleData() {
	mux.Lock()
	defer mux.Unlock()
	current := len(bufferChan)
	if current == 0 {
		fmt.Println("尚無資料")
		return
	}
	fmt.Println("handleData 啟動中...")
	time.Sleep(2 * time.Second)
	fmt.Println("取出資料時有", len(bufferChan))
	for i := 0; i < current; i++ {
		<-bufferChan
		fmt.Println("取出資料，目前總共 ", len(bufferChan), "個")
	}
	fmt.Println("結束時有", len(bufferChan))
	fmt.Println("handleData 關閉中...")
	time.Sleep(2 * time.Second)
}
