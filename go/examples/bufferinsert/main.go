package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mux = sync.Mutex{}

const HandleDataPeriod time.Duration = 3 * time.Second

func main() {
	forever := make(chan struct{})
	t := time.NewTimer(HandleDataPeriod)
	bufferChan := make(chan string, 10)

	go producer(bufferChan)
	go consumer(bufferChan, t)
	<-forever
}

func producer(bufferChan chan<- string) {
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

func consumer(bufferChan <-chan string, t *time.Timer) {
	fmt.Println("start consumer")
	for {
		select {
		case <-t.C:
			fmt.Println("定時：執行動作")
			t.Stop()
			handleData(bufferChan)
			t.Reset(HandleDataPeriod)
		default:
			if len(bufferChan) != cap(bufferChan) {
				continue
			}
			fmt.Println("緩存已滿：執行動作")
			if !t.Stop() {
				<-t.C
			}
			handleData(bufferChan)
			t.Reset(HandleDataPeriod)
		}
		time.Sleep(50 * time.Millisecond)
	}
}

func handleData(bufferChan <-chan string) {
	mux.Lock()
	defer mux.Unlock()
	current := len(bufferChan)
	if current == 0 {
		fmt.Println("尚無資料")
		return
	}
	fmt.Println("取出資料時有", len(bufferChan))
	for i := 0; i < current; i++ {
		<-bufferChan
		fmt.Println("取出資料，目前總共 ", len(bufferChan), "個")
	}
	fmt.Println("結束時有", len(bufferChan))
	fmt.Println("handleData 等待 2 秒")
	time.Sleep(2 * time.Second)
}
