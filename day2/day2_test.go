package main

import (
	"testing"

	"gotest.tools/assert"
)

var input string = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

func TestDay2(t *testing.T) {
	game1 := "Game 49: 11 red, 10 green, 12 blue; 4 green, 6 red, 6 blue; 9 red, 3 blue, 7 green; 6 red, 10 green, 14 blue; 8 blue, 10 red, 5 green; 6 blue, 17 green, 6 red"
	n1 := parseGameNumber(game1)

	assert.Equal(t, n1, 49)

	game2 := "Game 9: 11 red, 10 green, 12 blue; 4 green, 6 red, 6 blue; 9 red, 3 blue, 7 green; 6 red, 10 green, 14 blue; 8 blue, 10 red, 5 green; 6 blue, 17 green, 6 red"
	n2 := parseGameNumber(game2)

	assert.Equal(t, n2, 9)
}
