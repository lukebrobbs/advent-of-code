package utils

import (
	"fmt"
	"os"
)

func FileReader(year, day int) ([]byte, error) {
	return os.ReadFile(fmt.Sprintf("inputs/%d/day-%d.txt", year, day))
}
