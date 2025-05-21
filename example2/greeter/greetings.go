package greeter

import (
	"fmt"
)

type MyWriter interface {
	Write(data string) error
}

type MyColorizer interface {
	Color() string
}

// Greeter is a struct that holds a name and a writer.
type Greeter struct {
	name   string
	writer MyWriter
	color  MyColorizer
}

// NewGreeter creates a new Greeter with the specified name and writer.
func NewGreeter(name string, w MyWriter, c MyColorizer) *Greeter {
	return &Greeter{
		name:   name,
		writer: w,
		color:  c,
	}
}

// Greet writes a greeting message to the writer.
func (g *Greeter) Greet() {
	g.writer.Write(fmt.Sprintf("Hello, %s! Your color is %s\n", g.name, g.color.Color()))
}
