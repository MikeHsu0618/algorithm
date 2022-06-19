package main

import "fmt"

func main() {
	// 	map 是一個引用類型, 任何對 map 的修改都會改到原本的 map
	map1 := make(map[int]int)
	map1[1] = 90
	fmt.Println(map1)
	modifyMap(map1)
	fmt.Println(map1) // 結果將會影響原有的 map1 說明 map 是引用類型
}

func modifyMap(map1 map[int]int) {
	map1[1] = 900
}
