package main

import "fmt"

type Item struct {
	Value interface{}
	Next *Item
	Prev *Item
}

type List interface {
	Len() int
	Front() *Item
	Back() *Item
	PushFront(v interface{}) *Item
	PushBack(v interface{}) *Item
	Remove(i *Item)
	MoveToFront(i *Item)
}

func (elm *Item) Len() int {
	length := 0
	currElm := elm

	for currElm != nil {
		length += 1
		currElm = currElm.Next
	}

	return length
}

func (elm *Item) Front() *Item {
	currElm := elm

	for currElm.Prev != nil {
		currElm = currElm.Prev
	}

	return currElm
}

func (elm *Item) Back() *Item {
	currElm := elm

	for currElm.Next != nil {
		currElm = currElm.Next
	}

	return currElm
}

func main() {
	fmt.Println(Item{
		Value: "Hello",
		Next:  &Item{
			Value: "World",
			Next:  nil,
			Prev:  nil,
		},
		Prev:  nil,
	})
}
