package main

import (
	"log"
	"testing"

	"gotest.tools/assert"
)

func TestCheckLine(t *testing.T) {

	// 467..114..
	// ...*......
	// ..35..633.
	// ......#...
	// 617*......
	// .....+.58.
	// ..592.....
	// ......755.
	// ...$.*....
	// .664.598..

	prev := "...*......"
	current := "..35..633."
	next := "......#..."

	result := checkLineForParts(prev, current, next, 1)
	desired := []Coordinate{
		{Row: 1, Column: 2},
		{Row: 1, Column: 3},
		{Row: 1, Column: 6},
		{Row: 1, Column: 7},
	}

	assert.DeepEqual(t, result, desired)
}

func TestCoordinatesToNumbers(t *testing.T) {
	coordinates1 := []Coordinate{
		{Row: 1, Column: 2},
		{Row: 1, Column: 3},
		{Row: 1, Column: 6},
		{Row: 1, Column: 7},
	}

	desired1 := []string{
		"35",
		"633",
	}

	result1 := coordinatesToNumbers(coordinates1, "..35..633.")

	assert.DeepEqual(t, desired1, result1)

	coordinates2 := []Coordinate{
		{Row: 1, Column: 0},
		{Row: 1, Column: 1},
		{Row: 1, Column: 6},
	}

	desired2 := []string{
		"35",
		"633",
	}

	result2 := coordinatesToNumbers(coordinates2, "35..633")

	assert.DeepEqual(t, desired2, result2)
}

func TestCheckLineForGearNumbers(t *testing.T) {

	// 467..114..
	// ...*......
	// ..35..633.
	prev := "467..114.."
	current := "...*......"
	next := "..35..633."
	result := checkLineForGearNumbers(prev, current, next, 1)

	log.Println(result)

	desired := [][]Coordinate{
		{
			{Row: 0, Column: 2}, {Row: 2, Column: 3}, {Row: 2, Column: 2},
		},
	}

	assert.DeepEqual(t, desired, result)
}
