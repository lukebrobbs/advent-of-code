package day3

import (
	"regexp"
	"strconv"
)

func Day3(input string, part2 bool) int {
	total := 0
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := re.FindAllString(input, -1)

	for _, inst := range matches {
		dRe := regexp.MustCompile(`\d+`)
		nums := dRe.FindAllString(inst, -1)
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])
		total += (num1 * num2)
	}
	return total
}
