package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func day1() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		log.Println("*******************")
		first := 0
		last := 0
		s := scanner.Text()

		log.Println("original:", s)
		replaced := insertDigits(s)
		log.Println("replaced:", replaced)
		s = replaced

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

		log.Println("results: ", s, " ", first, last)
		sum = sum + (first * 10) + last

	}

	println(sum)
}

func insertDigits(s string) string {
	digits := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	for k, v := range digits {
		for strings.Index(s, k) != -1 {
			index := strings.Index(s, k)
			s = s[:index+1] + v + s[index+1:]
		}
	}

	return s
}
