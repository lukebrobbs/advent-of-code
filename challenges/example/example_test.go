package example

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected string
	}{
		"simple": {input: "123\n1234\n12345", expected: "2"},
		"longer": {input: "123\n1234\n12345\n1\n6353\n99\n45\n124124", expected: "4"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, Example(tc.input))
		})
	}
}
