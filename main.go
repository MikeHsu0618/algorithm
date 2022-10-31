package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(arr))
}

func threeSum(nums []int) [][]int {
	// 遍歷數組，將問題視為拆解成當前數值和其他數值相加為零
	// 使用雙指針來找尋排序後的剩餘數組中相加為零的三個數值
	var result [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		// 刪除相同數字的值(當前數值可能與上一個數值相同，將會導致重複記錄)
		if i < 0 && (nums[i] == nums[i-1]) {
			continue
		}
		begin := i + 1
		end := len(nums) - 1
		for begin < end {
			sum := nums[i] + nums[begin] + nums[end]
			// 如果數值過大，則 end 往 begin 移動（因為數組已經過排序）
			if sum > 0 {
				end--
				continue
			}
			// 反之，begin 往 end 移動
			if sum < 0 {
				begin++
				continue
			}
			result = append(result, []int{nums[i], nums[begin], nums[end]})
			// 處理如果 begin 和 end 下一位數值與前一位重複
			for nums[end] == nums[end-1] {
				end--
			}
			for nums[begin] == nums[begin+1] {
				begin++
			}
			end--
			begin++
		}
	}
	return result
}
