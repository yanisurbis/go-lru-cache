package main

import "fmt"

type listItem struct {
	Value interface{}
	Next *listItem
	Prev *listItem
}

type list struct {
	Head   *listItem
	Tail   *listItem
	Length int
}

type List interface {
	Len() int
	Front() *listItem
	Back() *listItem
	PushFront(v interface{}) *listItem
	PushBack(v interface{}) *listItem
	Remove(i *listItem)
	MoveToFront(i *listItem)
}

func (l *list) Len() int {
	return l.Length
}

func (l *list) Front() *listItem {
	return l.Head
}

func (l *list) Back() *listItem {
	return l.Tail
}

func (l *list) PushFront(v interface{}) *listItem {
	if l.Length == 0 {
		newElm := &listItem{
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
		newElm := &listItem{
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

func (l *list) PushBack(v interface{}) *listItem {
	if l.Length == 0 {
		newElm := &listItem{
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
		newElm := &listItem{
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

func (l *list) Remove(i *listItem) {
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

func (l *list) MoveToFront(i *listItem) {
	if i == nil {
		return
	}
	l.Remove(i)
	l.PushFront(i.Value)
}

func NewList() List {
	return &list{
		Head:   nil,
		Tail:   nil,
		Length: 0,
	}
}

type Cache interface {
	Set(key string, value interface{}) bool
	Get(key string) (interface{}, bool)
	Clear()
}

type lruCache struct {
	Size int
	Queue List
	Elements map[string]*listItem
}

func (cache *lruCache) Set(key string, value interface{}) bool {
	element, found := cache.Elements[key]
	if found {
		cache.Queue.Remove(element)
		cache.Elements[key] = cache.Queue.PushFront(value)
		return true
	} else {
		if cache.Queue.Len() == cache.Size {
			// delete element
			elmToDelete := cache.Queue.Back()
			cache.Queue.Remove(elmToDelete)
			for key, elm := range cache.Elements {
				if elm == elmToDelete {
					delete(cache.Elements, key)
				}
			}
		}
		elm := cache.Queue.PushFront(value)
		cache.Elements[key] = elm
		return false
	}
}

func (cache *lruCache) Get(key string) (interface{}, bool) {
	element, found := cache.Elements[key]
	if found {
		cache.Queue.MoveToFront(element)
		return element.Value, true
	} else {
		return nil, false
	}
}

func (cache *lruCache) Clear() {
	cache.Queue = NewList()
	cache.Elements = make(map[string]*listItem, cache.Size)
}

func NewCache(size int) Cache {
	return &lruCache{
		Size:     size,
		Queue:    NewList(),
		Elements: make(map[string]*listItem, size),
	}
}

func main() {
	fmt.Println(listItem{
		Value: "Hello",
		Next:  &listItem{
			Value: "World",
			Next:  nil,
			Prev:  nil,
		},
		Prev:  nil,
	})
}
