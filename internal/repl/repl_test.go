package repl

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestCleanInput(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected []string
	}{
		"two words":   {input: "yallah hallah", expected: []string{"yallah", "hallah"}},
		"casing":      {input: "YaLlah hAlLah", expected: []string{"yallah", "hallah"}},
		"whitespaces": {input: "     yallah      hallah       ", expected: []string{"yallah", "hallah"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := cleanInput(tc.input)
			diff := cmp.Diff(got, tc.expected)
			if diff != "" {
				t.Fatalf("%v", diff)
			}
		})
	}
}
