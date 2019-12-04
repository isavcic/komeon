package str

import (
	"fmt"
	"io"
)

// Printer represents a string factory with the ability to serve it on
// different outputs.
type Printer interface {
	fmt.Stringer
	Error() error
	Print(...io.Writer) (int, error)
}
