package main

import (
	"fmt"
)

type List struct {
	first *Item
	last  *Item
}

func (l List) Len() int {
	var counter int
	item := l.first
	for {
		if item != nil {
			counter++
			item = item.next
			// what about circle list?
		} else {
			return counter
		}
	}
}

func (l List) First() Item {
	return *l.first
}

func (l List) Last() Item {
	return *l.last
}

func (l *List) PushFront(v interface{}) {
	i := &Item{
		prev: nil,
		value: v,
		next:  l.first,
	}
	if l.first != nil {
		l.first.prev = i
	}
	l.first = i
	if l.last == nil {
		l.last = l.first
	}
}

func (l *List) PushBack(v interface{}) {
	i := Item{
		prev:  l.last,
		value: v,
		next: nil,
	}
	if l.last != nil {
		l.last.next = &i
	}
	l.last = &i
	if l.first == nil {
		l.first = l.last
	}
}

func (l *List) Remove(i Item) {
	if *l.first == i {
		i.next.prev = nil
		l.first = i.next
		return
	}
	if *l.last == i {
		i.prev.next = nil
		l.last = i.prev
		return
	}
	if *l.first == i && *l.last == i {
		l.first = nil
		l.last = nil
	} else {
		i.prev.next, i.next.prev = i.next, i.prev
	}
}

type Item struct {
	prev  *Item
	value interface{}
	next  *Item
}

func (i Item) Value() interface{} {
	return i.value
}

func (i Item) Next() *Item {
	return i.next
}

func (i Item) Prev() *Item {
	return i.prev
}

func main() {
	fmt.Println("Creating a list.")
	l := new(List)
	fmt.Println("Length of list:", l.Len())
	fmt.Println("Adding 1, 2, 3")
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	fmt.Println("Length of list:", l.Len())
	fmt.Println("First item:", l.First())
	fmt.Println("Last item:", l.Last())
	i := l.Last()
	i = *i.Prev()
	fmt.Println("Previous item value:", i.Value())
	fmt.Println("Removing the item.")
	l.Remove(*l.first)
	fmt.Println("Length of list:", l.Len())
}
