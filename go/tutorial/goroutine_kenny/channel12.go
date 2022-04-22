package main

import (
	"fmt"
	"time"
)

func main() {
	// time.NewTimer 設定一個兩秒鐘的定時器
	timer := time.NewTimer(2 * time.Second)
	<-timer.C // 兩秒鐘後可以從 timer 輸出值
	fmt.Println("timeout1")
	// <- timer.C // 第二次 timer.C 會造成 deadlock

	// timer.Reset 定時器重複使用
	timer.Reset(2 * time.Second)
	<-timer.C
	fmt.Println("timeout2")
}
