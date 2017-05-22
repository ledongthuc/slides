package main

import (
	"basic-demo/graph"

	flow "github.com/trustmaster/goflow"
)

func main() {
	network := graph.NewWordProcessor()
	in := make(chan string)
	defer close(in)
	network.SetInPort("In", in)
	flow.RunNet(network)
	in <- "Hanna"

	<-network.Wait()
}
