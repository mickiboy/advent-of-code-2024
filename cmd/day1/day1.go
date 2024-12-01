package main

import (
	_ "embed"

	"github.com/mickiboy/advent-of-code-2024/internal/day1"
)

//go:embed input
var list string

func main() {
	println("Part 1 (total distance):", day1.Part1(list))
	println("Part 2 (similarity score):", day1.Part2(list))
}
