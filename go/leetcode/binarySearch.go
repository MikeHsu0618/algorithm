package main

import "fmt"

// BinarySearch
// * 二分查找思路分析:
// * 1. arr 必須是一個有序的數組, 並且是從小到大排序
// * 2. 先找到中間的下標 (middle = (leftIndex + rightIndex) / 2), 然後讓中機下標的值和 findVal 進行比較
// * 預期會產生三種可能
// * 2.1 如果 arr[middle] > findVal, 就應該相 leftindex ---- (middle - 1)
// * 2.2 如果 arr[middle] < findVal, 就應該相 rightindex ---- (middle + 1)
// * 2,3 如果 arr[middle] == findVal , 就算直接找到解答
// * 3. 如果 leftIndex > rightIndex 則為找不到
//
func BinarySearch(arr *[6]int, leftIndex int, rightIndex int, findVal int) {
	// 判斷 leftIndex 是否大於 rightIndex
	if leftIndex > rightIndex {
		fmt.Println("找不到了壓 哥")
		return
	}
	// 先找到中間的下標 (int 除法會自動省略小數點)
	middle := (leftIndex + rightIndex) / 2
	fmt.Printf("l: %v r: %v mid: %v value: %v\n", leftIndex, rightIndex, middle, arr[middle])
	if (*arr)[middle] > findVal {
		// 要查找的數在左邊
		BinarySearch(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		// 要查找的數在右邊
		BinarySearch(arr, middle+1, rightIndex, findVal)
	} else {
		fmt.Printf("找到拉太神了 key 為 %v\n", middle)
	}
}

func main() {
	arr := [6]int{12, 24, 33, 55, 77, 78}
	BinarySearch(&arr, 0, len(arr)-1, 55)
}
