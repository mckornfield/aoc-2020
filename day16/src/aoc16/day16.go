package aoc16

import (
	"regexp"
	"strconv"
	"strings"
)

type Validator interface {
	validate(int) bool
	validatorName() string
}

const exp = `(.*): (\d+)-(\d+) or (\d+)-(\d+)`

func BuildValidator(line string) Validator {
	rangeExp, err := regexp.Compile(exp)
	if err != nil {
		panic(err)
	}
	vals := rangeExp.FindStringSubmatch(line)
	name := vals[1]
	bottomOne, _ := strconv.Atoi(vals[2])
	topOne, _ := strconv.Atoi(vals[3])
	bottomTwo, _ := strconv.Atoi(vals[4])
	topTwo, _ := strconv.Atoi(vals[5])

	return IntValidator{
		Name:      name,
		BottomOne: bottomOne,
		TopOne:    topOne,
		BottomTwo: bottomTwo,
		TopTwo:    topTwo,
	}

}

func (intv IntValidator) validate(val int) bool {
	return (val >= intv.BottomOne && val <= intv.TopOne) ||
		(val >= intv.BottomTwo && val <= intv.TopTwo)
}

func (intv IntValidator) validatorName() string {
	return intv.Name
}

type IntValidator struct {
	Name      string
	BottomOne int
	TopOne    int
	BottomTwo int
	TopTwo    int
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

func LineToInts(line string) []int {
	vals := []int{}
	for _, numStr := range strings.Split(strings.Trim(line, "\t"), ",") {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}
		vals = append(vals, num)
	}
	return vals
}
func FilterRowsByValidInput(input string) map[string][]int {
	buildingValidators := true
	rangeExp, err := regexp.Compile(exp)
	if err != nil {
		panic(err)
	}
	validators := []Validator{}
	validTickets := [][]int{}
	for _, line := range strings.Split(input, "\n") {
		if buildingValidators && rangeExp.MatchString(line) {
			validators = append(validators, BuildValidator(line))
		} else if !buildingValidators {
			invalidLine := false
			lineAsInts := LineToInts(line)
			for _, num := range lineAsInts {
				if !IsNumberValid(validators, num) {
					invalidLine = true
					break
				}
			}
			if !invalidLine {
				validTickets = append(validTickets, lineAsInts)
			}
		}
		if strings.Contains(line, "nearby tickets:") {
			buildingValidators = false
		}
	}

	columnToValidatorName := make(map[string][]int)
	for _, validator := range validators {
		// Super slow, go over all the int columns
		for column := 0; column < len(validTickets[0]); column++ {
			validForColumn := true
			for row := 0; row < len(validTickets); row++ {
				if !validator.validate(validTickets[row][column]) {
					// fmt.Println("Validator " + validator.validatorName() + " not valid for column " + fmt.Sprint(column))
					validForColumn = false
					break
				}
			}
			if validForColumn {
				name := validator.validatorName()
				if currentVal, ok := columnToValidatorName[name]; ok {
					columnToValidatorName[name] = append(currentVal, column)
				} else {
					columnToValidatorName[name] = []int{column}
				}
			}
		}
	}

	return columnToValidatorName
}
