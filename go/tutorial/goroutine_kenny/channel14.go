package main

import (
	"fmt"
	"time"
)

var boolChan = make(chan bool, 1)

func main() {
	// example 1
	// time.AfterFunc 前面放時間 後面放 func
	// 指定幾秒後呼叫指定的 func
	time.AfterFunc(2*time.Second, hello)

	// time.AfterFunc 等待的時間要自己設定
	time.Sleep(3 * time.Second)

	fmt.Println("done")

	// example 2
	time.AfterFunc(2*time.Second, hello2)
	// 可以利用 channel 的管道 block 住
	<-boolChan
	fmt.Println("over timer")
}

func hello2() {
	fmt.Println("Say Hello 2")
	boolChan <- true
}

func hello() {
	fmt.Println("Say Hello")
}
