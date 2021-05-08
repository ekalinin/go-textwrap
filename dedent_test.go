package textwrap

import (
	"reflect"
	"testing"
)

func TestDedentUnchanged(t *testing.T) {
	tests := []string{
		// No lines indented.
		"Hello there.\nHow are you?\nOh good, I'm glad.",
		// Similar, with a blank line.
		"Hello there.\n\nBoo!",
		// Some lines indented, but overall margin is still zero.
		"Hello there.\n  This is indented.",
		// Again, add a blank line.
		"Hello there.\n\n  Boo!\n",
	}

	for idx, test := range tests {
		got := Dedent(test)
		if test != got {
			t.Errorf("test #%d:\n want: %q\n  got: %q", idx, test, got)
		}
	}
}

func TestZip(t *testing.T) {
	expected := []zipped{
		{'1', 'a'},
		{'2', 'b'},
	}

	got := zip("1234", "ab")
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("want: %q, got: %q", expected, got)
	}
}

func TestDedentEven(t *testing.T) {
	tests := []struct {
		input    string
		expexted string
	}{
		// All lines indented by two spaces.
		{
			input:    "  Hello there.\n  How are ya?\n  Oh good.",
			expexted: "Hello there.\nHow are ya?\nOh good.",
		},
		// Same, with blank lines.
		{
			input:    "  Hello there.\n\n  How are ya?\n  Oh good.\n",
			expexted: "Hello there.\n\nHow are ya?\nOh good.\n",
		},
		// Now indent one of the blank lines.
		{
			input:    "  Hello there.\n  \n  How are ya?\n  Oh good.\n",
			expexted: "Hello there.\n\nHow are ya?\nOh good.\n",
		},
	}

	for idx, test := range tests {
		got := Dedent(test.input)
		if test.expexted != got {
			t.Errorf("[%d]\n want: %q\n  got: %q", idx, test.expexted, got)
		}
	}
}

func TestDedentUneven(t *testing.T) {
	tests := []struct {
		input    string
		expexted string
	}{
		// Lines indented unevenly.
		{
			`
        def foo():
            while 1:
                return foo`,
			`
def foo():
    while 1:
        return foo`,
		},
		// Uneven indentation with a blank line.
		{
			"  Foo\n    Bar\n\n   Baz\n",
			"Foo\n  Bar\n\n Baz\n",
		},
		// Uneven indentation with a whitespace-only line.
		{
			"  Foo\n    Bar\n \n   Baz\n",
			"Foo\n  Bar\n\n Baz\n",
		},
	}

	for idx, test := range tests {
		got := Dedent(test.input)
		if test.expexted != got {
			t.Errorf("[%d]\n want: %q\n  got: %q", idx, test.expexted, got)
		}
	}
}

func TestDedentDeclining(t *testing.T) {
	tests := []struct {
		input    string
		expexted string
	}{
		// Uneven indentation with declining indent level.
		{
			"     Foo\n    Bar\n", // 5 spaces, then 4
			" Foo\nBar\n",
		},
		// Declining indent level with blank line.
		{
			"     Foo\n\n    Bar\n", // 5 spaces, blank, then 4
			" Foo\n\nBar\n",
		},
		// Declining indent level with whitespace only line.
		{
			"     Foo\n    \n    Bar\n", //5 spaces, then 4, then 4
			" Foo\n\nBar\n",
		},
	}

	for idx, test := range tests {
		got := Dedent(test.input)
		if test.expexted != got {
			t.Errorf("[%d]\n want: %q\n  got: %q", idx, test.expexted, got)
		}
	}
}

// TestDedentPreserveIntTabs checks that dedent() should not mangle internal tabs
func TestDedentPreserveIntTabs(t *testing.T) {
	tests := []struct {
		input    string
		expexted string
	}{
		// should not mangle internal tabs
		{
			"  hello\tthere\n  how are\tyou?",
			"hello\tthere\nhow are\tyou?",
		},
	}

	for idx, test := range tests {
		got := Dedent(test.input)
		if test.expexted != got {
			t.Errorf("[%d]\n want: %q\n  got: %q", idx, test.expexted, got)
		}
	}
}

// TestDedentPreserveMarginTabs checks that dedent() should not mangle tabs in
// the margin (i.e. tabs and spaces both count as margin, but are *not*
// considered equivalent)
func TestDedentPreserveMarginTabs(t *testing.T) {
	tests := []struct {
		input    string
		expexted string
	}{
		// unchanged
		{
			"  hello there\n\thow are you?",
			"  hello there\n\thow are you?",
		},
		// same effect even if we have 8 spaces
		{
			"        hello there\n\thow are you?",
			"        hello there\n\thow are you?",
		},
		// dedent() only removes whitespace that can be uniformly removed!
		{
			"\thello there\n\thow are you?",
			"hello there\nhow are you?",
		},
		{
			"  \thello there\n  \thow are you?",
			"hello there\nhow are you?",
		},
		{
			"  \t  hello there\n  \t  how are you?",
			"hello there\nhow are you?",
		},
		{
			"  \thello there\n  \t  how are you?",
			"hello there\n  how are you?",
		},
		// test margin is smaller than smallest indent
		{
			"  \thello there\n   \thow are you?\n \tI'm fine, thanks",
			" \thello there\n  \thow are you?\n\tI'm fine, thanks",
		},
	}

	for idx, test := range tests {
		got := Dedent(test.input)
		if test.expexted != got {
			t.Errorf("[%d]\n want: %q\n  got: %q", idx, test.expexted, got)
		}
	}
}
