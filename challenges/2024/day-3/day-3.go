package day3

import (
	"regexp"
	"strconv"
)

const (
	DO   = `do()`
	DONT = `don't()`
)

func findViableMuls(items []string) []string {
	var tmp []string
	enabled := true

	for _, item := range items {
		if item == DONT {
			enabled = false
		} else if item == DO {
			enabled = true
		} else if enabled {
			tmp = append(tmp, item)
		}
	}

	return tmp
}

func Day3(input string, part2 bool) int {
	total := 0

	ptrn := `mul\(\d+,\d+\)`
	if part2 {
		ptrn += `|do\(\)|don\'t\(\)`
	}

	re := regexp.MustCompile(ptrn)
	matches := re.FindAllString(input, -1)

	if part2 {

		matches = findViableMuls(matches)
	}

	for _, inst := range matches {
		dRe := regexp.MustCompile(`\d+`)
		nums := dRe.FindAllString(inst, -1)
		if len(nums) == 2 {
			num1, _ := strconv.Atoi(nums[0])
			num2, _ := strconv.Atoi(nums[1])
			total += (num1 * num2)
		}
	}
	return total
}
