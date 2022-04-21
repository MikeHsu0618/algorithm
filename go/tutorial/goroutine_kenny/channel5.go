package main

import (
	"fmt"
	"sync"
)

func main() {
	num := 100
	intChan := make(chan int, num)

	// 單線程 -> 等到值全都進去後才會往下進行
	// for i := 0; i < num; i++ {
	// 	intChan <- i
	// }
	//
	// for i := 0; i < num; i++ {
	// 	fmt.Println(<-intChan)
	// }

	// 多協程的基本應用
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 200; i++ {
			intChan <- i
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < num; i++ {
			fmt.Println(<-intChan)
		}
	}()
	wg.Wait()
	fmt.Println("剩下：", len(intChan))
}
