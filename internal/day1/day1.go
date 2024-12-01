package day1

import (
	"log"
	"math"
	"slices"
	"strconv"
	"strings"
)

func Part1(list string) int {
	leftList, rightList := getLists(list)

	slices.Sort(leftList)
	slices.Sort(rightList)

	totalDistance := 0

	for i := range leftList {
		totalDistance += int(math.Abs(float64(leftList[i] - rightList[i])))
	}

	return totalDistance
}

func Part2(list string) int {
	leftList, rightList := getLists(list)

	similarityScore := 0

	for _, leftListNumber := range leftList {
		numContaining := 0

		for _, rightListNumber := range rightList {
			if leftListNumber == rightListNumber {
				numContaining++
			}
		}

		similarityScore += leftListNumber * numContaining
	}

	return similarityScore
}

func getLists(list string) ([]int, []int) {
	lines := strings.Split(list, "\n")

	leftList := []int{}
	rightList := []int{}

	for _, line := range lines {
		if !strings.Contains(line, "   ") {
			continue
		}

		trimmedLine := strings.TrimSpace(line)
		numbers := strings.Split(trimmedLine, "   ")

		leftNumber, err := strconv.Atoi(numbers[0])
		if err != nil {
			log.Fatalf("failed to convert %s (%v)", numbers[0], err)
		}

		rightNumber, err := strconv.Atoi(numbers[1])
		if err != nil {
			log.Fatalf("failed to convert %s (%v)", numbers[1], err)
		}

		leftList = append(leftList, leftNumber)
		rightList = append(rightList, rightNumber)
	}

	return leftList, rightList
}
