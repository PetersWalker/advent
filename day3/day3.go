package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	partNumberSum, gearRatioSum := processFile("./input.txt")
	log.Println("gearRatioSum: ", gearRatioSum)
	log.Println("partNumberSum: ", partNumberSum)
}

func processFile(path string) (int, int) {
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var prev string
	var next string
	nums := []string{}
	partNumberSum := 0

	gearRatioSum := 0
	for i, line := range lines {

		// prev validation
		if i > 0 {
			prev = lines[i-1]
		}

		// next validation
		if i < len(lines)-1 {
			next = lines[i+1]
		} else {
			next = ""
		}

		coordinates := checkLineForParts(prev, line, next, i)
		partNumbers := coordinatesToNumbers(coordinates, lines)
		nums = append(nums, partNumbers...)

		coordinatePairs := checkLineForGearNumbers(prev, line, next, i)

		for _, coordinatePair := range coordinatePairs {
			numbers := coordinatesToNumbers(coordinatePair, lines)

			if len(numbers) != 2 {
				log.Println("gearNumbers", numbers)
				continue
			}

			// log.Println("gearNumbers", numbers)

			first, err := strconv.Atoi(numbers[0])
			if err != nil {
				log.Panic(err)
			}

			second, err := strconv.Atoi(numbers[1])
			if err != nil {
				log.Panic(err)
			}

			gearRatio := first * second
			gearRatioSum += gearRatio
		}
	}

	for _, num := range nums {
		number, err := strconv.Atoi(num)

		if err != nil {
			log.Fatal(err)
		}

		partNumberSum += number
	}

	log.Println("sum is ", partNumberSum)
	return partNumberSum, gearRatioSum
}

type Coordinate struct {
	Row    int
	Column int
}

func checkLineForGearNumbers(prev string, current string, next string, rowIndex int) [][]Coordinate {
	isANumber := map[string]bool{
		"1": true,
		"2": true,
		"3": true,
		"4": true,
		"5": true,
		"6": true,
		"7": true,
		"8": true,
		"9": true,
		"0": true,
	}

	coordinatePairs := [][]Coordinate{}
	for i, v := range current {
		coordinates := []Coordinate{}
		appendCoordinate := func(rowIndex, columnIndex int) {
			coordinates = append(coordinates, Coordinate{
				Row:    rowIndex,
				Column: columnIndex,
			})
		}

		prevExists := (prev != "")
		nextExists := (next != "")

		notAtStart := (i != 0)
		notAtEnd := (i < len(current)-1)

		//if its not a gear, we don't care
		if string(v) != "*" {
			continue
		}

		//top
		if prevExists && isANumber[string(prev[i])] {
			appendCoordinate(rowIndex-1, i)
		}

		//top left
		if prevExists && notAtStart && isANumber[string(prev[i-1])] {
			appendCoordinate(rowIndex-1, i-1)
		}

		//top right
		if prevExists && notAtEnd && isANumber[string(prev[i+1])] {
			appendCoordinate(rowIndex-1, i+1)
		}

		//left
		if notAtStart && isANumber[string(current[i-1])] {
			appendCoordinate(rowIndex, i-1)
		}
		//right
		if notAtEnd && isANumber[string(current[i+1])] {
			appendCoordinate(rowIndex, i+1)
		}

		// bottom
		if nextExists && isANumber[string(next[i])] {
			appendCoordinate(rowIndex+1, i)
		}

		// bottom left
		if nextExists && notAtStart && isANumber[string(next[i-1])] {
			appendCoordinate(rowIndex+1, i-1)
		}

		// bottom right
		if nextExists && notAtEnd && isANumber[string(next[i+1])] {
			appendCoordinate(rowIndex+1, i+1)
		}

		// flush coordinates
		coordinatePairs = append(coordinatePairs, coordinates)
	}
	return coordinatePairs
}

func checkLineForParts(prev string, current string, next string, lineIndex int) []Coordinate {
	isNotAPart := map[string]bool{
		"1": true,
		"2": true,
		"3": true,
		"4": true,
		"5": true,
		"6": true,
		"7": true,
		"8": true,
		"9": true,
		"0": true,
		".": true,
	}

	coordinates := []Coordinate{}
	for i, v := range current {

		appendCoordinate := func(columnIndex int) {
			coordinates = append(coordinates, Coordinate{
				Row:    lineIndex,
				Column: columnIndex,
			})
		}

		_, err := strconv.Atoi(string(v))

		if err != nil {
			continue
		}

		prevExists := (prev != "")
		nextExists := (next != "")

		notAtStart := (i != 0)
		notAtEnd := (i < len(current)-1)

		//top
		if prevExists && !isNotAPart[string(prev[i])] {
			appendCoordinate(i)
			continue
		}

		//top left
		if prevExists && notAtStart && !isNotAPart[string(prev[i-1])] {
			appendCoordinate(i)
			continue
		}
		//top right
		if prevExists && notAtEnd && !isNotAPart[string(prev[i+1])] {
			appendCoordinate(i)
			continue
		}

		//left
		if notAtStart && !isNotAPart[string(current[i-1])] {
			appendCoordinate(i)
			continue
		}
		//right
		if notAtEnd && !isNotAPart[string(current[i+1])] {
			appendCoordinate(i)
			continue
		}

		// bottom
		if nextExists && !isNotAPart[string(next[i])] {
			appendCoordinate(i)
			continue
		}

		// bottom left
		if nextExists && notAtStart && !isNotAPart[string(next[i-1])] {
			appendCoordinate(i)
			continue
		}

		// bottom right
		if nextExists && notAtEnd && !isNotAPart[string(next[i+1])] {
			appendCoordinate(i)
			continue
		}
	}
	return coordinates
}

func coordinatesToNumbers(coordinates []Coordinate, lines []string) []string {
	nums := []string{}
	numSet := map[int]string{}
	for _, coordinate := range coordinates {
		initialLeft := coordinate.Column
		initialRight := coordinate.Column
		line := lines[coordinate.Row]
		left, right := searchLeftAndRight(line, initialLeft, initialRight)
		numSet[left] = line[left : right+1]
	}
	for _, num := range numSet {
		nums = append(nums, num)
	}
	return nums
}

func searchLeftAndRight(line string, left int, right int) (int, int) {
	isANumber := map[string]bool{
		"1": true,
		"2": true,
		"3": true,
		"4": true,
		"5": true,
		"6": true,
		"7": true,
		"8": true,
		"9": true,
		"0": true,
	}

	var nextLeft = ""
	var nextRight = ""

	nextLeftExists := ((left - 1) >= 0)
	nextRightExists := ((right + 1) < len(line))

	if nextLeftExists {
		nextLeft = string(line[left-1])
	}

	if nextRightExists {
		nextRight = string(line[right+1])
	}

	if isANumber[nextLeft] {
		left = left - 1
	}

	if isANumber[nextRight] {
		right = right + 1
	}

	if !isANumber[nextLeft] && !isANumber[nextRight] {
		return left, right
	}

	return searchLeftAndRight(line, left, right)
}

// 1. when its on line one youre not aware of line two

// 1. set as prev .....................................164.................429.35...........221....................................................34.........

// 2. prev exists so you ca compare that

// prev .....................................164.................429.35...........221....................................................34.........
// line ........................464...........*.................................../.......53*.....954.763.....................114*.764..............

// 2. after compare, set line to prev, and prev to prevprev

// 3.

// prevprev .....................................164.................429.35...........221....................................................34.........
// prev ........................464...........*.................................../.......53*.....954.763.....................114*.764..............
// line 223............275.....................725.....$.........460....176............................*............+.................&.267.........

// .....................................164.................429.35...........221....................................................34.........
// ........................464...........*.................................../.......53*.....954.763.....................114*.764..............
// 223............275.....................725.....$.........460....176............................*............+.................&.267.........
// .........854..........919.798...............541.....302...................723......$...............196.......275......$....@....*...+2...388
// ..........@.......284*............429..211..........*..........633.............503..66......865.....*....234..........21....918.779..../....
