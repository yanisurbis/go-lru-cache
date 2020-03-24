package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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

		assert.Equal(t, 3, first.Len())
	})
}