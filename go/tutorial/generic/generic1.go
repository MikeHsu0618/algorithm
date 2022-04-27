package main

import "fmt"

// 範例 -> 可以使用 官方的 any (interface{}) 來傳入各種類型
func print1[T any](arr []T) {
	for _, v := range arr {
		fmt.Print(v)
		fmt.Print(" ")
	}
	fmt.Println("")
}
func main() {
	strs := []string{"Hello", "World", "Generics"}
	decs := []float64{3.14, 1.14, 1.618, 2.718}
	nums := []int{2, 4, 6, 8}
	print1(strs)
	print1(decs)
	print1(nums)
}
