package main

import "testing"

func BenchmarkPrintGreeter(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PrintGreeter("Lazada")
	}
}
