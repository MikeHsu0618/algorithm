package main

// 管道的關閉, 一旦關閉之後就不能在寫入數據, 但可以讀取
func main() {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	// close 這個內置函數, 可以直接關閉 channel
	close(intChan)

	//// 這時不能夠再寫入數據
	//intChan <- 300
	//fmt.Printf("Ok") // panic: send on closed channel

	// channel 關閉後, 是可以讀取的
	n1 := <-intChan
	println("n1=", n1) // 100

	// channel 支持 for range 的方式進行遍歷, 有兩個重點
	// 1. 在遍歷時, 如果 channel 沒有關閉, 則會跳出 deadlock 錯誤
	// 2. 在遍歷時, 如果 channel 已經關閉, 則會正常遍歷數據, 遍歷完後就會退出遍歷

	// 遍歷 channel
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2 // 放入一百個數據到 channel intChan2
	}

	// 關閉 channel
	close(intChan2)

	// 遍歷時如果沒有關閉 channel, 到最後一個後則會噴錯
	for v := range intChan2 {
		println("v= ", v)
	}
}
