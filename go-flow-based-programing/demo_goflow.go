package main

import (
	"fmt"

	flow "github.com/trustmaster/goflow"
)

// START GREETING

type Greeter struct {
	flow.Component
	Name <-chan string // input port
	Res  chan<- string // output port
}

func (g *Greeter) OnName(name string) {
	greeting := fmt.Sprintf("Hello, %s!", name)
	g.Res <- greeting //send the greeting to the output port/
}

// END GREETING

// START PRINTER

type Printer struct {
	flow.Component
	Line <-chan string // inport
}

func (p *Printer) OnLine(line string) {
	fmt.Println(line)
}

// END PRINTER

// START NETWORKING

type GreetingApp struct {
	flow.Graph
}

func NewGreetingApp() *GreetingApp {
	// Init application
	n := new(GreetingApp)
	n.InitGraphState()

	// Add processes into application
	n.Add(new(Greeter), "greeter")
	n.Add(new(Printer), "printer")

	// Create the connection
	n.Connect("greeter", "Res", "printer", "Line")

	// Create the Application port
	n.MapInPort("In", "greeter", "Name")
	return n
}

// END NETWORKING

func PrintGreeter(in chan string) {
	in <- "Lazada"
}

// START RUN

func main() {
	// Create the application
	net := NewGreetingApp()

	// Open the application port `In`
	in := make(chan string)
	net.SetInPort("In", in)
	defer close(in)

	// Run the application and push the name
	flow.RunNet(net)
	in <- "Thuc Le"
	in <- "Lazada"

	// Wait until the app has done its job
	<-net.Wait()
}

// END RUN
