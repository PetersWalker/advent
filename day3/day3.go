package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
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
	sum := 0

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

		coordinates := checkLine(prev, line, next, i)
		partNumbers := coordinatesToNumbers(coordinates, line)
		nums = append(nums, partNumbers...)
	}

	for _, num := range nums {
		number, err := strconv.Atoi(num)

		if err != nil {
			log.Fatal(err)
		}

		sum += number
	}

	log.Println("sum is ", sum)
}

type Coordinate struct {
	Row    int
	Column int
}

func checkLine(prev string, current string, next string, lineIndex int) []Coordinate {
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

func coordinatesToNumbers(coordinates []Coordinate, line string) []string {
	nums := []string{}
	numSet := map[int]string{}
	for _, coordinate := range coordinates {
		initialLeft := coordinate.Column
		initialRight := coordinate.Column
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
