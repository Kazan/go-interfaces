package greeter

import (
	"github.com/kazan/go-interfaces/example1/writer"
)

// Greeter is a struct that holds a name and a writer.
type Greeter struct {
	name   string
	writer writer.Writer
}

// NewGreeter creates a new Greeter with the specified name and writer.
func NewGreeter(name string, w writer.Writer) *Greeter {
	return &Greeter{name: name, writer: w}
}

// Greet writes a greeting message to the writer.
func (g *Greeter) Greet() {
	g.writer.Write("Hello, " + g.name + "\n")
}
