package day2

import (
	"log"
	"math"
	"strconv"
	"strings"
)

func Part1(list string) int {
	return getNumSafeReports(list, false)
}

func Part2(list string) int {
	return getNumSafeReports(list, true)
}

func getNumSafeReports(list string, ignoreSingleError bool) int {
	numSafeReports := 0

	reports := strings.Split(list, "\n")

	for _, report := range reports {
		if !strings.Contains(report, " ") {
			continue
		}

		trimmedReport := strings.TrimSpace(report)
		levelsAsStrings := strings.Split(trimmedReport, " ")

		levels := []int{}

		for _, levelAsString := range levelsAsStrings {
			level, err := strconv.Atoi(levelAsString)
			if err != nil {
				log.Fatalf("failed to convert %s (%v)", levelAsString, err)
			}

			levels = append(levels, level)
		}

		if isSafetyGuaranteed(levels, ignoreSingleError) {
			numSafeReports++
		}
	}

	return numSafeReports
}

func isSafetyGuaranteed(levels []int, ignoreSingleError bool) bool {
	i := 0
	isIncreasing := true

	for i < len(levels)-1 {
		distance := levels[i] - levels[i+1]

		absDistance := int(math.Abs(float64(distance)))
		if absDistance < 1 || absDistance > 3 {
			if ignoreSingleError {
				return isSafetyGuaranteedWithOneLevelRemoved(levels, i)
			}

			return false
		}

		if i == 0 {
			isIncreasing = distance < 0
		} else {
			if isIncreasing && distance > 0 || !isIncreasing && distance < 0 {
				if ignoreSingleError {
					return isSafetyGuaranteedWithOneLevelRemoved(levels, i)
				}

				return false
			}
		}

		i++
	}

	return true
}

func isSafetyGuaranteedWithOneLevelRemoved(levels []int, i int) bool {
	if i > 0 {
		levelsWithoutLeftNumber := make([]int, len(levels))
		_ = copy(levelsWithoutLeftNumber, levels)
		levelsWithoutLeftNumber = append(levelsWithoutLeftNumber[:i-1], levelsWithoutLeftNumber[i:]...)

		if isSafetyGuaranteed(levelsWithoutLeftNumber, false) {
			return true
		}
	}

	levelsWithoutCurrentNumber := make([]int, len(levels))
	_ = copy(levelsWithoutCurrentNumber, levels)
	levelsWithoutCurrentNumber = append(levelsWithoutCurrentNumber[:i], levelsWithoutCurrentNumber[i+1:]...)

	if isSafetyGuaranteed(levelsWithoutCurrentNumber, false) {
		return true
	}

	if i <= len(levels)-2 {
		levelsWithoutRightNumber := make([]int, len(levels))
		_ = copy(levelsWithoutRightNumber, levels)
		levelsWithoutRightNumber = append(levelsWithoutRightNumber[:i+1], levelsWithoutRightNumber[i+2:]...)

		if isSafetyGuaranteed(levelsWithoutRightNumber, false) {
			return true
		}
	}

	return false
}
