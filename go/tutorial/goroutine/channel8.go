package main

import "time"

/**
如果我們起的眾多個協程裡有個協程出現 panic, 這時如果我們沒有捕捉這個 panic 那整個程序就會直接崩潰
所以我們可以直接在 goroutine 中使用 recover 來捕獲 panic 進行處理, 這樣即使這個協程發生問題
主線程仍然不受影響可以繼續執行
*/

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		println("hello world")
	}
}

func error1() {
	// 這裡我們可以使用 defer + recover 處理
	defer func() {
		if err := recover(); err != nil {
			println("error!() 發生錯誤", err)
		}
	}()

	var myMap map[int]string
	myMap[0] = "golang" // 沒有先 make 會噴 panic
}

func main() {
	go error1()
	go sayHello()

	for i := 0; i < 10; i++ {
		println("main ok 的", i)
		time.Sleep(time.Second)
	}
}
