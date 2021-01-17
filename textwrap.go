package textwrap

import (
	"bufio"
	"log"
	"regexp"
	"strings"
)

func isSpace(r rune) bool {
	return r == '\t' || r == ' '
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

var (
	leadSpaceRe = regexp.MustCompile(`(^[ \t]*)([^ \t\n])`)
	blankRowsRe = regexp.MustCompile(`(?m)^[ \t]*$`)
)

// Dedent removes any common leading whitespace from every line in `text`.
//
// This can be used to make multiline strings line up with the left
// edge of the display, while still presenting them in the source code
// in indented form.
//
// Note that tabs and spaces are both treated as whitespace, but they
// are not equal: the lines "  hello" and "\\thello" are
// considered to have no common leading whitespace.
//
// Entirely blank lines are normalized to a newline character.
//
// Sources:
// 	- https://docs.python.org/3/library/textwrap.html#textwrap.dedent
//  - https://github.com/python/cpython/blob/3.9/Lib/textwrap.py#L414
// 	- https://github.com/python/cpython/blob/3.9/Lib/test/test_textwrap.py#L694
func Dedent(text string) string {
	// Look for the longest leading string of spaces and tabs common to all lines
	margin := ""
	debug := false

	for idx, row := range strings.Split(text, "\n") {
		leadSpaces := leadSpaceRe.FindAllString(row, 1)
		if len(leadSpaces) == 0 {
			if debug {
				log.Printf("skip row[%d]: %q (only lead spaces)", idx, row)
			}
			continue
		}

		if margin == "" {
			margin = leadSpaces[0]
		} else if strings.HasPrefix(row, margin) {
			// Current line more deeply indented than previous winner:
			// no change (previous winner is still on top).
			if debug {
				log.Printf("skip row[%d]: %q (marging=%q)", idx, row, margin)
			}
			continue
		} else if strings.HasPrefix(margin, row) {
			// Current line consistent with and no deeper than previous winner:
			// it's the new winner.
			margin = row
		} else {
			// Find the largest common whitespace between current line and previous
			// winner.
			for i, c := range zip(margin, row) {
				if c.s1 != c.s2 {
					margin = margin[:i]
					break
				}
			}
		}
	}

	// sanity check (testing/debugging only)
	if debug && margin != "" {
		for idx, row := range strings.Split(text, "\n") {
			if row == "" || strings.HasPrefix(row, margin) {
				log.Printf("%d: line = %q, margin = %q", idx, row, margin)
			}
		}
	}

	if margin != "" {
		deleteLeadRe := regexp.MustCompile("(?m)^" + margin)
		return deleteLeadRe.ReplaceAllLiteralString(
			blankRowsRe.ReplaceAllLiteralString(text, ""),
			"")
	}
	return text
}
