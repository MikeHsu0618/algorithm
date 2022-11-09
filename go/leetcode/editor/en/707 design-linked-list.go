package main

import "fmt"

//Design your implementation of the linked list. You can choose to use a singly
//or doubly linked list. A node in a singly linked list should have two
//attributes: val and next. val is the value of the current node, and next is a pointer/
//reference to the next node. If you want to use the doubly linked list, you will
//need one more attribute prev to indicate the previous node in the linked list.
//Assume all nodes in the linked list are 0-indexed.
//
// Implement the MyLinkedList class:
//
//
// MyLinkedList() Initializes the MyLinkedList object.
// int get(int index) Get the value of the indexᵗʰ node in the linked list. If
//the index is invalid, return -1.
// void addAtHead(int val) Add a node of value val before the first element of
//the linked list. After the insertion, the new node will be the first node of the
//linked list.
// void addAtTail(int val) Append a node of value val as the last element of
//the linked list.
// void addAtIndex(int index, int val) Add a node of value val before the indexᵗ
//ʰ node in the linked list. If index equals the length of the linked list, the
//node will be appended to the end of the linked list. If index is greater than the
//length, the node will not be inserted.
// void deleteAtIndex(int index) Delete the indexᵗʰ node in the linked list, if
//the index is valid.
//
//
//
// Example 1:
//
//
//Input
//["MyLinkedList", "addAtHead", "addAtTail", "addAtIndex", "get",
//"deleteAtIndex", "get"]
//[[], [1], [3], [1, 2], [1], [1], [1]]
//Output
//[null, null, null, null, 2, null, 3]
//
//Explanation
//MyLinkedList myLinkedList = new MyLinkedList();
//myLinkedList.addAtHead(1);
//myLinkedList.addAtTail(3);
//myLinkedList.addAtIndex(1, 2);    // linked list becomes 1->2->3
//myLinkedList.get(1);              // return 2
//myLinkedList.deleteAtIndex(1);    // now the linked list is 1->3
//myLinkedList.get(1);              // return 3
//
//
//
// Constraints:
//
//
// 0 <= index, val <= 1000
// Please do not use the built-in LinkedList library.
// At most 2000 calls will be made to get, addAtHead, addAtTail, addAtIndex and
//deleteAtIndex.
//
//
// Related Topics Linked List Design 👍 1876 👎 1332

//leetcode submit region begin(Prohibit modification and deletion)
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

//leetcode submit region end(Prohibit modification and deletion)
