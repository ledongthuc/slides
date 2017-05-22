package process

import (
	"fmt"

	"github.com/trustmaster/goflow"
)

type Printer struct {
	flow.Component
	Message <-chan string
}

func (s *Printer) OnMessage(content string) {
	fmt.Println(content)
}
