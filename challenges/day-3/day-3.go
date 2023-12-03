package day3

import (
	"fmt"
	"image"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func Day3(input string, part2 bool) int {

	grid := map[image.Point]rune{}
	for y, s := range strings.Fields(string(input)) {
		fmt.Println(y, s)
		for x, r := range s {
			if r != '.' && !unicode.IsDigit(r) {
				grid[image.Point{x, y}] = r
			}
		}
	}

	parts := map[image.Point][]int{}
	for y, s := range strings.Fields(string(input)) {
		for _, m := range regexp.MustCompile(`\d+`).FindAllStringIndex(s, -1) {
			bounds := map[image.Point]struct{}{}
			for x := m[0]; x < m[1]; x++ {
				for _, d := range []image.Point{
					{-1, -1}, {-1, 0}, {-1, 1}, {0, -1},
					{0, 1}, {1, -1}, {1, 0}, {1, 1},
				} {
					bounds[image.Point{x, y}.Add(d)] = struct{}{}
				}
			}

			n, _ := strconv.Atoi(s[m[0]:m[1]])
			for p := range bounds {
				if _, ok := grid[p]; ok {
					parts[p] = append(parts[p], n)
				}
			}
		}
	}

	part1, part2Ret := 0, 0
	for p, ns := range parts {
		prod := 1
		for _, n := range ns {
			part1 += n
			prod *= n
		}
		if grid[p] == '*' && len(ns) == 2 {
			part2Ret += prod
		}
	}
	if part2 {
		return part2Ret
	}
	return part1

}
