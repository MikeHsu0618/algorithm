package main

import "fmt"

// 3. 取得最長不重複字段的長度
// 主要使用 slice window, hashtable 技巧
// 可參考 https://ithelp.ithome.com.tw/articles/10287926

func main() {
	fmt.Println(lengthOfLongestSubstring("aab"))
}

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	start, maxLen := 0, 0
	for i := 0; i < len(s); i++ {
		if v, ok := m[s[i]]; ok && v >= start {
			start = v + 1
		}
		m[s[i]] = i
		if (i - start + 1) > maxLen {
			maxLen = i - start + 1
		}
		if start+maxLen > len(s) {
			break
		}
	}
	return maxLen
}
