package main

import "testing"

func BenchmarkFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fast()
	}
}

func BenchmarkSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Slow()
	}
}

func BenchmarkFastSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fast()
		Slow()
	}
}