package main

import (
	"fmt"
	"time"
)

/**
實際應用
1. 開啟一個 writeData 協程, 向 channel intChan 中寫入 50 個整數
2. 開始一個 readData 協程, 從 channel 中讀取 writeData 寫入的數據
3. 注意： writeData 和 readData 操作的是同一個 channel
4. 主線程需要等待 writeData 和 readData 協程都完成工作後才能退出

思路分析：
    創造出一個存放數據的管道以及一個判斷協程完成的管道
    開啟一個讀的協程與一個寫的協程, 同步運行
    當資料不斷的寫入, 而讀取資料也同時不斷的 loop 從管道取出
    直到全部讀取完畢後, 讀取的協程會將判斷是否完成的管道放入數值
    而在主協程上等待的 loop 將會取出值, 以代表所有事務已經結束
*/

func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		// 放入數據
		intChan <- i
		println("放入一個數據拉：", i)
		time.Sleep(time.Millisecond * 100)
	}

	close(intChan) // 關閉
}

func readData(intChan chan int, exitChan chan bool) {
	// 如果該 channel 有 close 的話, 讀取時將會被鎖在外面繼續等到鎖釋放掉
	// 所以沒有用 close 關閉的話, 將會返回 deadlock 的錯誤
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		println("讀到一個數據了 ＝ ", v)
	}
	// readData 讀取完數據後, 即任務完成
	exitChan <- true
	close(exitChan)
}

func main() {
	// 創建兩個管道 -> 一個存取數據, 一個當作離開管道的依據
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	// 給個 for 循環不斷的去取 exitChan 是否已經完成
	for {
		// 方法一
		//if len(exitChan) != 0 {
		//    break
		//}
		println("exitChan Waiting")
		// 方法二
		_, ok := <-exitChan
		fmt.Printf("ok = %v \n", ok)
		if ok {
			break
		}
	}
}
