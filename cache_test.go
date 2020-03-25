package main

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func create123List() (*Item, *Item, *Item) {
	third := Item{
		Value: 3,
		Next:  nil,
		Prev:  nil,
	}
	second := Item{
		Value: 2,
		Next:  &third,
		Prev:  nil,
	}
	first := Item{
		Value: 1,
		Next:  &second,
		Prev:  nil,
	}
	third.Prev = &second
	second.Prev = &first

	return &first, &second, &third
}

func TestList(t *testing.T) {
	t.Run("len handles lists of one element", func(t *testing.T) {
		emptyList := Item{
			Value: nil,
			Next:  nil,
			Prev:  nil,
		}

		assert.Equal(t, 1, emptyList.Len())
	})

	t.Run("Len() works with lists more then 1 element", func(t *testing.T) {
		list, _, _ := create123List()
		assert.Equal(t, 3, list.Len())
	})

	t.Run("Front() returns first element of a list", func(t *testing.T) {
		first, second, third := create123List()

		assert.Equal(t, first, third.Front())
		assert.Equal(t, first, second.Front())
		assert.Equal(t, first, first.Front())
	})

	t.Run("Back() returns last element of a list", func(t *testing.T) {
		first, second, third := create123List()

		assert.Equal(t, third, third.Back())
		assert.Equal(t, third, second.Back())
		assert.Equal(t, third, first.Back())
	})

	t.Run("PushFront() adds an element to the beginning of the list", func(t *testing.T) {
		list, _, _ := create123List()

		randInt := rand.Intn(999999)
		list.PushFront(randInt)

		assert.Equal(t, randInt, list.Front().Value)
	})

	t.Run("PushBack() adds an element to the end of the list", func(t *testing.T) {
		list, _, _ := create123List()

		randInt := rand.Intn(999999)
		list.PushBack(randInt)

		assert.Equal(t, randInt, list.Back().Value)
	})

	t.Run("Remove() works in general case", func(t *testing.T) {
		list, second, third := create123List()

		list.Remove(second)

		assert.Equal(t, 2, list.Len())
		assert.Equal(t, third, list.Back())
	})

	t.Run("Remove() works with first or last element", func(t *testing.T) {
		list, second, third := create123List()

		list.Remove(third)

		assert.Equal(t, 2, list.Len())
		assert.Equal(t, second, list.Back())
	})

	t.Run("MoveToFront() remove an element from the list", func(t *testing.T) {
		list, _, third := create123List()

		list.MoveToFront(third)

		assert.Equal(t, 3, list.Len())
		assert.Equal(t, third.Value, list.Front().Value)
	})
}