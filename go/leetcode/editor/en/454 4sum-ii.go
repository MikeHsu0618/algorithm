package main

//Given four integer arrays nums1, nums2, nums3, and nums4 all of length n,
//return the number of tuples (i, j, k, l) such that:
//
//
// 0 <= i, j, k, l < n
// nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0
//
//
//
// Example 1:
//
//
//Input: nums1 = [1,2], nums2 = [-2,-1], nums3 = [-1,2], nums4 = [0,2]
//Output: 2
//Explanation:
//The two tuples are:
//1. (0, 0, 0, 1) -> nums1[0] + nums2[0] + nums3[0] + nums4[1] = 1 + (-2) + (-1)
// + 2 = 0
//2. (1, 1, 0, 0) -> nums1[1] + nums2[1] + nums3[0] + nums4[0] = 2 + (-1) + (-1)
// + 0 = 0
//
//
// Example 2:
//
//
//Input: nums1 = [0], nums2 = [0], nums3 = [0], nums4 = [0]
//Output: 1
//
//
//
// Constraints:
//
//
// n == nums1.length
// n == nums2.length
// n == nums3.length
// n == nums4.length
// 1 <= n <= 200
// -2²⁸ <= nums1[i], nums2[i], nums3[i], nums4[i] <= 2²⁸
//
//
// Related Topics Array Hash Table 👍 4219 👎 122

//leetcode submit region begin(Prohibit modification and deletion)
// 解法：拆成兩對，實現 a+b=-(c+d)，並計算相加為零的次數
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m := make(map[int]int, 0)
	count := 0
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			m[nums1[i]+nums2[j]]++
		}
	}
	for i := 0; i < len(nums3); i++ {
		for j := 0; j < len(nums4); j++ {
			// 這裡不需要判斷 m[-(nums3[i]+nums4[j])] 直接相加即可，因為不存在的值為零值不影響
			count += m[-(nums3[i] + nums4[j])]
		}
	}
	return count
}

//leetcode submit region end(Prohibit modification and deletion)
