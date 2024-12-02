package day2

import "testing"

const testList = `
7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`

func TestPart1(t *testing.T) {
	expected := 2

	actual := Part1(testList)

	if actual != expected {
		t.Fatalf("actual value %d is not equal to expected value %d", actual, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 4

	actual := Part2(testList)

	if actual != expected {
		t.Fatalf("actual value %d is not equal to expected value %d", actual, expected)
	}
}
