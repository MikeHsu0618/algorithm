package main

import "fmt"

func main() {
	list := Constructor()
	list.AddAtHead(1)
	list.Print()
	list.AddAtHead(4)
	list.AddAtHead(1)
	list.Print()
	list.AddAtHead(2)
	list.Print()
	fmt.Println(list.Get(2))
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

func (l *MyLinkedList) Print() {
	res := make([]interface{}, 0)
	cur := l
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}
	fmt.Println(res)
}

func (l *MyLinkedList) Get(index int) int {
	head := l.Next
	for index > 0 && head != nil {
		head = head.Next
		index--
	}
	if index > 0 || head == nil {
		return -1
	}
	return head.Val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (l *MyLinkedList) AddAtHead(val int) {
	newNode := &MyLinkedList{val, l.Next}
	l.Next = newNode
}

/** Append a node of value val to the last element of the linked list. */
func (l *MyLinkedList) AddAtTail(val int) {
	tail := l
	for tail.Next != nil {
		tail = tail.Next
	}
	newNode := &MyLinkedList{val, nil}
	tail.Next = newNode
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		return
	}
	head := l
	for index > 0 && head.Next != nil {
		head = head.Next
		index--
	}
	if index > 0 {
		return
	}
	newNode := &MyLinkedList{val, head.Next}
	head.Next = newNode
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 {
		return
	}
	head := l
	for index > 0 && head.Next != nil {
		head = head.Next
		index--
	}
	if index > 0 || head.Next == nil {
		return
	}
	head.Next = head.Next.Next
}
