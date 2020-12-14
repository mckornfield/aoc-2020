package aoc

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveAndFindManhattanDistanceSampleInput(t *testing.T) {
	assert := assert.New(t)
	input := `F10
N3
F7
R90
F11`
	distanceTravelled := MoveAndFindDistance(input)
	assert.Equal(25, distanceTravelled)
}

func TestPartTwo(t *testing.T) {
	assert := assert.New(t)
	input := ReadFile(t, "day12_puzzle.txt")
	distanceTravelled := MoveAndFindDistanceAroundWaypoint(input)
	assert.Equal(30761, distanceTravelled)
}

func TestPartTwoSample(t *testing.T) {
	assert := assert.New(t)
	input := `F10
N3
F7
R90
F11`
	distanceTravelled := MoveAndFindDistanceAroundWaypoint(input)
	assert.Equal(286, distanceTravelled)
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
