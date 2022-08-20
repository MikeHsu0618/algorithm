package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("bananas"))
}

func longestPalindrome(s string) int {
	// 將字母出現的次數存在 map 中
	sMap := make(map[rune]int)
	for _, v := range s {
		sMap[v]++
	}
	// 關鍵：計算奇數個字母出現的次數，因為奇數個字母一定會有一個落單
	count := 0
	for _, v := range sMap {
		if v%2 == 1 {
			count++
		}
	}
	// 迴文可以允許一個落單的字母，所以要加回去
	if count > 0 {
		return len(s) - count + 1
	}
	return len(s)
}
