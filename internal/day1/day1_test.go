package day1

import "testing"

const testList = `
3   4
4   3
2   5
1   3
3   9
3   3
`

func TestPart1(t *testing.T) {
	expected := 11

	actual := Part1(testList)

	if actual != expected {
		t.Fatalf("actual value %d is not equal to expected value %d", actual, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 31

	actual := Part2(testList)

	if actual != expected {
		t.Fatalf("actual value %d is not equal to expected value %d", actual, expected)
	}
}
