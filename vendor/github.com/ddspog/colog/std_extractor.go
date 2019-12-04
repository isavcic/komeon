package colog

import (
	"bytes"
	"regexp"

	"github.com/ddsgok/str"
)

var (
	keyPtn = `[\pL\d_]+`
	valuePtn = `[\pL\pN\\\/:.\-_'"(){}<>@ ]+`
	allButApostrophe  = `[^']+`
	allButQuote = `[^="]+`
	allButEqual = `[^=]+`

	// regex to extract key-value (or quoted value) from the logged message
	// if you can do this better please make a pull request
	// this is just the result of lots of trial and error
	fieldsRegex = str.Fmt(
		`(?P<key>%[1]s)\s*=\s*(?P<value>%[2]s|"%[3]s"|'%[4]s'|{%[5]s}|&(?:%[1]s|{%[5]s}))(?:\s+|$)`,
		keyPtn, valuePtn, allButQuote,allButApostrophe, allButEqual,
	)
)

// StdExtractor implements a regex based extractor for key-value pairs
// both unquoted foo=bar and quoted foo="some bar" are supported
type StdExtractor struct {
	rxFields *regexp.Regexp
}

// Extract finds key-value pairs in the message and sets them as Fields
// in the entry removing the pairs from the message.
func (se *StdExtractor) Extract(e *Entry) error {
	if se.rxFields == nil {
		se.rxFields = regexp.MustCompile(fieldsRegex)
	}
	matches := se.rxFields.FindAllSubmatch(e.Message, -1)
	if matches == nil {
		return nil
	}

	var key, value []byte
	captures := make(map[string]interface{})

	// Look for positions with: fmt.Printf("%#v \n", rxFields.SubexpNames())
	// Will find positions []string{"", "key", "", "", "value", "", "quoted", ""}
	//                                    1               4            6

	for _, match := range matches {
		// First group, simple key-value detected
		if len(match[1]) > 0 && len(match[2]) > 0 {
			key, value = match[1], match[2]
		}

		captures[string(key)] = string(value)
	}

	if captures != nil {
		// Eliminate key=value from text and trim from the right
		e.Message = bytes.TrimRight(se.rxFields.ReplaceAll(e.Message, nil), " \n")
		for k, v := range captures {
			e.Fields[k] = v
		}
	}

	return nil
}
