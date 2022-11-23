package main

import (
	"container/heap"
	"fmt"
)

// 先來看一個簡單例子，基於整型integer 實現一個最小堆。
//
// 首先定義一個自己的類型，在這個例子中是int，所以這一步跳過；
// 定義一個Heap 類型，這裡我們使用type IntHeap []int；
// 實現自定義Heap 類型的5 個方法，三個sort 的，加上Push 和Pop。

//  Heap Interface
// type Interface interface {
// 	sort.Interface
// 	Push(x any) //  add x as element Len()
// 	Pop() any   //  remove and return element Len() - 1.
// }
// Sort Interface
// type Interface interface {
// 	Len() int
// 	Less(i, j int) bool
// 	Swap(i, j int)
// }

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h *IntHeap) Push(val any) {
	*h = append(*h, val.(int))
}

func (h *IntHeap) Pop() any {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

func (h *IntHeap) Len() int {
	return len(*h)
}

// i = next j = current
func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func main() {
	h := &IntHeap{1, 6, 2, 3}
	heap.Init(h)
	heap.Push(h, 2)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	// minimum: 1
	// 1 2 2 3 6
}
