package main

import "sort"

//Given two integer arrays nums1 and nums2, return an array of their
//intersection. Each element in the result must appear as many times as it shows in both
//arrays and you may return the result in any order.
//
//
// Example 1:
//
//
//Input: nums1 = [1,2,2,1], nums2 = [2,2]
//Output: [2,2]
//
//
// Example 2:
//
//
//Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//Output: [4,9]
//Explanation: [9,4] is also accepted.
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
//
// Follow up:
//
//
// What if the given array is already sorted? How would you optimize your
//algorithm?
// What if nums1's size is small compared to nums2's size? Which algorithm is
//better?
// What if elements of nums2 are stored on disk, and the memory is limited such
//that you cannot load all elements into the memory at once?
//
//
// Related Topics Array Hash Table Two Pointers Binary Search Sorting 👍 5797 👎
// 807

//leetcode submit region begin(Prohibit modification and deletion)
// 使用 sort + 雙指針
func intersect(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)

	sort.Ints(nums1)
	sort.Ints(nums2)

	p1, p2 := 0, 0
	for p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] > nums2[p2] {
			p2++
			continue
		}
		if nums1[p1] < nums2[p2] {
			p1++
			continue
		}

		res = append(res, nums1[p1])
		p1++
		p2++
	}

	return res
}

// 使用 hash map 判斷是否重複
func intersect1(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	m := make(map[int]int, 0)

	for i := 0; i < len(nums1); i++ {
		m[nums1[i]]++
	}

	for i := 0; i < len(nums2); i++ {
		if v, ok := m[nums2[i]]; ok && v > 0 {
			res = append(res, nums2[i])
			m[nums2[i]]--
		}
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
