package main

//Given a positive integer num, write a function which returns True if num is a
//perfect square else False.
//
// Follow up: Do not use any built-in library function such as sqrt.
//
//
// Example 1:
// Input: num = 16
//Output: true
//
// Example 2:
// Input: num = 14
//Output: false
//
//
// Constraints:
//
//
// 1 <= num <= 2^31 - 1
//
//
// Related Topics Math Binary Search 👍 3105 👎 265

//leetcode submit region begin(Prohibit modification and deletion)
func isPerfectSquare(num int) bool {
	l, r := 1, num

	for l <= r {
		mid := (l + r) / 2

		if mid*mid < num {
			l = mid + 1
			continue
		}

		if mid*mid > num {
			r = mid - 1
			continue
		}

		return true
	}

	return false
}

// Likely something more efficient is possible, but this was simple to code. The
// idea is actually borrowed from the discussion list:
// 1 = 1
// 4 = 1 + 3
// 9 = 1 + 3 + 5
// 16 = 1 + 3 + 5 + 7
// .... and so on

func isPerfectSquare1(num int) bool {
	var s int
	i := 1
	for s < num {
		s += i
		i += 2
	}

	if s == num {
		return true
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
