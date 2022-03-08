package main

import "fmt"

func main() {
    // 定義 Arr
    arr := [5]int{66, 24, 33, 55, 12}
    BubbleSort(&arr)
    fmt.Println("傳入指針排序後的陣列", arr)
}

/*
 * 核心概念：
 * 每次遞迴都將最大的數值排序到最右側, 直到將所有數值遞迴一次
 */

func BubbleSort(arr *[5]int) {
    fmt.Println("排序前的 arr =", *arr)
    temp := 0
    for i := 0; i < len(arr)-1; i++ {
        for j := 0; j < len(arr)-1-i; j++ {
            if (*arr)[j] > (*arr)[j+1] {
                // 交換
                temp = (*arr)[j]
                (*arr)[j] = (*arr)[j+1]
                (*arr)[j+1] = temp
            }
        }
    }
    fmt.Println("第一次排序的 arr =", *arr)
}
