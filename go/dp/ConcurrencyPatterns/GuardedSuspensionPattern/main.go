package main

import (
	"fmt"
	"math/rand"
	"time"
)

//  Guarded Suspension Pattern 漫談模式

// 如果 thread 執行時條件不符，就停下等待，等到條件符合時再開始開始執行

// 情境：當設計線上即時顯示留言板時，會接受大量的留言，並將其顯示出來，
// 我們需確保留言不出現 race condition，也需確保沒留言時不執行顯示。
// 思路：利用 channel 的阻塞來實現有任務才執行

// channel 有幾個特性：
// 在讀寫時不會產生 race condition
// 在讀取時如果沒有元素，會等待到有值再執行
// 如果使用 unbuffered channel，e.g. make(chan string)，寫入後如果沒有讀取，會在寫入處等待
// 如果使用 buffered channel，e.g. make(chan string, 100)，
// 寫入後如果沒有讀取，可以繼續寫入，直到到達 buffer 數才會等待讀取

type MessageBoard struct {
	messages chan string
}

func (m *MessageBoard) SendMessage(message string) {
	m.messages <- message
}

func (m *MessageBoard) Show() {
	for {
		fmt.Println(<-m.messages)
	}
}

func main() {
	messageBoard := new(MessageBoard)
	messageBoard.messages = make(chan string, 100) // 最多可容納 100 條訊息

	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond) // 模擬各個client發送的隨機處理時間
			messageBoard.SendMessage("cool")
		}
	}()
	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond) // 模擬各個client發送的隨機處理時間
			messageBoard.SendMessage("wow")
		}
	}()
	go func() {
		messageBoard.Show()
	}()
	time.Sleep(1 * time.Second) // 等待goroutine執行完畢
}
