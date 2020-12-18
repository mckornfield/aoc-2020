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

	// Find validators that satisfy columns
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

func contains(val int, values []int) (int, bool) {
	for idx, valToMatch := range values {
		if valToMatch == val {
			return idx, true
		}
	}
	return -1, false
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func ValidatorToColumnMapping(validatorToColumns map[string][]int) map[string]int {
	validatorToColumnsCopy := make(map[string][]int)
	for k, elems := range validatorToColumns {
		newElems := make([]int, len(elems))
		copy(newElems, elems)
		validatorToColumnsCopy[k] = newElems
	}
	validatorToColumnSingular := make(map[string]int)
	// Crappy perf, but doesn't take too long
	originalLen := len(validatorToColumns)
	for len(validatorToColumnSingular) != originalLen {
		// Find column with fewest matches
		lowestColumnName := ""
		lowestColumnCount := 4000
		columnVal := -1
		for validatorName, columns := range validatorToColumnsCopy {
			columnCount := len(columns)
			if columnCount < lowestColumnCount && columnCount > 0 {
				lowestColumnName = validatorName
				columnVal = columns[0]
				lowestColumnCount = columnCount
			}
		}
		// Now reduce
		for validatorName, columns := range validatorToColumnsCopy {
			if index, ok := contains(columnVal, columns); ok {
				validatorToColumnsCopy[validatorName] = remove(columns, index)
			}
			if validatorName == lowestColumnName {
				delete(validatorToColumnsCopy, validatorName)
			}
		}
		validatorToColumnSingular[lowestColumnName] = columnVal
	}
	return validatorToColumnSingular
}

func GetAllDepartureColumnsAndSum(validatorToColumnSingular map[string]int, input string) int {
	onTicketLine := false
	var myTicketInfo []int
	for _, line := range strings.Split(input, "\n") {
		if onTicketLine {
			myTicketInfo = LineToInts(line)
			break
		}
		if line == "your ticket:" {
			onTicketLine = true
		}
	}
	product := 1
	for name, column := range validatorToColumnSingular {
		if strings.Contains(name, "departure") {
			product = myTicketInfo[column] * product
		}
	}
	return product
}
