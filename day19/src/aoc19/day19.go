package aoc19

import (
	"strings"
)

func GetRuleZeroMatches(input string) int {
	ruleMap := make(map[string]Rule)
	linesToCheck := []string{}
	for _, line := range strings.Split(input, "\n") {
		if strings.Contains(line, ":") {
			segments := strings.Split(line, ": ")
			if strings.Contains(line, "\"") {
				id := segments[0]
				literal := strings.Replace(segments[1], "\"", "", 2)
				ruleMap[id] = LiteralRule{literal: literal}
			} else {
				id := segments[0]
				rules := strings.Split(segments[1], " | ")
				orReferences := make([][]string, len(rules))
				for idx, ruleSet := range rules {
					orReferences[idx] = strings.Split(ruleSet, " ")
				}
				ruleMap[id] = ReferenceRule{orReferences: orReferences}
			}
		} else if len(line) > 0 {
			linesToCheck = append(linesToCheck, line)
		}
	}
	count := 0
	for _, line := range linesToCheck {
		ruleZero := ruleMap["0"]
		// fmt.Println(ruleZero)
		if ok, finalIndex := ruleZero.matches(ruleMap, line, 0); ok {
			if finalIndex != len(line) {
				// fmt.Println("Line " + line + " matches the rules but is not the right length, final index " + fmt.Sprint(finalIndex))
			} else {
				// fmt.Println("Line " + line + " matches the rules")
				count++
			}
		}
	}
	return count
}

type LiteralRule struct {
	literal string
}

type ReferenceRule struct {
	orReferences [][]string
}

func (r ReferenceRule) matches(ruleSet map[string]Rule, stringToMatch string, startingIndex int) (bool, int) {

	for _, references := range r.orReferences {
		matches := true
		newStartIndex := startingIndex
		for _, singleReference := range references {
			ruleToMatch := ruleSet[singleReference]
			// fmt.Println(fmt.Sprintf("Rule to match %v", ruleToMatch))
			// fmt.Println(fmt.Sprintf("Char to match %s", string(stringToMatch[newStartIndex])))
			matches, newStartIndex = ruleToMatch.matches(ruleSet, stringToMatch, newStartIndex)
			if !matches {
				break
			}
		}
		// fmt.Println(fmt.Sprintf("references: %v , matches %s", references, strconv.FormatBool(matches)))
		if matches {
			return true, newStartIndex
		}
	}
	// fmt.Println("Made it all the way here")
	return false, startingIndex
}

func (l LiteralRule) matches(ruleSet map[string]Rule, stringToMatch string, startingIndex int) (bool, int) {
	if len(stringToMatch)-1 < startingIndex {
		return false, startingIndex
	}
	return l.literal == string(stringToMatch[startingIndex]), startingIndex + 1
}

type Rule interface {
	// This returns whether or not it matches the string, and how far it advanced
	matches(map[string]Rule, string, int) (bool, int)
}
