package day1

import (
	"testing"

	"gotest.tools/assert"
)

func TestReplaceWords(t *testing.T) {
	result := insertDigits("eightwothree")
	assert.Equal(t, result, "e8ight2wot3hree")

	result2 := insertDigits("two1nine")
	assert.Equal(t, result2, "t2wo1n9ine")

	result3 := insertDigits("xtwone3four")
	assert.Equal(t, result3, "xt2wo1ne3f4our")

	result5 := insertDigits("eighthree")
	assert.Equal(t, result5, "e8ight3hree")
}
