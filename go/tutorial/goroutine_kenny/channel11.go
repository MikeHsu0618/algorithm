package main

import (
	"fmt"
	"time"
)

func main() {
	quitChan := make(chan bool, 1)

	go func() {
		time.Sleep(time.Second)
		quitChan <- true
	}()

	for {
		flag := false
		select {
		case <-quitChan:
			fmt.Println("done")
			flag = true

			// done 後 return 的話是直接跳出此 func
			// return
		default:
			fmt.Println("wait")
			time.Sleep(time.Second)
		}
		if flag {
			break
		}
	}

	fmt.Println("結束了 哥")
}
