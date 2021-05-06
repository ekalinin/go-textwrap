package textwrap

import (
	"strings"
	"testing"
)

// The roundtrip cases are separate, because textwrap.Dedent doesn't
// handle Windows line endings.
var roundtripCases = []string{
	// Basic test case
	"Hi.\nThis is a test.\nTesting.",
	// Include a blank line
	"Hi.\nThis is a test.\n\nTesting.",
	// Include leading and trailing blank lines
	"\nHi.\nThis is a test.\nTesting.\n",
}

var cases = []string{
	// Use Windows line endings
	"Hi.\r\nThis is a test.\r\nTesting.\r\n",
	// Pathological case
	"\nHi.\r\nThis is a test.\n\r\nTesting.\r\n\n",
}

func allCases() []string {
	return append(roundtripCases, cases...)
}

// TestIndentNomarginDefault checks that Indent should do nothing
// if 'prefix' is empty.
func TestIndentNomarginDefault(t *testing.T) {
	for idx, test := range allCases() {
		got := Indent(test, "", nil)
		if test != got {
			t.Errorf("test #%d:\n want: %q\n  got: %q", idx, test, got)
		}
	}
}

// TestIndentNomarginAllLines is the same check as above one, but using
// the optional predicate argument.
func TestIndentNomarginAllLines(t *testing.T) {
	for idx, test := range allCases() {
		got := Indent(test, "", Any)
		if test != got {
			t.Errorf("test #%d:\n want: %q\n  got: %q", idx, test, got)
		}
	}
}

func TestIndentNomarginNoLines(t *testing.T) {
	for idx, test := range allCases() {
		got := Indent(test, "", None)
		if test != got {
			t.Errorf("test #%d:\n want: %q\n  got: %q", idx, test, got)
		}
	}
}

func TestIndentRoundtripSpaces(t *testing.T) {
	for idx, test := range roundtripCases {
		got := Dedent(Indent(test, "    ", nil))
		if test != got {
			t.Errorf("test #%d:\n want: %q\n  got: %q", idx, test, got)
		}
	}
}

func TestIndentRoundtripTabs(t *testing.T) {
	for idx, test := range roundtripCases {
		got := Dedent(Indent(test, "\t\t", nil))
		if test != got {
			t.Errorf("test #%d:\n want: %q\n  got: %q", idx, test, got)
		}
	}
}

func TestIndentRoundtripMixed(t *testing.T) {
	for idx, test := range roundtripCases {
		got := Dedent(Indent(test, " \t  \t ", nil))
		if test != got {
			t.Errorf("test #%d:\n want: %q\n  got: %q", idx, test, got)
		}
	}
}

// TestIndentDefault checks that Indent adds 'prefix' to all lines, including
// whitespace-only ones.
func TestIndentAllLines(t *testing.T) {
	prefix := "  "
	expected := []string{
		// Basic test case
		"  Hi.\n  This is a test.\n  Testing.",
		// Include a blank line
		"  Hi.\n  This is a test.\n  \n  Testing.",
		// Include leading and trailing blank lines
		"  \n  Hi.\n  This is a test.\n  Testing.\n  ",
		// Use Windows line endings
		"  Hi.\r\n  This is a test.\r\n  Testing.\r\n  ",
		// Pathological case
		"  \n  Hi.\r\n  This is a test.\n  \r\n  Testing.\r\n  \n  ",
	}

	for idx, test := range allCases() {
		want := expected[idx]
		got := Indent(test, prefix, Any)
		if want != got {
			t.Errorf("test #%d:\n want: %q\n  got: %q", idx, want, got)
		}
	}
}

// TestIndentEmptyLines checks that Indent adds 'prefix' solely to
// whitespace-only lines.
func TestIndentEmptyLines(t *testing.T) {
	prefix := "  "
	expected := []string{
		// Basic test case
		"Hi.\nThis is a test.\nTesting.",
		// Include a blank line
		"Hi.\nThis is a test.\n  \nTesting.",
		// Include leading and trailing blank lines
		"  \nHi.\nThis is a test.\nTesting.\n  ",
		// Use Windows line endings
		"Hi.\r\nThis is a test.\r\nTesting.\r\n  ",
		// Pathological case
		"  \nHi.\r\nThis is a test.\n  \r\nTesting.\r\n  \n  ",
	}
	isEmpty := func(s string) bool {
		return strings.TrimSpace(s) == ""
	}

	for idx, test := range allCases() {
		want := expected[idx]
		got := Indent(test, prefix, isEmpty)
		if want != got {
			t.Errorf("test #%d:\n want: %q\n  got: %q", idx, want, got)
		}
	}
}
