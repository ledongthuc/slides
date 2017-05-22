package graph

import (
	"basic-demo/process"

	flow "github.com/trustmaster/goflow"
)

type WordProcessor struct {
	flow.Graph
}

func NewWordProcessor() *WordProcessor {
	n := new(WordProcessor)
	n.InitGraphState()

	n.Add(new(process.WordCounter), "WordCounter")
	n.Add(new(process.SplitWord), "SplitWord")
	n.Add(new(process.Printer), "Printer")

	n.Connect("SplitWord", "Words", "WordCounter", "Words")
	n.Connect("WordCounter", "MessageCounter", "Printer", "Message")

	n.MapInPort("In", "SplitWord", "Document")

	return n
}
