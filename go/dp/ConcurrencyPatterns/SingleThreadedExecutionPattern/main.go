package main

import (
	"context"
	"fmt"
	"log"
	"sync"
)

// 透過一個 lock，使同一時間只有一個 goroutine 執行此區段的代碼
// 情境：許多人同時執行一個點讚的動作同時也有一群人執行讀讚，將會有機會讀到還沒有更新的讚數
type Like struct {
	sync.Mutex
	Count uint16
}

// 在執行的動作中加入互斥鎖 sync.Mutex 來防止多協程時的資源競爭
func (l *Like) Add(writerID string) {
	l.Lock()
	defer l.Unlock()
	l.Count++
}

func AddLikes(cancel context.CancelFunc, writerID string, like *Like) {
	for i := 0; i < 10000; i++ {
		like.Add(writerID)
	}
	log.Println("Canceled")
	cancel()
}

func main() {
	ctx1, cancel1 := context.WithCancel(context.Background())
	ctx2, cancel2 := context.WithCancel(context.Background())
	ctx3, cancel3 := context.WithCancel(context.Background())
	ctx4, cancel4 := context.WithCancel(context.Background())

	like := new(Like)
	go AddLikes(cancel1, "A", like)
	go AddLikes(cancel2, "B", like)
	go AddLikes(cancel3, "C", like)
	go AddLikes(cancel4, "D", like)

	// 等待goroutine執行完畢
	<-ctx1.Done()
	<-ctx2.Done()
	<-ctx3.Done()
	<-ctx4.Done()

	fmt.Println(like.Count)
}
