package process

import (
	"strings"

	"github.com/trustmaster/goflow"
)

type SplitWord struct {
	flow.Component
	Document <-chan string
	Words    chan<- []string
}

func (s *SplitWord) OnDocument(document string) {
	words := strings.Split(document, " ")
	s.Words <- words
}
