package main

//Given two integer arrays nums1 and nums2, return an array of their
//intersection. Each element in the result must be unique and you may return the result in
//any order.
//
//
// Example 1:
//
//
//Input: nums1 = [1,2,2,1], nums2 = [2,2]
//Output: [2]
//
//
// Example 2:
//
//
//Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//Output: [9,4]
//Explanation: [4,9] is also accepted.
//
//
//
// Constraints:
//
//
// 1 <= nums1.length, nums2.length <= 1000
// 0 <= nums1[i], nums2[i] <= 1000
//
//
// Related Topics Array Hash Table Two Pointers Binary Search Sorting 👍 3915 👎
// 2027

//leetcode submit region begin(Prohibit modification and deletion)
// 解法一(hash table): 建立一個 hash table 儲存不重複的 num 在 key 中
// 並將 value 用來判斷是否已經存在於 res 之中
func intersection(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	m := make(map[int]bool, 0)
	for i := 0; i < len(nums1); i++ {
		m[nums1[i]] = false
	}
	for i := 0; i < len(nums2); i++ {
		if _, ok := m[nums2[i]]; !ok || m[nums2[i]] == true {
			continue
		}
		res = append(res, nums2[i])
		m[nums2[i]] = true
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
