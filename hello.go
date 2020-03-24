package main

import "fmt"

type Item struct {
	Value interface{}
	Next *Item
	Prev *Item
}

//type List interface {
//	Len() int
//}

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
