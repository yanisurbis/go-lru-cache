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
	lengthRight := 0
	elmRight := elm

	for elmRight != nil {
		lengthRight += 1
		elmRight = elmRight.Next
	}

	lengthLeft := 0
	elmLeft := elm

	for elmLeft.Prev != nil {
		lengthLeft += 1
		elmLeft = elmLeft.Prev
	}

	return lengthRight + lengthLeft
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

func (elm *Item) PushFront(v interface{}) *Item {
	head := elm.Front()
	newElm := Item{
		Value: v,
		Next:  head,
		Prev:  nil,
	}
	head.Prev = &newElm

	return &newElm
}

func (elm *Item) PushBack(v interface{}) *Item {
	tail := elm.Back()
	newElm := Item{
		Value: v,
		Next:  nil,
		Prev:  tail,
	}
	tail.Next = &newElm

	return &newElm
}

func (elm *Item) Remove(i *Item) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
}

func (elm *Item) MoveToFront(i *Item) {
	elm.Remove(i)
	elm.PushFront(i.Value)
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
