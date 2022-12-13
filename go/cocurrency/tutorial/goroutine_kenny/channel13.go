package main

import (
	"fmt"
	"time"
)

//
func main() {
	quitChan := make(chan bool, 1)
	go func() {
		time.Sleep(2 * time.Second)
		quitChan <- true
	}()

	// select 將會等待 channel 直到傳值進來，如果兩秒以上將會接收到 timer 回傳並 timeout
	select {
	case <-quitChan:
		fmt.Println("over")
	case <-time.After(2 * time.Second):
		fmt.Println("timeout")
	}
}
