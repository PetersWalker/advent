package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// The Elf would first like to know which games would have been
// possible if the bag contained only 12 red cubes, 13 green cubes,
// and 14 blue cubes?

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Panic(err)
	}

	defer file.Close()

	scanner1 := bufio.NewScanner(file)

	scanner2 := bufio.NewScanner(os.Stdin)
	count := processLines(scanner1)
	log.Println("total count for possible games: ", count)

	minPowerCube := minPowerCube(scanner2)
	log.Println("minPowerCube: ", minPowerCube)

}

func minPowerCube(s *bufio.Scanner) int {
	sum := 0

	for s.Scan() {
		log.Println("powerCube :")
		line := s.Text()

		minRed := 0
		minGreen := 0
		minBlue := 0

		colorCounts := parseColorValues(line)

		for _, colorValue := range colorCounts {
			if colorValue["red"] > minRed {
				minRed = colorValue["red"]
			}

			if colorValue["green"] > minGreen {
				minGreen = colorValue["green"]
			}

			if colorValue["blue"] > minBlue {
				minBlue = colorValue["blue"]
			}
		}
		powerCube := minRed * minGreen * minBlue

		sum += powerCube
	}
	return sum
}

func processLines(s *bufio.Scanner) int {
	possibleGames := 0
	red := 12
	green := 13
	blue := 14

	for s.Scan() {
		possible := true
		line := s.Text()

		gameNumber := parseGameNumber(line)
		colorCounts := parseColorValues(line)

		for _, colorValues := range colorCounts {
			if colorValues["red"] > red {
				possible = false
			}

			if colorValues["green"] > green {
				possible = false
			}

			if colorValues["blue"] > blue {
				possible = false
			}
		}

		if possible {
			possibleGames += gameNumber
		}

	}

	return possibleGames
}

// "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
func parseColorValues(line string) []map[string]int {
	colorCounts := []map[string]int{}
	gameRemovedString := strings.Split(line, ":")[1]
	iterations := strings.Split(gameRemovedString, ";") // ["3 blue, 4 red" , "1 red, 2 green, 6 blue" , "2 green"]

	for _, v := range iterations {
		iteration := strings.Split(v, ",") // ["3 blue " , " 4 red"]

		colorCountMap := map[string]int{}
		for _, colorString := range iteration {
			trimmed := strings.Trim(colorString, " ")
			colorArray := strings.Split(trimmed, " ")
			count, err := strconv.Atoi(colorArray[0])
			color := colorArray[1]
			if err != nil {
				log.Panic(err)
			}
			colorCountMap[color] = count
		}
		colorCounts = append(colorCounts, colorCountMap)
	}
	return colorCounts
}

// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
func parseGameNumber(line string) int {
	gameString := strings.Split((strings.Split(line, ":")[0]), " ")[1]
	i, err := strconv.Atoi(string(gameString))

	if err != nil {
		log.Panic(err)
	}

	return i
}
