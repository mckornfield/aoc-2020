package aoc15

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountingGame(t *testing.T) {
	assert := assert.New(t)
	numbers := "0,3,6"
	spokenNumber := FindSpokenNumberForTurn(numbers, 2)
	assert.Equal(3, spokenNumber)
	spokenNumber = FindSpokenNumberForTurn(numbers, 3)
	assert.Equal(6, spokenNumber)
	spokenNumber = FindSpokenNumberForTurn(numbers, 4)
	assert.Equal(0, spokenNumber)
	spokenNumber = FindSpokenNumberForTurn(numbers, 5)
	assert.Equal(3, spokenNumber)
	spokenNumber = FindSpokenNumberForTurn(numbers, 6)
	assert.Equal(3, spokenNumber)
	spokenNumber = FindSpokenNumberForTurn(numbers, 7)
	assert.Equal(1, spokenNumber)
	spokenNumber = FindSpokenNumberForTurn(numbers, 9)
	assert.Equal(4, spokenNumber)
	spokenNumber = FindSpokenNumberForTurn(numbers, 10)
	assert.Equal(0, spokenNumber)
}

func TestCountingGameSecondSample(t *testing.T) {
	assert := assert.New(t)
	numbers := "1,3,2"
	spokenNumber := FindSpokenNumberForTurn(numbers, 2020)
	assert.Equal(1, spokenNumber)
}

func TestCountGamesPartOne(t *testing.T) {
	assert := assert.New(t)
	numbers := "6,13,1,15,2,0"
	spokenNumber := FindSpokenNumberForTurn(numbers, 2020)
	assert.Equal(1194, spokenNumber)
}

func TestCountGamesPartTwoSample(t *testing.T) {
	assert := assert.New(t)
	numbers := "0,3,6"
	spokenNumber := FindSpokenNumberForTurn(numbers, 30000000)
	assert.Equal(175594, spokenNumber)
}

func TestCountGamesPartTwo(t *testing.T) {
	assert := assert.New(t)
	numbers := "6,13,1,15,2,0"
	spokenNumber := FindSpokenNumberForTurn(numbers, 30000000)
	assert.Equal(48710, spokenNumber)
}

func ReadFile(t *testing.T, fileName string) string {

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		t.Fatal(err.Error())
	}
	input := string(content)
	input = strings.TrimRight(input, "\n")
	return input
}
