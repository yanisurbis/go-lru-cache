package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func createBasicList() *Item {
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

	return &first
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

	t.Run("len works with lists more then 1 element", func(t *testing.T) {
		list := createBasicList()
		assert.Equal(t, 3, list.Len())
	})

	t.Run("front returns first element of a list", func(t *testing.T) {
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

		assert.Equal(t, &first, third.Front())
		assert.Equal(t, &first, second.Front())
		assert.Equal(t, &first, first.Front())
	})

	t.Run("back returns last element of a list", func(t *testing.T) {
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

		assert.Equal(t, &third, third.Back())
		assert.Equal(t, &third, second.Back())
		assert.Equal(t, &third, first.Back())
	})
}