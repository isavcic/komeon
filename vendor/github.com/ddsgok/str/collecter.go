package str

import (
	"io"
	"reflect"
)

// collected represents an array of splitted with Factory pattern
// behaviour.
type collected []splitted

// Collecter it's an array of splitted factory. It will mount sets of
// splitted through various methods, and then outputted in desired form.
type Collecter interface {
	Printer
	Array() [][]string
	AnomArray() [][]interface{}
	FmtAll(string, ...interface{}) Splitter
	JoinAll(...string) Splitter
}

// Array returns the Collecter as set of string arrays.
func (c collected) Array() (sa [][]string) {
	sa = make([][]string, len(c))
	for i := 0; i < len(c); i++ {
		sa[i] = c[i].Array()
	}
	return
}

// AnomArray returns the Collecter as set of interface{} arrays.
func (c collected) AnomArray() (sa [][]interface{}) {
	sa = make([][]interface{}, len(c))
	for i := 0; i < len(c); i++ {
		sa[i] = c[i].AnomArray()
	}
	return
}

// FmtAll applies an fmt.Sprintf with the formatted string received,
// using the different sets of strings as arguments. This is basically
// a map function to the set of strings presented on collect using a
// format string. It will return an array of all formats.
func (c collected) FmtAll(f string, args ...interface{}) (s Splitter) {
	formatter := f
	if len(args) > 0 {
		formatter = Fmt(f, args...)
	}

	arr := make([]string, len(c))
	for i := 0; i < len(c); i++ {
		arr[i] = Fmt(formatter, c[i].AnomArray()...)
	}

	s = splitted(arr)
	return
}

// JoinAll join all set of strings on the collect using a separator. It
// will return an array of all joins.
func (c collected) JoinAll(optSep ...string) (s Splitter) {
	var sep string
	if len(optSep) == 0 {
		sep = ""
	} else {
		sep = optSep[0]
	}

	arr := make([]string, len(c))
	for i := 0; i < len(c); i++ {
		arr[i] = c[i].Join(sep).String()
	}

	s = splitted(arr)
	return
}

// String returns the Collecter as a string.
func (c collected) String() (r string) {
	r = c.JoinAll().Join(" ").String()
	return
}

// Error returns a new error with Collecter as the message.
func (c collected) Error() (err error) {
	err = c.JoinAll().Join(" ").Error()
	return
}

// Print will log Collecter content to writer received, or os.Stdout as
// default writer.
func (c collected) Print(wa ...io.Writer) (n int, err error) {
	n, err = c.JoinAll().Join(" ").Print(wa...)
	return
}

// Collect creates the Collecter converting any value received to an array
// of string arrays.
func Collect(set interface{}) (c Collecter) {
	switch reflect.TypeOf(set).Kind() {
	case reflect.Slice:
		rset := reflect.ValueOf(set)

		splitArr := make([]splitted, rset.Len())
		for i := 0; i < rset.Len(); i++ {
			val := rset.Index(i).Interface()

			switch reflect.TypeOf(val).Kind() {
			case reflect.Slice:
				rval := reflect.ValueOf(val)

				strArr := make([]string, rval.Len())
				for j := 0; j < rval.Len(); j++ {
					strArr[j] = Fmt(rval.Index(j))
				}

				splitArr[i] = splitted(strArr)
			default:
				splitArr[i] = splitted([]string{Fmt(val)})
			}
		}

		c = collected(splitArr)
	default:
		c = collected([]splitted{splitted([]string{Fmt(set)})})
	}

	return
}
