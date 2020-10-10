package reuse

import "encoding/json"

// TODO(a.telyshev): Profile

type A struct {
	I int
}

func (a *A) Reset() {
	*a = A{}
}

func Slow() {
	for i := 0; i < 1000; i++ {
		a := &A{}
		_ = json.Unmarshal([]byte("{\"i\": 32}"), a)
	}
}

func Fast() {
	a := &A{}
	for i := 0; i < 1000; i++ {
		a.Reset()
		_ = json.Unmarshal([]byte("{\"i\": 32}"), a)
	}
}
