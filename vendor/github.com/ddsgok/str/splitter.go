package str

import (
	"io"
	"strings"
)

// splitted represents a string array type with Factory pattern
// behaviour.
type splitted []string

// Splitter it's a string array factory. It will mount strings arrays
// through various methods, and then outputted in desired form.
type Splitter interface {
	Printer
	Array() []string
	AnomArray() []interface{}
	Join(...string) Chainer
}

// Array returns the Splitter as a string array.
func (s splitted) Array() (a []string) {
	a = []string(s)
	return
}

// AnomArray returns the Splitter as a interface{} array.
func (s splitted) AnomArray() (a []interface{}) {
	a = make([]interface{}, len(s))
	for i := 0; i < len(s); i++ {
		a[i] = s[i]
	}

	return
}

// Join will mount Splitted into a Chainer using separator received.
// If nothing is received, it will use "" as separator. Uses
// strings.Join on this operation.
func (s splitted) Join(ss ...string) (c Chainer) {
	if len(ss) == 0 {
		c = chained(strings.Join([]string(s), ""))
	} else {
		c = chained(strings.Join([]string(s), ss[0]))
	}
	return
}

// String returns the Splitter as a string.
func (s splitted) String() (r string) {
	r = s.Join().String()
	return
}

// Error returns a new error with Splitter as the message.
func (s splitted) Error() (err error) {
	err = s.Join().Error()
	return
}

// Print will log Splitter content to writer received, or os.Stdout as
// default writer.
func (s splitted) Print(wa ...io.Writer) (n int, err error) {
	n, err = s.Join().Print(wa...)
	return
}

// With will take an string array and turn into a Splitter.
func With(arr []string) (s Splitter) {
	s = splitted(arr)
	return
}
