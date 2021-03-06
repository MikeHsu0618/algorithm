package main

import (
	"fmt"
	"time"
)

// 尋找一到一個數字間的質數
func main() {
	num := 300000
	start := time.Now()
	for i := 1; i <= num; i++ {
		// 沒有使用 gorutine 的情況
		// findPrimesNormal(i)

		// 使用 gorutine 的情況
		go findPrimes(i)
	}
	time.Sleep(5 * time.Second)
	end := time.Now()
	fmt.Println(end.Unix()-start.Unix(), "seconds")
}

func findPrimes(num int) {
	if num == 1 {
		return
	} else if num == 2 {
		println(num)
	} else {
		for i := 2; i < num; i++ {
			if num%i == 0 {
				return
			}
		}
		println(num)
	}
}

func findPrimesNormal(num int) {
	if num == 1 {
		return
	} else if num == 2 {
		fmt.Println(num)
	} else {
		for i := 2; i < num; i++ {
			if num%i == 0 {
				return
			}
		}
		fmt.Println(num)
	}
}
