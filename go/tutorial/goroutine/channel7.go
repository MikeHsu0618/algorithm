package main

import (
	"fmt"
	"time"
)

/**

 */

func main() {
	// 使用 select 可以解決從管道取數據的阻塞問題
	// 1. 定義一個管道 10 個數據 int
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	// 2. 定一一個管道 5 個數據 string
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	// 傳統方法在遍歷管道時, 如果不關閉管道, 會阻塞而導致 deadlock

	// 問題, 在實際開發中, 可能我們不好確定什麼時候應該關閉管道
	// 可以使用 select 方式可以解決
	for {
		select {
		case v := <-intChan: // 注意：這裡如果 intChan 一直沒有關閉, 不會一直阻塞而 deadlock
			fmt.Printf("從 intChan 讀取到值 ＝ %d \n", v)
		case v := <-stringChan:
			fmt.Printf("從 stringChan 讀取到值 ＝ %s \n", v)
		default:
			time.Sleep(1 * time.Second)
			println("都讀不到了拉 不玩了, 可以自己加入邏輯")
			return
		}
	}
}
