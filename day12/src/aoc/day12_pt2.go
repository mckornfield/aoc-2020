package aoc

import (
	"math"
	"strings"
)

func MoveAndFindDistanceAroundWaypoint(input string) int {
	location := Location{X: 0, Y: 0}
	direction := 0 // along the positive x axis, east
	waypointLocation := Location{X: 10, Y: 1}
	for _, line := range strings.Split(input, "\n") {
		actionAndMag := ParseInput(line)
		magnitude := actionAndMag.Magnitude
		originalDirection := direction
		switch actionAndMag.Action {
		case ("F"):
			location.moveTowardWaypoint(magnitude, waypointLocation)
			break
		case ("N"):
			waypointLocation.moveNorth(magnitude)
			break
		case ("S"):
			waypointLocation.moveSouth(magnitude)
			break
		case ("E"):
			waypointLocation.moveEast(magnitude)
			break
		case ("W"):
			waypointLocation.moveWest(magnitude)
			break
		case ("L"): // Clockwise is positive
			direction += magnitude
			break
		case ("R"):
			direction -= magnitude
			break
		}
		for direction >= 360 {
			direction -= 360
		}
		for direction < 0 {
			direction += 360
		}
		if originalDirection != direction {
			waypointLocation.rotateLocation(originalDirection, direction)
			// Do the rotate thing
		}
		// fmt.Println(actionAndMag)
		// fmt.Println(location)
		// fmt.Println(waypointLocation)
	}
	return int(math.Abs(float64(location.X))) + int(math.Abs(float64(location.Y)))

}

func (l *Location) moveTowardWaypoint(magnitude int, waypoint Location) {
	l.X += waypoint.X * magnitude
	l.Y += waypoint.Y * magnitude
}

func (l *Location) rotateLocation(originalDirection, currentDirection int) {
	rotation := originalDirection - currentDirection
	// Unit circle, (1,0) East, (0,-1) South
	// (0,1 ) North, (-1,0) West
	if rotation == 90 || rotation == -270 {
		l.X, l.Y = l.Y, -l.X
	} else if rotation == -90 || rotation == 270 {
		l.X, l.Y = -l.Y, l.X
	} else if math.Abs(float64(rotation)) == 180 {
		l.Y, l.X = -l.Y, -l.X
	}
}
