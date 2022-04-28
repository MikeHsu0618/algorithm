package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Fan-Out Fan-In Pattern
// 將 input 由一個 producer 分發多個 goroutine 運行，
// 再將多個 task goroutine 運行的結果由一個 consumer 收集資料合併為 output

// Fan Out: input 傳入 producer 後開啟多個 goroutine 運行，
// 直到 producer 不再接收 input，就像是分發任務一般，所以稱為 Fan Out
// Fan In: 由多個 input 傳入 consumer 後合併為 output 傳出，
// 直到 consumer 不再接收 input，就像是收集資料一般，所以稱為 Fan In

func Producer(serverNames ...string) <-chan string {
	producerCh := make(chan string, len(serverNames))
	go func() {
		defer close(producerCh)
		for _, serverName := range serverNames {
			producerCh <- serverName
		}
	}()
	return producerCh
}

func Task(producerCh <-chan string) <-chan string {
	taskCh := make(chan string)
	go func() {
		defer close(taskCh)
		for serverName := range producerCh {
			taskCh <- GetServerData(serverName)
		}
	}()
	return taskCh
}

func Consumer(taskChs ...<-chan string) <-chan string {
	consumerCh := make(chan string)

	var wg sync.WaitGroup
	wg.Add(len(taskChs))
	go func() {
		wg.Wait()
		close(consumerCh)
	}()

	for _, task := range taskChs {
		go func(task <-chan string) {
			defer wg.Done()
			for news := range task {
				consumerCh <- news
			}
		}(task)
	}

	return consumerCh
}

func GetServerData(serverName string) string {
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second) // 模擬取得server data消耗的時間
	return fmt.Sprintf("%s server data", serverName)
}

func ShowNews(news ...interface{}) {
	fmt.Println(news...)
}

func main() {
	start := time.Now()
	producerCh := Producer("A", "B", "C")

	task1 := Task(producerCh)
	task2 := Task(producerCh)
	task3 := Task(producerCh)

	consumerCh := Consumer(task1, task2, task3)

	for news := range consumerCh {
		ShowNews(news)
	}
	fmt.Printf("cost %s", time.Since(start))
}
