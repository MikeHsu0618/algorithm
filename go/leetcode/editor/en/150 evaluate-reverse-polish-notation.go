package main

import "strconv"

//Evaluate the value of an arithmetic expression in Reverse Polish Notation.
//
// Valid operators are +, -, *, and /. Each operand may be an integer or
//another expression.
//
// Note that division between two integers should truncate toward zero.
//
// It is guaranteed that the given RPN expression is always valid. That means
//the expression would always evaluate to a result, and there will not be any
//division by zero operation.
//
//
// Example 1:
//
//
//Input: tokens = ["2","1","+","3","*"]
//Output: 9
//Explanation: ((2 + 1) * 3) = 9
//
//
// Example 2:
//
//
//Input: tokens = ["4","13","5","/","+"]
//Output: 6
//Explanation: (4 + (13 / 5)) = 6
//
//
// Example 3:
//
//
//Input: tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
//Output: 22
//Explanation: ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
//= ((10 * (6 / (12 * -11))) + 17) + 5
//= ((10 * (6 / -132)) + 17) + 5
//= ((10 * 0) + 17) + 5
//= (0 + 17) + 5
//= 17 + 5
//= 22
//
//
//
// Constraints:
//
//
// 1 <= tokens.length <= 10⁴
// tokens[i] is either an operator: "+", "-", "*", or "/", or an integer in the
//range [-200, 200].
//
//
// Related Topics Array Math Stack 👍 4256 👎 747

//leetcode submit region begin(Prohibit modification and deletion)
// 解法：將數值持續丟進 stack，直到遇到運算子才取出兩個數值進行運算
func evalRPN(tokens []string) int {
	if len(tokens) == 1 {
		num, _ := strconv.Atoi(string(tokens[0]))
		return num
	}
	amount := 0
	stack := make([]int, 0)
	operators := map[string]struct{}{
		"+": struct{}{},
		"-": struct{}{},
		"*": struct{}{},
		"/": struct{}{},
	}
	for i := 0; i < len(tokens); i++ {
		if _, ok := operators[tokens[i]]; !ok {
			num, _ := strconv.Atoi(string(tokens[i]))
			stack = append(stack, num)
			continue
		}
		if tokens[i] == "+" {
			amount = (stack[len(stack)-2]) + stack[len(stack)-1]
		}
		if tokens[i] == "-" {
			amount = (stack[len(stack)-2]) - stack[len(stack)-1]
		}
		if tokens[i] == "*" {
			amount = (stack[len(stack)-2]) * stack[len(stack)-1]
		}
		if tokens[i] == "/" {
			amount = (stack[len(stack)-2]) / stack[len(stack)-1]
		}
		stack = append(stack[:len(stack)-2], amount)
	}
	return amount
}

//leetcode submit region end(Prohibit modification and deletion)
