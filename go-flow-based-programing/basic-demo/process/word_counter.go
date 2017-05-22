package process

import (
	"fmt"

	"github.com/trustmaster/goflow"
)

type WordCounter struct {
	flow.Component
	Words          <-chan []string
	MessageCounter chan<- string
}

func (s *WordCounter) OnWords(words []string) {
	numberOfWords := len(words)
	s.MessageCounter <- fmt.Sprintf("Number of words: %d", numberOfWords)
}
