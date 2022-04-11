package main

import (
	"fmt"
	"time"
)

/**
實際應用3: 要求統計 1-8000 的數字中, 哪些是素數？

分析思路：
    傳統方法, 使用一個循環, 判斷各個數是不是素數
    使用併發/併行的方式, 機統計的素數的任務分配給多個 goroutine 去完成 (時間短)
=>
1. 開啟一個 goroutine 的 for 循環將 1-8000 寫入 intChan channel, 完成後就 close
2. 開啟數個 goroutine 的 for 循環不斷的從 intChan 取值, 判斷為素數的就存入 primeChan
3. 當每個判斷素數的 goroutine 執行完畢後, 就對 exitChan 寫入值代表完成
4. 主線程在外面不斷的查詢 exitChan 是否全部完成
*/

func main() {
	GoNum := 4
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000) // 存入結果
	exitChan := make(chan bool, GoNum)

	start := time.Now().UnixMicro()

	// 開啟一個協程向 intChan 放入 1-8000個數
	go putNum(intChan)
	// 開啟四個協程, 從 intChan 取出數據, 並判斷是否為素數, 如果是, 就放入 primeChan
	for i := 0; i < GoNum; i++ {
		go readNum(intChan, primeChan, exitChan)
	}

	go func() {
		// 主線程退出處理
		for i := 0; i < 4; i++ {
			<-exitChan
		}

		end := time.Now().UnixMicro()
		fmt.Println("使用協程時間", end-start)

		// 當我們從這個管道取出了四個結果, 就可以放心的關閉 primeChan
		close(primeChan)
	}()

	for {
		_, ok := <-primeChan
		if !ok {
			break
		}
	}
	println("主線程退出")
}

func readNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool
	for {
		num, ok := <-intChan
		if !ok {
			break
		}

		// 判斷是否為素數
		flag = true
		for i := 2; i < num; i++ {
			// 說明不是素數
			if num%i == 0 {
				flag = false
				break
			}
		}

		if flag {
			// 將這個數放進 primeChan
			primeChan <- num
		}
	}

	println("有一個 prime協程 因為取不到數據退出嚕")
	// 這時我們還不能馬上關閉 primeChan channel

	exitChan <- true
}

func putNum(intChan chan int) {

	for i := 1; i <= 8000; i++ {

		intChan <- i
	}
	close(intChan)
}
