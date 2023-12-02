package day2

import (
	"strconv"
	"strings"
)

func getCubes() map[string]int {
	cubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	return cubes
}

func Day2(input string, part2 bool) int {
	values := strings.Split(input, "\n")
	cubes := getCubes()
	validGames := []int{}
gameLoop:
	for _, game := range values {
		gr := strings.Split(game, ":")
		gameNumber, _ := strconv.Atoi(strings.Split(gr[0], " ")[1])
		games := strings.Split(gr[1], ";")

		for _, round := range games {
			picks := strings.Split(round, ",")
			for _, pick := range picks {
				pick = strings.TrimSpace(pick)
				total := strings.Split(pick, " ")
				amount, _ := strconv.Atoi(total[0])
				color := total[1]
				if cubes[color] < amount {
					continue gameLoop
				}
			}
		}
		validGames = append(validGames, gameNumber)

	}
	var total int

	for _, v := range validGames {
		total += v
	}
	return total
}
