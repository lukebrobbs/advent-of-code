package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay3(t *testing.T) {
	tests := map[string]struct {
		input    string
		part2    bool
		expected int
	}{
		"part 1": {input: "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))", part2: false, expected: 161},
		"part 2": {input: "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))", part2: true, expected: 161},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, Day3(tc.input, tc.part2))
		})
	}
}
