package main

import "math"

//You are given an array of unique integers salary where salary[i] is the
//salary of the iᵗʰ employee.
//
// Return the average salary of employees excluding the minimum and maximum
//salary. Answers within 10⁻⁵ of the actual answer will be accepted.
//
//
// Example 1:
//
//
//Input: salary = [4000,3000,1000,2000]
//Output: 2500.00000
//Explanation: Minimum salary and maximum salary are 1000 and 4000 respectively.
//
//Average salary excluding minimum and maximum salary is (2000+3000) / 2 = 2500
//
//
// Example 2:
//
//
//Input: salary = [1000,2000,3000]
//Output: 2000.00000
//Explanation: Minimum salary and maximum salary are 1000 and 3000 respectively.
//
//Average salary excluding minimum and maximum salary is (2000) / 1 = 2000
//
//
//
// Constraints:
//
//
// 3 <= salary.length <= 100
// 1000 <= salary[i] <= 10⁶
// All the integers of salary are unique.
//
//
// Related Topics Array Sorting 👍 1091 👎 126

//leetcode submit region begin(Prohibit modification and deletion)
func average(salary []int) float64 {
	sum := 0
	min := math.MaxInt
	max := math.MinInt
	for i := 0; i < len(salary); i++ {
		if salary[i] < min {
			min = salary[i]
		}
		if salary[i] > max {
			max = salary[i]
		}

		sum += salary[i]
	}

	return float64(sum-min-max) / float64(len(salary)-2)
}

//leetcode submit region end(Prohibit modification and deletion)
