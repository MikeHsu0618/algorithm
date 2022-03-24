package main

/**
管道可以聲明為只讀或只寫
*/

func main() {
	// 1. 在默認情況下管道是雙向的
	//var chan1 chan int // 可讀可寫

	// 2. 聲明為只寫
	//var chan2 chan<- int
	//chan2 = make(chan int, 3)
	//num := <-chan2 // 會報錯, 只能寫

	// 3. 聲明為只讀
	//var chan3 <-chan int
	//num2 := <-chan3
	//chan3<-30 // 報錯, 只能讀

}
