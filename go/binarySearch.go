package main

/*
 * 二分查找思路分析:
 * 1. arr 必須是一個有序的數組, 並且是從小到大排序
 * 2. 先找到中間的下標 (middle = (leftIndex + rightIndex) / 2), 然後讓中機下標的值和 findVal 進行比較
 * 預期會產生三種可能
 * 2.1 如果 arr[middle] > findVal, 就應該相 leftindex ---- (middle - 1)
 * 2.2 如果 arr[middle] < findVal, 就應該相 rightindex ---- (middle + 1)
 * 2,3 如果 arr[middle] == findVal , 就算直接找到解答
 */
func main() {

}
