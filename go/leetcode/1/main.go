package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{1, 3, 5, 7, 8}, 9))
	fmt.Println(twoSum1([]int{1, 3, 5, 7, 8}, 9))
}

// v1 版本 -> 用兩個迴圈又需要多判斷條件較不優雅
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	var res []int
	for i := 0; i < len(nums); i++ {
		m[target-nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		v, ok := m[nums[i]]
		if !ok || i == v {
			continue
		}
		return append(res, i, v)
	}
	return res
}

// v2 將兩個迴圈合而為一
func twoSum1(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if v, ok := m[another]; ok {
			return []int{v, i}
		}
		m[nums[i]] = i
	}
	return nil
}
