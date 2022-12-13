package main

import (
	"fmt"
	"sort"
)

func main() {
	// map 預設是沒有排序的，且每次遍歷出來順序都會不一樣
	map1 := make(map[int]int, 10)
	map1[99] = 1
	map1[100] = 10
	map1[1] = 1
	map1[2] = 2
	map1[10] = 10
	map1[3] = 3
	map1[4] = 4
	fmt.Println(map1)

	// 如果要按照 map 的 key 順序進行輸出
	// 1. 先將 map 的 key 放入切片中
	// 2. 對切片排序
	// 3. 遍歷切片，然後按照 key 來輸出 map 的值
	var keys []int
	for k, _ := range map1 {
		keys = append(keys, k)
	}
	// 排序
	sort.Ints(keys)
	fmt.Println(keys)

	for _, key := range keys {
		fmt.Println(key, map1[key])
	}
}
