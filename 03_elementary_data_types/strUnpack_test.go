package main

import "testing"

func TestUnpackString(t *testing.T) {
	s := []string{"a4bc2d5e", "abcd", "45", `qwe\4\5`, `qwe\45`, `qwe\\5`}
	e := []string{"aaaabccddddde", "abcd", "", `qwe45`, `qwe44444`, `qwe\\\\\`}
	for i := range s {
		if c := UnpackString(s[i]); c != e[i] {
			t.Fatalf("wrong unpacking for \"%s\": got \"%s\", expected \"%s\"", s[i], c, e[i])
		}
	}
}
