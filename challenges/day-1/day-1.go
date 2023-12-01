package day1

import (
	"strconv"
	"strings"
)

var numberStringMap = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func Day1(input string, part2 bool) (o int) {
	values := strings.Split(input, "\n")

	for _, v := range values {
		var first string = ""
		var last string = ""
		cp := v
		rv := v

	forward:
		for len(cp) > 0 {
			str := ""
			for _, e := range cp {
				str += string(e)
				if _, err := strconv.Atoi(str); err == nil {
					first = str
					break forward
				}
				if val, ok := numberStringMap[str]; ok {
					first = strconv.Itoa(val)
					break forward
				}
			}

			var builder strings.Builder
			builder.WriteString(cp[1:])
			cp = builder.String()
		}
	reverse:
		for len(rv) > 0 {
			str := ""
			for i := len(rv) - 1; i >= 0; i-- {
				str = string(rv[i]) + str
				if _, err := strconv.Atoi(str); err == nil {
					last = str
					break reverse
				}
				if val, ok := numberStringMap[str]; ok {
					last = strconv.Itoa(val)
					break reverse
				}
			}

			var builder strings.Builder
			builder.WriteString(rv[:len(rv)-1])
			rv = builder.String()
		}

		d, _ := strconv.Atoi(first + last)
		o += d
	}

	return
}
