package aoc

import (
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
