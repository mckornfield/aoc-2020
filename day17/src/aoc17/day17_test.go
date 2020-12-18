package aoc17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneSample(t *testing.T) {
	assert := assert.New(t)
	input := `.#.
..#
###`
	assert.Equal(5, EvolveCoordinates(input, 0))
	assert.Equal(11, EvolveCoordinates(input, 1))
	assert.Equal(21, EvolveCoordinates(input, 2))
	assert.Equal(112, EvolveCoordinates(input, 6))

}

func TestPartOne(t *testing.T) {
	assert := assert.New(t)
	input := `##......
.##...#.
.#######
..###.##
.#.###..
..#.####
##.####.
##..#.##`
	assert.Equal(306, EvolveCoordinates(input, 6))
}

func TestPartTwoSample(t *testing.T) {
	assert := assert.New(t)
	input := `.#.
..#
###`
	assert.Equal(5, EvolveCoordinates4d(input, 0))
	// assert.Equal(11, EvolveCoordinates(input, 1))
	// assert.Equal(21, EvolveCoordinates(input, 2))
	assert.Equal(848, EvolveCoordinates4d(input, 6))

}

func TestPartTwoAnswer(t *testing.T) {
	assert := assert.New(t)
	input := `##......
.##...#.
.#######
..###.##
.#.###..
..#.####
##.####.
##..#.##`
	assert.Equal(2572, EvolveCoordinates4d(input, 6))

}
