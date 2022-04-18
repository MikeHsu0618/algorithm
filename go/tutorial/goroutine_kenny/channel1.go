package main

import "fmt"

func main() {
	// 此方式宣告是沒有容量的
	var intChan chan int
	// 可以設定容量
	intChan = make(chan int, 1)
	fmt.Println(intChan) // nil

	// 沒給容量：deadlock 因為不知道容量為多少
	// 有給容量：小於容量返回指針，大於容量 deadlock
	intChan <- 10

	// 取得長度
	fmt.Println(len(intChan))
	// 取出值(先進先出)
	fmt.Println(<-intChan)

	// 如果只有進沒有出就會堵塞造成 deadlock
	intChan = make(chan int, 3)
	intChan <- 1
	intChan <- 2
	intChan <- 3
	b := <-intChan
	intChan <- 4
	println("b=", b)

	strChan := make(chan string)
	// 此處會照成 deadlock 因為塞入值後沒有被取出， 必須使用 gorutine 不斷的往裡面取
	// strChan <- "hello world"
	// <-strChan

	// 需要開啟 gorutine 同時取資料 塞資料
	go func() {
		strChan <- "hello world"
	}()

	fmt.Println(<-strChan)
}
