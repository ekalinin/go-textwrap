package textwrap

import (
	"log"
	"strings"
)

// PredicateFunc is a function type for using in Indent function
type PredicateFunc = func(s string) bool

// IsNotEmpty returns true if s is not empty. false in other cases.
func IsNotEmpty(s string) bool {
	return strings.TrimSpace(s) != ""
}

// Any returns true for any string.
func Any(_ string) bool {
	return true
}

// None returns false for any string.
func None(_ string) bool {
	return false
}

// Indent adds 'prefix' to the beginning of selected lines in 'text'.
//
// If 'predicate' is provided, 'prefix' will only be added to the lines
// where 'predicate(line)' is True. If 'predicate' is not provided,
// it will default to adding 'prefix' to all non-empty lines that do not
// consist solely of whitespace characters.
//
// Sources:
// - https://docs.python.org/3/library/textwrap.html#textwrap.indent
// - https://github.com/python/cpython/blob/3.9/Lib/textwrap.py#L465
// - https://github.com/python/cpython/blob/3.9/Lib/test/test_textwrap.py#L816
func Indent(text, prefix string, pred PredicateFunc) string {
	prefixed := []string{}
	sep := "\n"

	predicate := pred
	if pred == nil {
		predicate = IsNotEmpty
	}

	for _, row := range strings.Split(text, sep) {
		if debug {
			log.Printf("Indent row: %q", row)
		}
		if predicate(row) {
			row = prefix + row
		}
		if debug {
			log.Printf("Indented row: %q", row)
		}
		prefixed = append(prefixed, row)
	}

	return strings.Join(prefixed, sep)
}
