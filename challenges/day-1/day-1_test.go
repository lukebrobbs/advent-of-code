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
		"part 1": {input: "1000\n2000\n3000\n\n4000\n\n5000\n6000\n\n7000\n8000\n9000\n\n10000", part2: false, expected: 24000},
		"part 2": {input: "1000\n2000\n3000\n\n4000\n\n5000\n6000\n\n7000\n8000\n9000\n\n10000", part2: true, expected: 45000},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, Day1(tc.input, tc.part2))
		})
	}
}
