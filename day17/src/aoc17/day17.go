package aoc17

import (
	"strings"
)

func EvolveCoordinates(input string, cycles int) int {
	actives := make(map[Coordinate]bool)
	for y, line := range strings.Split(input, "\n") {
		for x, c := range line {
			if c == '#' {
				coordinate := Coordinate{
					X: x,
					Y: y,
					Z: 0,
				}
				actives[coordinate] = true
			}
		}
	}
	for i := 0; i < cycles; i++ {
		neighborsAndCurrents := GetNeighborsAndCurrents(actives)
		newActives := make(map[Coordinate]bool)
		for coord := range neighborsAndCurrents {
			activeCubesNearby := GetNearbyActiveCubes(actives, coord)
			if _, ok := actives[coord]; ok {
				if activeCubesNearby != 2 && activeCubesNearby != 3 {
					// Deactivate
					delete(newActives, coord)
				} else {
					newActives[coord] = true
				}
			} else if activeCubesNearby == 3 {
				newActives[coord] = true
			}
		}
		actives = newActives
	}
	// fmt.Println(actives)
	return len(actives)
}

func EvolveCoordinates4d(input string, cycles int) int {
	actives := make(map[HCoordinate]bool)
	for y, line := range strings.Split(input, "\n") {
		for x, c := range line {
			if c == '#' {
				coordinate := HCoordinate{
					X: x,
					Y: y,
					Z: 0,
					W: 0,
				}
				actives[coordinate] = true
			}
		}
	}
	for i := 0; i < cycles; i++ {
		neighborsAndCurrents := GetNeighborsAndCurrents4d(actives)
		newActives := make(map[HCoordinate]bool)
		for coord := range neighborsAndCurrents {
			activeCubesNearby := GetNearbyActiveCubes4d(actives, coord)
			if _, ok := actives[coord]; ok {
				if activeCubesNearby != 2 && activeCubesNearby != 3 {
					// Deactivate
					delete(newActives, coord)
				} else {
					newActives[coord] = true
				}
			} else if activeCubesNearby == 3 {
				newActives[coord] = true
			}
		}
		actives = newActives
	}
	// fmt.Println(actives)
	return len(actives)
}

func GetNeighborsAndCurrents(actives map[Coordinate]bool) map[Coordinate]bool {
	neighborsAndCurrents := make(map[Coordinate]bool)
	for coordinate := range actives {
		for x0 := -1; x0 < 2; x0++ {
			for y0 := -1; y0 < 2; y0++ {
				for z0 := -1; z0 < 2; z0++ {
					neighborOrCurrent := coordinate.displace(Coordinate{
						X: x0,
						Y: y0,
						Z: z0,
					})
					if _, ok := neighborsAndCurrents[neighborOrCurrent]; !ok {
						neighborsAndCurrents[neighborOrCurrent] = true
					}
				}
			}
		}
	}
	return neighborsAndCurrents
}

func GetNeighborsAndCurrents4d(actives map[HCoordinate]bool) map[HCoordinate]bool {
	neighborsAndCurrents := make(map[HCoordinate]bool)
	for coordinate := range actives {
		for x0 := -1; x0 < 2; x0++ {
			for y0 := -1; y0 < 2; y0++ {
				for z0 := -1; z0 < 2; z0++ {
					for w0 := -1; w0 < 2; w0++ {
						neighborOrCurrent := coordinate.displace(HCoordinate{
							X: x0,
							Y: y0,
							Z: z0,
							W: w0,
						})
						if _, ok := neighborsAndCurrents[neighborOrCurrent]; !ok {
							neighborsAndCurrents[neighborOrCurrent] = true
						}
					}
				}
			}
		}
	}
	return neighborsAndCurrents
}

func GetNearbyActiveCubes(activeCubes map[Coordinate]bool, currentCoords Coordinate) int {
	active := 0
	counter := 0
	for x0 := -1; x0 < 2; x0++ {
		for y0 := -1; y0 < 2; y0++ {
			for z0 := -1; z0 < 2; z0++ {
				if x0 == 0 && y0 == 0 && z0 == 0 {
					continue
				}
				counter++
				displacement := Coordinate{
					X: x0,
					Y: y0,
					Z: z0,
				}
				if _, ok := activeCubes[currentCoords.displace(displacement)]; ok {
					active++
				}
			}
		}
	}
	return active
}

func GetNearbyActiveCubes4d(activeCubes map[HCoordinate]bool, currentCoords HCoordinate) int {
	active := 0
	counter := 0
	for x0 := -1; x0 < 2; x0++ {
		for y0 := -1; y0 < 2; y0++ {
			for z0 := -1; z0 < 2; z0++ {
				for w0 := -1; w0 < 2; w0++ {
					if x0 == 0 && y0 == 0 && z0 == 0 && w0 == 0 {
						continue
					}
					counter++
					displacement := HCoordinate{
						X: x0,
						Y: y0,
						Z: z0,
						W: w0,
					}
					if _, ok := activeCubes[currentCoords.displace(displacement)]; ok {
						active++
					}
				}
			}
		}
	}
	return active
}

type Coordinate struct {
	X int
	Y int
	Z int
}

type HCoordinate struct {
	X int
	Y int
	Z int
	W int
}

func (c Coordinate) displace(displacement Coordinate) Coordinate {
	return Coordinate{
		X: c.X + displacement.X,
		Y: c.Y + displacement.Y,
		Z: c.Z + displacement.Z,
	}

}

func (c HCoordinate) displace(displacement HCoordinate) HCoordinate {
	return HCoordinate{
		X: c.X + displacement.X,
		Y: c.Y + displacement.Y,
		Z: c.Z + displacement.Z,
		W: c.W + displacement.W,
	}

}
