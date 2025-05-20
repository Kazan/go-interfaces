package writer

import (
	"io"
)

// Writer is an interface that defines a method for writing data.
type Writer interface {
	Write(data string) error
}

// writer is a concrete implementation of the Writer interface.
// It writes data to an io.Writer.
type writer struct {
	output io.Writer
}

// NewWriter creates a new Writer that writes to the specified io.Writer.
// It returns a Writer interface.
func NewWriter(output io.Writer) Writer {
	return &writer{output: output}
}

// Write writes the given data to the output.
// It implements the Write method of the Writer interface.
func (w *writer) Write(data string) error {
	_, err := w.output.Write([]byte(data))
	return err
}
