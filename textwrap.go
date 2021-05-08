package textwrap

import (
	"bufio"
	"strings"
)

const Version = "0.0.1"

var debug bool

// SetDebug sets debug output on/off.
func SetDebug(d bool) {
	debug = d
}

type zipped struct {
	s1 rune
	s2 rune
}

// zip returns a slice of zipped structs, where the i-th elemnt contains the
// i-th element from each of the argument strings. It stops when the shortest
// input string is exhausted.
func zip(s1, s2 string) []zipped {
	r1 := bufio.NewReader(strings.NewReader(s1))
	r2 := bufio.NewReader(strings.NewReader(s2))

	var res []zipped
	for {
		c1, _, err1 := r1.ReadRune()
		c2, _, err2 := r2.ReadRune()
		if err1 != nil || err2 != nil {
			break
		}
		res = append(res, zipped{c1, c2})

	}
	return res
}
