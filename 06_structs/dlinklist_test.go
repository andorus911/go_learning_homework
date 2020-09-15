package main

import "testing"

// List tests

func TestLen(t *testing.T) {
	e := 3
	l := new(List)
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	if g := l.Len(); g != e {
		t.Fatalf("wrong list length for %v: got %v, expected %v", l, g, e)
	}
}

func TestFirst(t *testing.T) {
	l := new(List)
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	e := l.first
	if g := l.First(); g != *e {
		t.Fatalf("wrong first item for %v: got %v, expected %v", l, g, e)
	}
}

func TestLast(t *testing.T) {
	l := new(List)
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	e := l.last
	if g := l.Last(); g != *e {
		t.Fatalf("wrong last item for %v: got %v, expected %v", l, g, e)
	}
}

func TestPushFront(t *testing.T) {
	l := new(List)
	l.PushFront(1)
	l.PushFront(2)
	e := 3
	l.PushFront(e)
	if g := l.first; g.value != e {
		t.Fatalf("wrong front item after front push for %v: got %v, expected %v", l, g, e)
	}
}

func TestPushBack(t *testing.T) {
	l := new(List)
	l.PushBack(1)
	l.PushBack(2)
	e := 3
	l.PushBack(e)
	if g := l.last; g.value != e {
		t.Fatalf("wrong last item after back push for %v: got %v, expected %v", l, g, e)
	}
}

func TestRemove(t *testing.T) {
	l := new(List)
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	e := 2
	l.Remove(*l.first)
	if g := l.Len(); g != e {
		t.Fatalf("wrong length for %v after removing: got %v, expected %v", l, g, e)
	}
}

// Item tests

func TestValue(t *testing.T) {
	i := new(Item)
	e := 1
	i.value = e

	if g := i.Value(); g != e {
		t.Fatalf("wrong value of item %v: got %v, expected %v", i, g, e)
	}
}

func TestNext(t *testing.T) {
	l := new(List)
	l.PushBack(1)
	e := 5
	l.PushBack(e)
	l.PushBack(3)

	if g := l.first.Next(); g.value != e {
		t.Fatalf( "wrong next item value for %v: got %v, expected %v", l, g, e)
	}
}

func TestPrev(t *testing.T) {
	l := new(List)
	l.PushBack(1)
	e := 5
	l.PushBack(e)
	l.PushBack(3)

	if g := l.last.Prev(); g.value != e {
		t.Fatalf("wrong previous item value for %v: got %v, expected %v", l, g, e)
	}
}