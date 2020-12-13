package aoc

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func MoveAndFindDistance(input string) int {
	location := Location{X: 0, Y: 0}
	direction := 0 // along the positive x axis, east
	for _, line := range strings.Split(input, "\n") {
		actionAndMag := ParseInput(line)
		magnitude := actionAndMag.Magnitude
		switch actionAndMag.Action {
		case ("F"):
			location.moveForward(magnitude, direction)
			break
		case ("N"):
			location.moveNorth(magnitude)
			break
		case ("S"):
			location.moveSouth(magnitude)
			break
		case ("E"):
			location.moveEast(magnitude)
			break
		case ("W"):
			location.moveWest(magnitude)
			break
		case ("L"): // Clockwise is positive
			direction += magnitude
			break
		case ("R"):
			direction -= magnitude
			break
		}
		for direction > 360 {
			direction -= 360
		}
		for direction < 0 {
			direction += 360
		}
		fmt.Println(actionAndMag)
		fmt.Println(location)
	}
	return int(math.Abs(float64(location.X))) + int(math.Abs(float64(location.Y)))

}

func ParseInput(line string) ActionAndMagnitude {
	action := string(line[0])
	magnitude, err := strconv.Atoi(line[1:])
	if err != nil {
		panic(err)
	}
	return ActionAndMagnitude{
		Action:    action,
		Magnitude: magnitude,
	}
}

type Location struct {
	X, Y int
}

func (l *Location) moveForward(magnitude int, direction int) {
	if direction == 0 { // East
		l.moveEast(magnitude)
	} else if direction == 90 { // North
		l.moveNorth(magnitude)
	} else if direction == 180 { // West
		l.moveWest(magnitude)
	} else if direction == 270 { // South
		l.moveSouth(magnitude)
	} else {
		panic(fmt.Sprintf("Cannot move in given direction %d", direction))
	}

}

func (l *Location) moveEast(magnitude int) {
	l.X += magnitude
}

func (l *Location) moveNorth(magnitude int) {
	l.Y += magnitude
}
func (l *Location) moveWest(magnitude int) {
	l.X -= magnitude
}

func (l *Location) moveSouth(magnitude int) {
	l.Y -= magnitude
}

type ActionAndMagnitude struct {
	Action    string
	Magnitude int
}
