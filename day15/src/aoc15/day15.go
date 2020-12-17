package aoc15

import (
	"strconv"
	"strings"
)

func UpdateTimesSpokenMap(spokenToTurnMap map[int][]int, currentVal int, currentTurn int) {
	if val, ok := spokenToTurnMap[currentVal]; ok {
		spokenToTurnMap[currentVal] = []int{val[len(val)-1], currentTurn}
	} else {
		spokenToTurnMap[currentVal] = []int{currentTurn}
	}
}

func FindSpokenNumberForTurn(numbers string, turn int) int {
	splitNumbers := strings.Split(numbers, ",")
	startingNumbersInts := make([]int, len(splitNumbers))
	for idx, numStr := range strings.Split(numbers, ",") {
		num, _ := strconv.Atoi(numStr)
		startingNumbersInts[idx] = num
	}
	currentVal := 0
	previousSpokenToTurnMap := make(map[int][]int)
	for currentTurn := 1; currentTurn < turn+1; currentTurn++ {
		index := currentTurn - 1
		if len(startingNumbersInts) > index {
			currentVal = startingNumbersInts[index]
			// ENDGAME boys
		} else if twoPreviousTurns, ok := previousSpokenToTurnMap[currentVal]; ok && len(twoPreviousTurns) > 1 {
			// Some fancy turn comparison s**t
			currentVal = twoPreviousTurns[1] - twoPreviousTurns[0]
		} else { // Nah this is BRAND NEW
			currentVal = 0
		}
		UpdateTimesSpokenMap(previousSpokenToTurnMap, currentVal, currentTurn)
	}
	return currentVal

}
