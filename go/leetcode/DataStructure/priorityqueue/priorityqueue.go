package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value    string // Item 的內容(可自訂)
	priority int    // 優先級判斷的依據(可自訂)
	index    int    // index 為 head.Interface 需要的欄位
}

type PriorityQueue []*Item

func (pq *PriorityQueue) Pop() any {
	res := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return res
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Len() int {
	return len(*pq)
}

// 這裡的 Less 判斷可以使其成為 Max-heap
func (pq *PriorityQueue) Less(i, j int) bool {
	return (*pq)[i].priority > (*pq)[j].priority
}

func (pq *PriorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
	(*pq)[i].index = i
	(*pq)[j].index = j
}

func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value, item.priority = value, priority
	heap.Fix(pq, item.index)
}

func main() {
	// Some items and their priorities.
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	// Insert a new item and then modify its priority.
	item := &Item{
		value:    "orange",
		priority: 1,
	}
	heap.Push(&pq, item)
	pq.update(item, item.value, 5)

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%s ", item.priority, item.value)
	}
}
