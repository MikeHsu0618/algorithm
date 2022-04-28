package main

import (
	"fmt"
	"time"
)

// Future Pattern

// 呼叫者將 task 交給 goroutine 執行，執行完畢後 goroutine 將 task 運行得到的結果傳回呼叫者
// 需要一個中間物件去得 goroutine 未來運行的結果，golang 通常採用 channel 實作

// 情境：延續推播新聞的情境，將新的新聞直接推播出去，還需要紀錄推播完成的時間

func PushNews(news string, startTime time.Time) <-chan time.Time {
	newsCh := make(chan time.Time)
	go func() {
		time.Sleep(time.Duration(3 * time.Second)) // 模擬推播運行的時間
		fmt.Printf("%s cost %s\n", news, time.Since(startTime))
		newsCh <- time.Now()
	}()
	return newsCh
}

func main() {
	start := time.Now()
	allNews := []string{
		"中秋節來了",
		"記得",
		"不要戶外烤肉～",
	}

	newsChs := []<-chan time.Time{}

	for _, news := range allNews {
		newsChs = append(newsChs, PushNews(news, start))
	}
	// do something

	// 這裡將會等待每次的 <-newsCh 輸出，達到 Future Pattern 的目的
	for index, newsCh := range newsChs {
		fmt.Printf("news %d is sent at %s\n", index, <-newsCh)
	}
}
