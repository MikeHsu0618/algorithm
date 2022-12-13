package main

import (
	"fmt"
	"time"
)

func main() {
	// 定時不斷的塞值進去 timer channel
	ticker := time.NewTicker(time.Second)

	go func() {
		// ticker 底層也是一個 chanel 所以可以用 ticker.C 去 for range 取值出來
		for t := range ticker.C {
			fmt.Println(" tick at :", t)
		}
	}()

	time.Sleep(3 * time.Second)
	// ticker.Stop() // ticker.Stop 會停止 timer
	fmt.Println("over")
}
