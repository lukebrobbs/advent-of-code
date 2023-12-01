package utils

import (
	"fmt"
	"io/ioutil"
)

func FileReader(day int) ([]byte, error) {
	return ioutil.ReadFile(fmt.Sprintf("inputs/day-%d.txt", day))
}
