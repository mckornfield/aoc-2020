package aoc16

import (
	"regexp"
	"strconv"
	"strings"
)

type Validator interface {
	validate(int) bool
}

const exp = `(\d+)-(\d+) or (\d+)-(\d+)`

func BuildValidator(line string) Validator {
	rangeExp, err := regexp.Compile(exp)
	if err != nil {
		panic(err)
	}
	vals := rangeExp.FindStringSubmatch(line)
	bottomOne, _ := strconv.Atoi(vals[1])
	topOne, _ := strconv.Atoi(vals[2])
	bottomTwo, _ := strconv.Atoi(vals[3])
	topTwo, _ := strconv.Atoi(vals[4])

	return IntValidator{
		BottomOne: bottomOne,
		TopOne:    topOne,
		BottomTwo: bottomTwo,
		TopTwo:    topTwo,
	}

}

func IsNumberValid(validators []Validator, number int) bool {
	for _, validator := range validators {
		if validator.validate(number) {
			return true
		}
	}
	return false
}

func FindSumOfInvalidInputs(input string) int {
	buildingValidators := true
	rangeExp, err := regexp.Compile(exp)
	if err != nil {
		panic(err)
	}
	validators := []Validator{}
	invalidNums := []int{}
	invalidSum := 0
	for _, line := range strings.Split(input, "\n") {
		if buildingValidators && rangeExp.MatchString(line) {
			validators = append(validators, BuildValidator(line))
		} else if !buildingValidators {
			for _, numStr := range strings.Split(strings.Trim(line, "\t"), ",") {
				num, err := strconv.Atoi(numStr)
				if err != nil {
					panic(err)
				}
				if !IsNumberValid(validators, num) {
					invalidNums = append(invalidNums, num)
					invalidSum += num
				}
			}
		}
		if strings.Contains(line, "nearby tickets:") {
			buildingValidators = false
		}
	}
	return invalidSum
}

func (intv IntValidator) validate(val int) bool {
	return (val >= intv.BottomOne && val <= intv.TopOne) ||
		(val >= intv.BottomTwo && val <= intv.TopTwo)
}

type IntValidator struct {
	BottomOne int
	TopOne    int
	BottomTwo int
	TopTwo    int
}
