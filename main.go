package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		first := 0
		last := 0
		s := scanner.Text()
		for i := 0; i < len(s); i++ {
			integer, err := strconv.Atoi(string(s[i]))

			if err != nil {
				continue
			}

			if first == 0 {
				first = integer
			}

			last = integer
		}

		// if last == 0 {
		// 	last = first
		// }

		log.Print(s, " ", first, last)
		sum = sum + (first * 10) + last
	}

	println(sum)
}
