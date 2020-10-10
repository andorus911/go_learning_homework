package syncpool

import (
	"strings"
	"sync"
)

// TODO(a.telyshev): Profile

var pool = sync.Pool{
	New: func() interface{} {
		return new(strings.Builder)
	},
}

func Slow() (res string) {
	for i := 0; i < 100; i++ {
		builder := new(strings.Builder)
		builder.WriteString("Hello")
		res = builder.String()
	}
	return
}

func Fast() (res string) {
	for i := 0; i < 100; i++ {
		builder := pool.Get().(*strings.Builder)
		builder.Reset()
		builder.WriteString("Hello")
		res = builder.String()
		pool.Put(builder) // Try to comment it out!
	}
	return
}
