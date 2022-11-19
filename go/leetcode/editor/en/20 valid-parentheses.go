package main

//Given a string s containing just the characters '(', ')', '{', '}', '[' and ']
//', determine if the input string is valid.
//
// An input string is valid if:
//
//
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.
//
//
//
// Example 1:
//
//
//Input: s = "()"
//Output: true
//
//
// Example 2:
//
//
//Input: s = "()[]{}"
//Output: true
//
//
// Example 3:
//
//
//Input: s = "(]"
//Output: false
//
//
//
// Constraints:
//
//
// 1 <= s.length <= 10⁴
// s consists of parentheses only '()[]{}'.
//
//
// Related Topics String Stack 👍 16440 👎 837

//leetcode submit region begin(Prohibit modification and deletion)
// [V] 複習
func isValid(s string) bool {
	// 處理 edge case ，單數個字串必不滿足條件
	if len(s)%2 != 0 {
		return false
	}
	// 建立左右成對查詢表
	m := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	// 使用 stack 概念堆疊出所有的左側符號
	stack := []rune{}
	for _, r := range s {
		// 將左側符號放入 stack 中
		if _, ok := m[r]; ok {
			stack = append(stack, r)
			continue
		}
		// 從遍歷到的右側符號與 stack top 的左側符號比較
		if len(stack) == 0 || m[stack[len(stack)-1]] != r {
			return false
		}
		// 比較通過後就將 stack top 取出，繼續下一輪
		stack = stack[:len(stack)-1]
	}
	// 理相狀態會把 stack 取空
	return len(stack) == 0
}

//leetcode submit region end(Prohibit modification and deletion)
