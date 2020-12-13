package aoc

import (
	"math"
	"strconv"
	"strings"
)

func FindNearestTimeStampAndBusIdProduct(input string) int {
	selectedBusId := 0
	overallMinWaitTime := math.MaxInt32 // Does golang have maxint?
	timestamp := 0
	busIDs := []int{}
	for _, line := range strings.Split(input, "\n") {
		// Parse lines
		if line == "" || line == "\n" {
			continue
		} else if strings.Contains(line, ",") {
			busIDs = ParseBusIds(line)
		} else {
			var err error
			timestamp, err = strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
		}
		if timestamp != 0 && len(busIDs) != 0 {
			for _, busID := range busIDs {
				minutesToWait := CalculateMinutesToWait(timestamp, busID)
				if overallMinWaitTime > minutesToWait {
					overallMinWaitTime = minutesToWait
					selectedBusId = busID
				}
			}
			timestamp = 0
			busIDs = []int{}
		}
	}

	return selectedBusId * overallMinWaitTime
}

func ParseBusIds(line string) []int {
	busIds := []int{}
	for _, id := range strings.Split(line, ",") {
		if id != "x" {
			parsedId, err := strconv.Atoi(id)
			if err != nil {
				panic(err)
			}
			busIds = append(busIds, parsedId)
		}
	}
	return busIds
}

func CalculateMinutesToWait(timestamp int, busId int) int {
	// Essentially this is just a divide, add to get the next number
	// If necessary, then subtract
	quotient := float64(timestamp) / float64(busId)
	if quotient == math.Floor(quotient) {
		// Zero wait time, crazy, not possible?
		return 0
	}
	// Otherwise ceil and subtract
	return int(math.Ceil(quotient))*int(busId) - timestamp
}
