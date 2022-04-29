package main

import "fmt"

// Iterator Pattern
// 將不同資料物件透過一致的方式取得其中的元素

// 設計一個Iteratorinterface 介面，裡頭 HasNext()用來確認是否還擁有下一個元素，
// Next()用來取得元素與把元素 index 往後移。

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type IterableSliceString []string

func (i IterableSliceString) Iterator() Iterator {
	return &SliceStringIterator{
		original: i,
		index:    0,
	}
}

type SliceStringIterator struct {
	original IterableSliceString
	index    int
}

func (s *SliceStringIterator) HasNext() bool {
	return s.index < len(s.original)
}

func (s *SliceStringIterator) Next() interface{} {
	item := s.original[s.index]
	s.index++
	return item
}

type IterableString string

func (i IterableString) Iterator() Iterator {
	return &StringIterator{
		original: i,
		index:    0,
	}
}

type StringIterator struct {
	original IterableString
	index    int
}

func (s *StringIterator) HasNext() bool {
	return s.index < len(s.original)
}

func (s *StringIterator) Next() interface{} {
	item := string(s.original[s.index])
	s.index++
	return item
}

func PrintAllItems(iterator Iterator) {
	for iterator.HasNext() {
		fmt.Println(iterator.Next())
	}
}

func main() {
	PrintAllItems(IterableSliceString{"a", "b", "c"}.Iterator())
	PrintAllItems(IterableString("abcd").Iterator())
}
