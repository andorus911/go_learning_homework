package main

import (
	"reflect"
	"testing"
)

func TestWordCounter(t *testing.T) {
	s := `Lorem (lorem) ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.`
	e := map[string]int{"lorem": 2, "ipsum": 1, "dolor": 1, "sit": 1, "amet": 1, "consectetur": 1, "adipiscing": 1, "elit": 1, "sed": 1, "do": 1, "eiusmod": 1, "tempor": 1, "incididunt": 1, "ut": 1, "labore": 1, "et": 1, "dolore": 1, "magna": 1, "aliqua": 1}
	if c := WordCounter(s); !reflect.DeepEqual(c, e) {
		t.Fatalf("wrong word counter for \"%s\": got \"%v\", expected \"%v\"", s, c, e)
	}
}

func TestTop(t *testing.T) {
	n := 3
	d := map[string]int{"foolish": 3, "tom": 4, "jerry": 7, "pie": 1}
	e := []Pair{
		{
			Word:  "foolish",
			Count: 3,
		},
		{
			Word:  "tom",
			Count: 4,
		},
		{
			Word:  "jerry",
			Count: 7,
		},
	}
	if c := Top(n, d); !reflect.DeepEqual(c, e) {
		t.Fatalf("wrong top-%v for %v: got %v, expected %v", n, d, c, e)
	}
}
