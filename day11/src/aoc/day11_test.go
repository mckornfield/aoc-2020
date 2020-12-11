package aoc

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstExample(t *testing.T) {
	assert := assert.New(t)
	input := `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`
	grid := ParseInput(input)
	initialGrid := grid
	assert.Equal(grid.equal(grid), true)
	assert.Equal("L", grid.get(0, 0))
	assert.Equal("L", grid.get(0, 9))
	assert.Equal("L", grid.get(9, 9))
	assert.Equal(".", grid.get(7, 9))

	grid = IterateSeats(grid)
	assert.Equal(initialGrid.equal(grid), false)
	assert.Equal("#", grid.get(0, 0))
	assert.Equal("#", grid.get(0, 9))
	assert.Equal("#", grid.get(9, 9))
	assert.Equal(".", grid.get(7, 9))

	grid = IterateSeats(grid)
	assert.Equal("#", grid.get(0, 0))
	assert.Equal("#", grid.get(0, 9))
	assert.Equal("#", grid.get(9, 9))
	assert.Equal("L", grid.get(3, 9))
	assert.Equal(".", grid.get(7, 9))
}

func TestFirstExampleGetIterations(t *testing.T) {
	assert := assert.New(t)
	input := `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`
	assert.Equal(37, IterateUntilRepeat(input))
}

func TestFirstPart(t *testing.T) {
	assert := assert.New(t)
	content, err := ioutil.ReadFile("day11_puzzle.txt")
	if err != nil {
		t.Fatal(err.Error())
	}
	input := string(content)
	input = strings.TrimRight(input, "\n")
	assert.Equal(2194, IterateUntilRepeat(input))
	// fmt.Println(text)
}
