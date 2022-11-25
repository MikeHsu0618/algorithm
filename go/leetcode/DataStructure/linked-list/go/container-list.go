package main

import (
	"container/list"
	"fmt"
)

// type Element
//
//type Element struct {
//
//        Value interface{}    //在元素中存儲的值
//}
//
//func (e *Element) Next() *Element //返回該元素的下一個元素，如果沒有下一個元素則返回nil
//func (e *Element) Prev() *Element//返回該元素的前一個元素，如果沒有前一個元素則返回nil。
//type List
//func New() *List //返回一個初始化的list
//func (l *List) Back() *Element //獲取list l的最後一個元素
//func (l *List) Front() *Element //獲取list l的第一個元素
//func (l *List) Init() *List //list l初始化或者清除list l
//func (l *List) InsertAfter(v interface{}, mark *Element) *Element //在list l中元素mark之後插入一個值為v的元素，並返回該元素，如果mark不是list中元素，則list不改變。
//func (l *List) InsertBefore(v interface{}, mark *Element) *Element//在list l中元素mark之前插入一個值為v的元素，並返回該元素，如果mark不是list中元素，則list不改變。
//func (l *List) Len() int //獲取list l的長度
//func (l *List) MoveAfter(e, mark *Element) //將元素e移動到元素mark之後，如果元素e或者mark不屬於list l，或者e==mark，則list l不改變。
//func (l *List) MoveBefore(e, mark *Element)//將元素e移動到元素mark之前，如果元素e或者mark不屬於list l，或者e==mark，則list l不改變。
//func (l *List) MoveToBack(e *Element)//將元素e移動到list l的末尾，如果e不屬於list l，則list不改變。
//func (l *List) MoveToFront(e *Element)//將元素e移動到list l的首部，如果e不屬於list l，則list不改變。
//func (l *List) PushBack(v interface{}) *Element//在list l的末尾插入值為v的元素，並返回該元素。
//func (l *List) PushBackList(other *List)//在list l的尾部插入另外一個list，其中l和other可以相等。
//func (l *List) PushFront(v interface{}) *Element//在list l的首部插入值為v的元素，並返回該元素。
//func (l *List) PushFrontList(other *List)//在list l的首部插入另外一個list，其中l和other可以相等。
//func (l *List) Remove(e *Element) interface{}//如果元素e屬於list l，將其從list中刪除，並返回元素e的值。
//————————————————

func main() {
	l := list.New() // 創建一個新的 list
	for i := 0; i < 5; i++ {
		l.PushBack(i) // 將數值推到尾部
	}
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value) // 遍歷輸出 list : 0,1,2,3,4
	}
	fmt.Println("輸出首個元素", l.Front().Value)
	fmt.Println("輸出末個元素", l.Back().Value)
	l.InsertAfter(6, l.Front()) // 在首個元素後插入一個值為 6 的元素
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value) // 遍歷輸出 list : 0,6,1,2,3,4
	}
	fmt.Println("")
	l.MoveBefore(l.Front().Next(), l.Front()) // 將首兩個元素位置互換
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value) // 遍歷輸出 list : 6,0,1,2,3,4
	}
	fmt.Println("")
	l.MoveToFront(l.Back()) // 將末個元素移動到首部
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value) // 遍歷輸出 list : 4,6,0,1,2,3
	}
	fmt.Println("")
	l2 := list.New()
	l2.PushBackList(l) // 將 l 放到 l2 後面
	for e := l2.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value) // 遍歷輸出 list : 4,6,0,1,2,3
	}

	l.Init() // 清空 list
	fmt.Println("此時 list 的長度為 ", l.Len())
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value) // 遍歷輸出 list : 無內容
	}
}
