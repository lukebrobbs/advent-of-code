package example

import (
	"fmt"
	"strconv"
	"strings"
)

func Example(input string) string {
	values := strings.Split(input, "\n")
	var prev, total int
	for _, v := range values[1:] {
		st, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		if st > prev {
			total++
		}
		prev = st

	}
	return fmt.Sprint(total)
}
