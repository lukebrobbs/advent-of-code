package main

import (
	"flag"
	"fmt"

	"github.com/lukebrobbs/advent-of-code-2023/challenges"
	"github.com/lukebrobbs/advent-of-code-2023/utils"
)

var day = flag.Int("day", 1, "Advent of code day to run")
var part2 = flag.Bool("part-2", false, "Should we run part 2 of the code")

func main() {
	flag.Parse()
	challenges := challenges.DayChallenges
	bytes, err := utils.FileReader(*day)
	utils.Check(err)
	fmt.Println(challenges[*day](string(bytes), *part2))
}
