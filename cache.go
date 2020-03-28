package main

import "fmt"

type Item struct {
	Value interface{}
	Next *Item
	Prev *Item
}

type LinkedList struct {
	Head   *Item
	Tail   *Item
	Length int
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

func (l *LinkedList) Len() int {
	return l.Length
}

func (l *LinkedList) Front() *Item {
	return l.Head
}

func (l *LinkedList) Back() *Item {
	return l.Tail
}

func (l *LinkedList) PushFront(v interface{}) *Item {
	if l.Length == 0 {
		newElm := &Item{
			Value: v,
			Next:  nil,
			Prev:  nil,
		}
		l.Head = newElm
		l.Tail = newElm
		l.Length = 1

		return newElm
	} else {
		head := l.Head
		newElm := &Item{
			Value: v,
			Next:  head,
			Prev:  nil,
		}
		head.Prev = newElm
		l.Head = newElm
		l.Length += 1

		return newElm
	}
}

func (l *LinkedList) PushBack(v interface{}) *Item {
	if l.Length == 0 {
		newElm := &Item{
			Value: v,
			Next:  nil,
			Prev:  nil,
		}
		l.Head = newElm
		l.Tail = newElm
		l.Length = 1

		return newElm
	} else {
		tail := l.Tail
		newElm := &Item{
			Value: v,
			Next:  nil,
			Prev:  tail,
		}
		tail.Next = newElm
		l.Tail = newElm
		l.Length += 1

		return newElm
	}
}

func (l *LinkedList) Remove(i *Item) {
	// QUESTION: how to make sure the list contains this particular item?
	if i == nil || l.Length == 0 {
		return
	}

	if i.Prev != nil && i.Next != nil {
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
		l.Length -= 1
		return
	}

	if i.Prev == nil && i.Next != nil {
		l.Head = i.Next
		l.Head.Prev = nil
		l.Length -= 1
		return
	}

	if i.Prev != nil && i.Next == nil {
		l.Tail = i.Prev
		l.Tail.Next = nil
		l.Length -= 1
		return
	}

	if i.Prev == nil && i.Next == nil {
		l.Head = nil
		l.Tail = nil
		l.Length = 0
		return
	}
}

func (l *LinkedList) MoveToFront(i *Item) {
	if i == nil {
		return
	}
	l.Remove(i)
	l.PushFront(i.Value)
}

type Cache interface {
	Set(key string, value interface{}) bool
	Get(key string) (interface{}, bool)
	Clear()
}

type LRUCache struct {
	size int
	elements List
	dict map[string]*Item
}

func (elm *Item) Set(key string, value interface{}) {

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
