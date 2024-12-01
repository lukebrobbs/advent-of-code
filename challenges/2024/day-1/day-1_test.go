package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1(t *testing.T) {
	tests := map[string]struct {
		input    string
		part2    bool
		expected int
	}{
		"part 1": {input: "3   4\n4   3\n2   5\n1   3\n3   9\n3   3", part2: false, expected: 11},
		"part 2": {input: "3   4\n4   3\n2   5\n1   3\n3   9\n3   3", part2: true, expected: 11},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, Day1(tc.input, tc.part2))
		})
	}
}
