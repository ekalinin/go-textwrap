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
			t.Errorf("[%d] want: %q, got: %q", idx, test, got)
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
