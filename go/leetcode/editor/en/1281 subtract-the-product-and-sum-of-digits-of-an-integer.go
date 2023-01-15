package main

//Given an integer number n, return the difference between the product of its
//digits and the sum of its digits.
//
//
// Example 1:
//
//
//Input: n = 234
//Output: 15
//Explanation:
//Product of digits = 2 * 3 * 4 = 24
//Sum of digits = 2 + 3 + 4 = 9
//Result = 24 - 9 = 15
//
//
// Example 2:
//
//
//Input: n = 4421
//Output: 21
//Explanation:
//Product of digits = 4 * 4 * 2 * 1 = 32
//Sum of digits = 4 + 4 + 2 + 1 = 11
//Result = 32 - 11 = 21
//
//
//
// Constraints:
//
//
// 1 <= n <= 10^5
//
//
// Related Topics Math 👍 1851 👎 204

//leetcode submit region begin(Prohibit modification and deletion)
func subtractProductAndSum(n int) int {
	sum, product := 0, 1

	for n > 0 {
		mod := n % 10
		sum += mod
		product *= mod

		n /= 10
	}

	return product - sum
}

func subtractProductAndSum1(n int) int {
	arr := make([]int, 0)
	prod := 1
	sum := 0

	for n != 0 {
		arr = append(arr, n%10)
		n /= 10
	}

	for i := 0; i < len(arr); i++ {
		prod *= arr[i]
		sum += arr[i]
	}

	return prod - sum
}

//leetcode submit region end(Prohibit modification and deletion)
