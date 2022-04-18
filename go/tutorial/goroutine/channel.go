package main

import "fmt"

// channel 底層是使用 mutex 互斥鎖來實現的, 所以在底層會使用隊列來等待數據的存入或取出

func main() {
	// 演示一下管道的使用
	// 1. 初始化：創建一個可以存放三個 int 類型的管道
	var intChan chan int
	intChan = make(chan int, 3)

	// 2. 看看 intChan 是什麼
	fmt.Printf("intChan 的值 = %v, intChan 本身的地址= %v \n", intChan, &intChan)

	// 3. 向管道寫入數據
	intChan <- 10
	num := 200
	intChan <- num
	intChan <- 100

	// 4. 看看管道的長度和 cap（容量）
	// 注意點：當我們給管道寫入數據時, 不能超過其容量
	fmt.Printf("channel len=%v cap=%v \n", len(intChan), cap(intChan)) // 2, 3

	// 5. 從管道裡面讀取數據, 讀了長度會變短, 但容量不變
	// 先進先出
	num2 := <-intChan
	println(num2)                                                      // 10
	fmt.Printf("channel len=%v cap=%v \n", len(intChan), cap(intChan)) // 1, 3

	// 6. 再沒有使用協程的情況下 , 如果我們的管道數據已經全部取出, 再取就會報告 deadlock
	num3 := <-intChan
	num4 := <-intChan
	println(num3, num4) // 200, 100

	// num5 := <-intChan   // 報錯, 需要保護機制

	intChan <- 55 // 當 channel 取出後可以再繼續放入
	<-intChan     // 可以單獨取出, 不賦予變數
}
