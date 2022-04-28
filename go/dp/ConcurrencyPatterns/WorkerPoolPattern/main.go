package main

import (
	"fmt"
	"time"
)

// Worker Pool Pattern
// 設定好 pool 的 goroutine 數量，預先啟動多個 goroutine，把 job 傳給這些 goroutine 執行

// 與 Thread-Per-Message Pattern 類似，都是將 message 或 job 傳給 goroutine 執行的 pattern，不同的是：
//   Thread-Per-Message Patternt 當接收到 message 的時候啟動 goroutine
//   Worker Pool Pattern 是預先啟動好 goroutine，稱之為pool

// 像 socket 這種啟動需要消耗比較多時間的可以預先開啟！

// Worker Pool Pattern 可以限制 pool 的 goroutine 數量，以避免系統負載過大，
// 但也需要考慮 pool 是否過小，導致 news jobs 常常需等待 pool 的情形。

// 情境：設計一個推播新聞系統，會將新的新聞直接推播出去，我們希望推播系統效率要高，
// 並且每次推播都會跟某 service 建立 socket 拿取資料

type SendInfo struct {
	NewsName   string
	FinishTime time.Time
}

func Worker(id int, jobs <-chan string, results chan<- SendInfo, startTime time.Time) {
	fmt.Println("work starting id : ", id)
	time.Sleep(5 * time.Second) // 模擬與某service建立socket的時間
	for job := range jobs {
		time.Sleep(1 * time.Second) // 模擬推播運行的時間
		fmt.Printf("%s cost %s by worker %d\n", job, time.Since(startTime), id)
		results <- SendInfo{job, time.Now()}
	}
}

func main() {
	start := time.Now()
	allNews := []string{
		"中秋節來了",
		"記得",
		"不要戶外烤肉～",
		"也不要吃太撐",
		"中秋節來了",
		"記得",
		"不要戶外烤肉～",
		"也不要吃太撐",
	}
	jobs := make(chan string, len(allNews))
	results := make(chan SendInfo, len(allNews))

	for w := 1; w <= 2; w++ {
		go Worker(w, jobs, results, start)
	}

	for _, news := range allNews {
		jobs <- news
	}

	// do something

	for r := 1; r <= len(allNews); r++ {
		result := <-results
		fmt.Printf("news %s is sent at %s\n", result.NewsName, result.FinishTime)
	}
}
