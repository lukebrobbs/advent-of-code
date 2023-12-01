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
		"part 1": {input: "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet", part2: false, expected: 142},
		"part 2": {input: "two1nine\neightwothree\nabcone2threexyz\n\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen\n", part2: true, expected: 281},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, Day1(tc.input, tc.part2))
		})
	}
}
