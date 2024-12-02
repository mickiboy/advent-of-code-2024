package main

import (
	_ "embed"

	"github.com/mickiboy/advent-of-code-2024/internal/day2"
)

//go:embed input
var list string

func main() {
	println("Part 1 (safe reports):", day2.Part1(list))
	println("Part 1 (safe reports with problem dampener):", day2.Part2(list))
}
