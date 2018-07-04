package ksuid

import (
	"testing"
	)

func BenchmarkStandardGenerator_Next(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Next()
	}
}

func BenchmarkAsyncGenerator_Next(b *testing.B) {
	g := NewAsyncGenerator()
	go g.Run()
	
	for i := 0; i < b.N; i++ {
		g.Next()
	}
}