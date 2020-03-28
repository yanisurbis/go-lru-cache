package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func create123List() List {
	ll := &LinkedList{
		Head:   nil,
		Tail:   nil,
		Length: 0,
	}
	ll.PushBack(1)
	ll.PushBack(2)
	ll.PushBack(3)

	return ll
}

func createEmptyList() List {
	return &LinkedList{
		Head:   nil,
		Tail:   nil,
		Length: 0,
	}
}

func TestList(t *testing.T) {
	t.Run("Len() returns length", func(t *testing.T) {
		emptyList := createEmptyList()
		list123 := create123List()

		assert.Equal(t, 0, emptyList.Len())
		assert.Equal(t, 3, list123.Len())
	})

	t.Run("Front() returns first element of a list", func(t *testing.T) {
		emptyList := createEmptyList()
		list123 := create123List()

		assert.Nil(t, emptyList.Front())
		assert.Equal(t, 1, list123.Front().Value)
	})

	t.Run("Back() returns last element of a list", func(t *testing.T) {
		emptyList := createEmptyList()
		list123 := create123List()

		assert.Nil(t, emptyList.Back())
		assert.Equal(t, 3, list123.Back().Value)
	})

	t.Run("PushFront() adds an element to the beginning of the list", func(t *testing.T) {
		emptyList := createEmptyList()
		emptyList.PushFront("xxx")
		assert.EqualValues(t, emptyList.Front(), emptyList.Back())
		assert.Equal(t, 1, emptyList.Len())
		assert.Equal(t, "xxx", emptyList.Front().Value)

		list123 := create123List()
		list123.PushFront("xxx")
		assert.Equal(t, "xxx", list123.Front().Value)
		assert.Equal(t, 4, list123.Len())
	})

	t.Run("PushBack() adds an element to the end of the list", func(t *testing.T) {
		emptyList := createEmptyList()
		emptyList.PushBack("xxx")
		assert.EqualValues(t, emptyList.Front(), emptyList.Back())
		assert.Equal(t, 1, emptyList.Len())
		assert.Equal(t, "xxx", emptyList.Back().Value)

		list123 := create123List()
		list123.PushBack("xxx")
		assert.Equal(t, "xxx", list123.Back().Value)
		assert.Equal(t, 4, list123.Len())
	})

	t.Run("Remove() removes an element from a list", func(t *testing.T) {
		emptyList := createEmptyList()
		emptyList.Remove(emptyList.Front())
		assert.Equal(t, 0, emptyList.Len())

		list123 := create123List()
		list123.Remove(list123.Front())
		assert.Equal(t, 2, list123.Len())
		assert.Equal(t, 2, list123.Front().Value)
		assert.Equal(t, 3, list123.Back().Value)
		list123.Remove(list123.Back())
		assert.Equal(t, 1, list123.Len())
		assert.Equal(t, 2, list123.Front().Value)
		assert.Equal(t, 2, list123.Back().Value)
		list123.Remove(list123.Back())
		assert.Equal(t, 0, list123.Len())
		assert.Nil(t, list123.Front())
		assert.Nil(t, list123.Back())
	})

	t.Run("MoveToFront() moves an element to the beginning of the list", func(t *testing.T) {
		//emptyList := createEmptyList()
		//emptyList.MoveToFront(emptyList.Front())

		list123 := create123List()
		list123.MoveToFront(list123.Back())
		assert.Equal(t, 3, list123.Front().Value)
		assert.Equal(t, 2, list123.Back().Value)
	})
}