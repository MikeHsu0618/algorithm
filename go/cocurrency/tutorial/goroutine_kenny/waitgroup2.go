package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	jobs := 100
	// 常見錯誤1 如果 jobs 數大於 gorutine 數目時，會產生 deadlock 注意！
	// 常見錯誤2 如果 jobs 數小於 gorutine 數目時，會產生 waitGroup Counter 不得為負的錯誤！ 不能多也不能少
	wg.Add(jobs)
	for i := 0; i < jobs; i++ {
		go doTask(wg)
	}
	wg.Wait()

	// waitgroup 可以重複使用，會以 wg.wait() 來當分界線！
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// }()
	// wg.wait()

	fmt.Println("jobs all done")
}

func doTask(wg *sync.WaitGroup) {
	defer func() {
		println("do a job")
		wg.Done()
	}()
}
