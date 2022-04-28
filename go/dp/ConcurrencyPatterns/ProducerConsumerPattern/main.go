package main

import (
	"fmt"
	"time"
)

// Producer Consumer Pattern
// 多個 Producer(生產者)提供任 Job 任務，多個 Consumer(消費者)消費任務

// 有時候系統的任務不會直接執行，而由多個 Producer 程序存到一個 queue 中，再由其他 Consumer 程序讀取 queue 執行，
// 這樣的話可以使 Producer 與 Consumer 程序間沒有直接關係，他們只依賴 queue，即可解耦。

// 有時候系統的任務不會直接執行，而由多個 Producer 程序存到一個 queue 中，
// 再由其他 Consumer 程序讀取 queue 執行，這樣的話可以使 Producer 與 Consumer 程序間沒有直接關係，
// 他們只依賴 queue，即可解耦。

// 例如在微服務的系統下，會利用 kafka 來做 message queue system，
// 這樣即使微服務 auto scaling(水平擴增)也不會為服務找不到彼此。

type UserInfo struct {
	ID   uint32
	Name string
}

var userInfos = []UserInfo{
	{
		1,
		"糖糖",
	},
	{
		2,
		"鹽鹽",
	},
	{
		3,
		"乖乖",
	},
}

func UberProducer(job chan<- UserInfo, i int) {
	job <- userInfos[i]
}

func UberConsumer(job <-chan UserInfo, id int) {
	// 利用 channel 跟 for range 的阻塞來實現
	for userInfo := range job {
		fmt.Printf("uber consumer %d get %s user\n", id, userInfo.Name)
	}
}

func main() {
	job := make(chan UserInfo)
	UberProducerCount := len(userInfos)
	UberConsumerCount := 5

	for i := 0; i < UberProducerCount; i++ {
		go UberProducer(job, i)
	}

	// 開啟多數的 consumer 來等待取得 job
	for i := 0; i < UberConsumerCount; i++ {
		go UberConsumer(job, i)
	}

	time.Sleep(10 * time.Second) // 等待goroutine執行完畢
}
