package main

import (
	"testing"
)

// benchmarking the decode function
func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		decode("post.json")
	}
}

// benchmarking the unmarshal function
func BenchmarkUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unmarshal("post.json")
	}
}
