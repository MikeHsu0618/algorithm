package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type server struct {
	bufferChan chan string
	timer      *time.Timer
	mux        sync.Mutex
}

const HandleDataPeriod time.Duration = 5 * time.Second

func main() {
	forever := make(chan struct{})

	s := server{
		bufferChan: make(chan string, 10),
		timer:      time.NewTimer(HandleDataPeriod),
		mux:        sync.Mutex{},
	}
	s.RunBufferInsert()

	<-forever
}

func (s *server) RunBufferInsert() {
	go s.producer()
	go s.consumer()
}

func (s *server) producer() {
	fmt.Println("start producer")
	for {
		func() {
			s.mux.Lock()
			defer s.mux.Unlock()
			s.bufferChan <- "我是一筆資料"
			fmt.Println(fmt.Sprintf("插入完成，目前總共 %v 個", len(s.bufferChan)))
		}()
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	}
}

func (s *server) consumer() {
	fmt.Println("start consumer")
	for {
		select {
		case <-s.timer.C:
			fmt.Println("定時：執行動作")
			s.timer.Stop()
			s.handleData()
			s.timer.Reset(HandleDataPeriod)
		default:
			if len(s.bufferChan) != cap(s.bufferChan) {
				continue
			}
			fmt.Println("緩存已滿：執行動作")
			if !s.timer.Stop() {
				<-s.timer.C
			}
			s.handleData()
			s.timer.Reset(HandleDataPeriod)
		}
		time.Sleep(50 * time.Millisecond)
	}
}

func (s *server) handleData() {
	s.mux.Lock()
	defer s.mux.Unlock()
	current := len(s.bufferChan)
	if current == 0 {
		fmt.Println("尚無資料")
		return
	}
	fmt.Println("取出資料時有", len(s.bufferChan))
	for i := 0; i < current; i++ {
		<-s.bufferChan
		fmt.Println("取出資料，目前總共 ", len(s.bufferChan), "個")
	}
	fmt.Println("結束時有", len(s.bufferChan))
	fmt.Println("handleData 等待 2 秒")
	time.Sleep(2 * time.Second)
}
