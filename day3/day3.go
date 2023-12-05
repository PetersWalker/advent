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

	// prev := []string{}
	// next := []string{}
	for scanner.Scan() {
		line := scanner.Text()

		for _, v := range line {
			s := string(v)

			if number, err := strconv.Atoi(s); err != nil {
				log.Println(number)

			}
		}

	}

}
