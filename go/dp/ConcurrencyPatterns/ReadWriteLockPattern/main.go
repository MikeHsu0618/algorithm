package main

import (
	"context"
	"fmt"
	"log"
	"sync"
)

// 多讀單寫。將 lock 分為 read 與 write 兩種，讓 lock 效能更佳，read 行為不會改變資料，
// 所以 read lock 時可以再同時 read，達到「多讀」，而 write 行為會改變資料，
// 所以 write lock 不能再同時 read 與 write 時，達到「單寫」

// 情境：許多人同時執行點讚以及讀讚的動作，因為

type Like struct {
	sync.RWMutex
	Count uint16
}

// 在執行的動作中加入讀寫鎖 sync.RWMutex 來防止多協程時的資源競爭
func (l *Like) Add(writerID string) {
	l.Lock()
	defer l.Unlock()
	l.Count++
}

func (l *Like) Show(readerID string) {
	l.RLock()
	defer l.RUnlock()
	fmt.Printf("%s read count: %d\n", readerID, l.Count)
}

func AddLikes(cancel context.CancelFunc, writerID string, like *Like) {
	defer cancel()

	for i := 0; i < 10000; i++ {
		like.Add(writerID)
	}
	log.Println("Canceled")
}

func ReadLikes(cancel context.CancelFunc, readerID string, like *Like) {
	defer cancel()
	for i := 0; i < 200; i++ {
		like.Show(readerID)
	}
}

func main() {
	ctx1, cancel1 := context.WithCancel(context.Background())
	ctx2, cancel2 := context.WithCancel(context.Background())
	ctx3, cancel3 := context.WithCancel(context.Background())
	ctx4, cancel4 := context.WithCancel(context.Background())

	like := new(Like)
	go AddLikes(cancel1, "A", like)
	go ReadLikes(cancel2, "B", like)
	go ReadLikes(cancel3, "C", like)
	go AddLikes(cancel4, "D", like)

	// 等待goroutine執行完畢
	<-ctx1.Done()
	<-ctx2.Done()
	<-ctx3.Done()
	<-ctx4.Done()

	fmt.Println(like.Count)
}
