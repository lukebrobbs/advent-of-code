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
		"part 1": {input: "467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598..", part2: false, expected: 4361},
		"part 2": {input: "467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598..", part2: true, expected: 467835},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, Day3(tc.input, tc.part2))
		})
	}
}
