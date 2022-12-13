package main

import (
	"fmt"
	"sync"
	"time"
)

// 需求：現在計算一到兩百的各個數的階乘, 並且把各個數的階乘放入到 map 中
// 最後顯示出來 要求使用 goroutine 完成

// 思路
// 1. 編寫一個函數, 來計算各個數的階乘, 並放入到 map 中
// 2. 我們啟動多個協程, 統一的將結果放入到 map 中
// 3. map 應該做出一個全局的
// 問題點
// map 的寫入不能是併發操作, 必須要有保護機制
// 主線程不知道協程什麼時候運行完畢, 必須要等到全部完畢才能結束程序

// 使用互斥鎖只能算是一個低水平的資源競爭解法, 通常是使用 Channel 來解決

var result = make(map[int]int, 10)

// 聲明一個全局的互斥鎖
// sync 是一個 package, Mutex 是互斥
var lock sync.Mutex

// test 函數就是計算 n!, 將這個結果放進 map
func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	// 問題一解法: 只讓一個進程去寫入
	lock.Lock()
	result[n] = res
	lock.Unlock()
}

func main() {
	// 在這裡開啟多個協程完成這個任務(200 個)
	for i := 1; i <= 200; i++ {
		go test(i)
	}

	// 休眠十秒, 確認程序全都完成
	time.Sleep(10 * time.Second)
	lock.Lock()
	for i, v := range result {
		fmt.Printf("map[%d]= %d \n", i, v)
	}
	lock.Unlock()
}
