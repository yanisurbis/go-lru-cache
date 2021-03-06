package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func create123List() List {
	ll := NewList()
	ll.PushBack(1)
	ll.PushBack(2)
	ll.PushBack(3)

	return ll
}

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, l.Len(), 0)
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, l.Len(), 3)

		middle := l.Back().Prev // 20
		l.Remove(middle)        // [10, 30]
		require.Equal(t, l.Len(), 2)

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, l.Len(), 7)
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Back(); i != nil; i = i.Prev {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{50, 30, 10, 40, 60, 80, 70}, elems)
	})

	t.Run("Len() returns length", func(t *testing.T) {
		emptyList := NewList()
		list123 := create123List()

		require.Equal(t, 0, emptyList.Len())
		require.Equal(t, 3, list123.Len())
	})

	t.Run("Front() returns first element of a list", func(t *testing.T) {
		emptyList := NewList()
		list123 := create123List()

		require.Nil(t, emptyList.Front())
		require.Equal(t, 1, list123.Front().Value)
	})

	t.Run("Back() returns last element of a list", func(t *testing.T) {
		emptyList := NewList()
		list123 := create123List()

		require.Nil(t, emptyList.Back())
		require.Equal(t, 3, list123.Back().Value)
	})

	t.Run("PushFront() adds an element to the beginning of the list", func(t *testing.T) {
		emptyList := NewList()
		emptyList.PushFront("xxx")
		require.EqualValues(t, emptyList.Front(), emptyList.Back())
		require.Equal(t, 1, emptyList.Len())
		require.Equal(t, "xxx", emptyList.Front().Value)

		list123 := create123List()
		list123.PushFront("xxx")
		require.Equal(t, "xxx", list123.Front().Value)
		require.Equal(t, 4, list123.Len())
	})

	t.Run("PushBack() adds an element to the end of the list", func(t *testing.T) {
		emptyList := NewList()
		emptyList.PushBack("xxx")
		require.EqualValues(t, emptyList.Front(), emptyList.Back())
		require.Equal(t, 1, emptyList.Len())
		require.Equal(t, "xxx", emptyList.Back().Value)

		list123 := create123List()
		list123.PushBack("xxx")
		require.Equal(t, "xxx", list123.Back().Value)
		require.Equal(t, 4, list123.Len())
	})

	t.Run("Remove() removes an element from a list", func(t *testing.T) {
		emptyList := NewList()
		emptyList.Remove(emptyList.Front())
		require.Equal(t, 0, emptyList.Len())

		list123 := create123List()
		list123.Remove(list123.Front())
		require.Equal(t, 2, list123.Len())
		require.Equal(t, 2, list123.Front().Value)
		require.Equal(t, 3, list123.Back().Value)
		list123.Remove(list123.Back())
		require.Equal(t, 1, list123.Len())
		require.Equal(t, 2, list123.Front().Value)
		require.Equal(t, 2, list123.Back().Value)
		list123.Remove(list123.Back())
		require.Equal(t, 0, list123.Len())
		require.Nil(t, list123.Front())
		require.Nil(t, list123.Back())
	})

	t.Run("MoveToFront() moves an element to the beginning of the list", func(t *testing.T) {
		//emptyList := NewList()
		//emptyList.MoveToFront(emptyList.Front())

		list123 := create123List()
		list123.MoveToFront(list123.Back())
		require.Equal(t, 3, list123.Front().Value)
		require.Equal(t, 2, list123.Back().Value)
	})
}

func TestCache(t *testing.T) {
	t.Run("Cache works", func(t *testing.T) {
		cache := NewCache(2)
		cache.Set("1", 1)
		cache.Set("2", 2)
		cache.Set("3", 3)

		_, found1 := cache.Get("1")
		require.False(t, found1)

		v2, found2 := cache.Get("2")
		require.True(t, found2)
		require.Equal(t, 2, v2)

		v3, found3 := cache.Get("3")
		require.True(t, found3)
		require.Equal(t, 3, v3)

		cache.Clear()
		_, found4 := cache.Get("3")
		require.False(t, found4)
	})

	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})
}