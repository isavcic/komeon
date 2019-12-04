package str

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// chained represents the central object in this package, a string type
// with Factory pattern behaviour.
type chained string

// Chainer it's a string factory. It will mount strings through various
// methods, and then outputted in desired form.
type Chainer interface {
	Printer
	Split(string) Splitter
	SplitN(string, int) Splitter
}

// Split separates string to a Splitter, an array of strings. It uses
// separator received in strings.Split to do the conversion.
func (c chained) Split(sep string) (s Splitter) {
	s = splitted(strings.Split(string(c), sep))
	return
}

// SplitN separates string to a Splitter, an array of strings. It uses
// separator received in strings.Split to do the conversion. Contains
// a limiter N to count desired number of substrings, following:
//	n > 0: at most n substrings; the last substring will be the unsplit remainder.
//	n == 0: the result is nil (zero substrings)
//	n < 0: all substrings
func (c chained) SplitN(sep string, n int) (s Splitter) {
	s = splitted(strings.SplitN(string(c), sep, n))
	return
}

// String returns the Chainer as a string.
func (c chained) String() (s string) {
	s = string(c)
	return
}

// Error returns a new error with Chainer as the message.
func (c chained) Error() (err error) {
	err = errors.New(string(c))
	return
}

// Print will log Chainer content to writer received, or os.Stdout as
// default writer.
func (c chained) Print(optWt ...io.Writer) (n int, err error) {
	var wt io.Writer = os.Stdout

	if len(optWt) != 0 {
		wt = optWt[0]
	}

	n, err = fmt.Fprint(wt, string(c))
	return
}

// New creates the Chainer using fmt.Sprintf to ensures string
// formatting.
func New(s interface{}, args ...interface{}) (c Chainer) {
	c = chained(Fmt(s, args...))
	return
}

// Fmt creates the Chainer using fmt.Sprintf and already returns the
// string value.
func Fmt(formatter interface{}, args ...interface{}) (s string) {
	if len(args) > 0 {
		s = fmt.Sprintf(Fmt(formatter), args...)
	} else {
		s = fmt.Sprintf("%v", formatter)
	}

	return
}
