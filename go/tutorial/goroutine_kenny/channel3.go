package main

// 只讀, 只寫, 讀寫
func main() {
	intChan := make(chan<- int, 3) // 只寫
	strChan := make(<-chan int, 3) // 只讀
	intChan <- 1
	// <- intChan 編譯器將會提示錯誤 不能讀
	<-strChan
	// strChan <- "hello"
}
