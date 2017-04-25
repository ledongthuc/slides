package main

import (
	"testing"

	flow "github.com/trustmaster/goflow"
)

func BenchmarkPrintGreeter(b *testing.B) {
	var net = NewGreetingApp()
	var in = make(chan string)
	net.SetInPort("In", in)
	defer close(in)

	flow.RunNet(net)

	for n := 0; n < b.N; n++ {
		PrintGreeter(in)
	}
}
