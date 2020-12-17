package aoc

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSampleInputRoundTrip(t *testing.T) {
	assert := assert.New(t)
	input := `939
7,13,x,x,59,x,31,19
955
7,13,x,x,59,x,31,19`
	answer := FindNearestTimeStampAndBusIdProduct(input)
	assert.Equal(28, answer)
}

func TestParseBusIds(t *testing.T) {
	assert := assert.New(t)
	input := `7,13,x,x,59,x,31,19`
	busIds := ParseBusIds(input)
	assert.Equal([]int{7, 13, 59, 31, 19}, busIds)
}

func TestCalculateMinutesToWait(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(5, CalculateMinutesToWait(939, 59))
}

func TestPartOne(t *testing.T) {
	assert := assert.New(t)
	input := ReadFile(t, "day13_puzzle.txt")
	answer := FindNearestTimeStampAndBusIdProduct(input)
	assert.Equal(28, answer)
}

func TestPartTwoSample(t *testing.T) {
	assert := assert.New(t)
	input := "7,13,x,x,59,x,31,19"
	answer := FindFirstTimestampAllBusesWork(input)
	assert.Equal(28, answer)
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
