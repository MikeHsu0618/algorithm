package main

import "fmt"

func main() {
	list := Constructor()
	list.Print()
	//list.AddAtTail(3)
	//list.Print()
	//list.AddAtIndex(0, 6)
	//list.AddAtIndex(0, 6)
	//list.AddAtIndex(3, 3)
	//list.AddAtIndex(2, 2)
	//list.AddAtHead(5)
	//list.Print()
	//list.AddAtHead(1)
	//list.Print()
	list.AddAtHead(1)
	list.Print()
	//fmt.Println(list.Get(3))
	//fmt.Println(list.Get(0))
	//fmt.Println(list.Get(1))
	//fmt.Println(list.Get(2))
	list.AddAtTail(3)
	list.Print()
	//list.AddAtTail(9)
	//list.Print()
	list.AddAtIndex(0, 6)
	list.AddAtIndex(3, 2)
	list.Print()
	fmt.Println(list.Get(0))

	fmt.Println(list.Get(1))
	fmt.Println(list.Get(2))
	list.AddAtIndex(4, 8)

	//list.AddAtIndex(2, 22)
	list.Print()
	//list.AddAtIndex(7, 33)
	//list.Print()
	//list.DeleteAtIndex(0)
	//list.Print()
	//list.DeleteAtIndex(1)
	//list.Print()
	//list.DeleteAtIndex(1)
	//list.Print()
	//list.DeleteAtIndex(1)
	//list.Print()
	//list.DeleteAtIndex(1)
	//list.Print()
	//list.DeleteAtIndex(11)
	//list.Print()
}

//[0 1]
//[0 1 4 1]
//[0 2 1 4 1]
// 此寫法將會預設一個 dummy node 來方便操作
type MyLinkedList struct {
	Val  int
	Next *MyLinkedList
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Print() {
	res := make([]interface{}, 0)
	cur := this
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}
	fmt.Println(res)
}

func (this *MyLinkedList) Get(index int) int {
	cur := this.Next
	for index > 0 && cur != nil {
		cur = cur.Next
		index--
	}
	if index > 0 || cur == nil {
		return -1
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.Next = &MyLinkedList{Val: val, Next: this.Next}
}

func (this *MyLinkedList) AddAtTail(val int) {
	cur := this
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &MyLinkedList{Val: val}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	cur := this
	// 這裡的 cur != nil 為了滿足 index == len(list) 在 list 尾端插入 node
	for index > 0 && cur != nil {
		cur = cur.Next
		index--
	}
	// index == len(list) 情況
	if index > 0 || cur == nil {
		return
	}
	cur.Next = &MyLinkedList{Val: val, Next: cur.Next}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	cur := this
	for index > 0 && cur.Next != nil {
		cur = cur.Next
		index--
	}
	// 處理不存在的 index
	if index > 0 || cur.Next == nil {
		return
	}
	cur.Next = cur.Next.Next
}
