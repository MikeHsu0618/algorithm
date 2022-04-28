package main

import (
	"fmt"
	"sync"
	"time"
)

// Thread-Per-Message Pattern

// 每一個 message 開啟一個 goroutine 來運行
// 如果有多個 task 需要運行，每個 task 的執行時間過長，比如 IO 或者是複雜的運算，
// 那一個一個執行 task 是相當浪費時間的，只要 task 沒有執行順序的需求，可以將所有 task 都同時執行。

// 情境：設計一個推播新聞系統，會將新的新聞直接推播出去，我們希望推播系統效率要高。

func PushNews(wg *sync.WaitGroup, news string, startTime time.Time) {
	defer wg.Done()
	time.Sleep(3 * time.Second) // 模擬推播運行的時間
	fmt.Printf("%s cost %s\n", news, time.Since(startTime))
}

func main() {

	start := time.Now()
	allNews := []string{
		"中秋節來了",
		"記得",
		"不要戶外烤肉～",
	}
	wg := new(sync.WaitGroup)
	wg.Add(len(allNews))
	for _, news := range allNews {
		// 每一條訊息都開一個 gorutine 來執行
		go PushNews(wg, news, start)
	}
	wg.Wait()
	fmt.Printf("cost %s", time.Since(start))
}
