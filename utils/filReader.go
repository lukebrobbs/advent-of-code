package utils

import (
	"fmt"
	"os"
)

func FileReader(day int) ([]byte, error) {
	return os.ReadFile(fmt.Sprintf("inputs/day-%d.txt", day))
}
