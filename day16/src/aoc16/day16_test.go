package aoc16

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidator(t *testing.T) {
	assert := assert.New(t)
	input := "class: 1-3 or 5-7"
	validator := BuildValidator(input)
	assert.Equal(false, validator.validate(4))
	assert.Equal(true, validator.validate(1))
	assert.Equal(true, validator.validate(7))
	assert.Equal(false, validator.validate(8))
}

func TestSetOfValidators(t *testing.T) {
	assert := assert.New(t)
	input := `class: 1-3 or 5-7
	row: 6-11 or 33-44
	seat: 13-40 or 45-50
	
	your ticket:
	7,1,14
	
	nearby tickets:
	7,3,47
	40,4,50
	55,2,20
	38,6,12`
	invalidSum := FindSumOfInvalidInputs(input)
	assert.Equal(71, invalidSum)
}

func TestSumPt1(t *testing.T) {
	assert := assert.New(t)
	input := ReadFile(t, "day16_puzzle.txt")
	invalidSum := FindSumOfInvalidInputs(input)
	assert.Equal(28884, invalidSum)
}

func TestSumPt2Sample(t *testing.T) {
	assert := assert.New(t)
	input := `class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9`
	validatorsToColumns := FilterRowsByValidInput(input)
	validatorToSingleColumn := ValidatorToColumnMapping(validatorsToColumns)
	assert.Equal(map[string]int{
		"class": 1, "row": 0, "seat": 2,
	}, validatorToSingleColumn)
}

func TestSumPt2(t *testing.T) {
	assert := assert.New(t)
	input := ReadFile(t, "day16_puzzle.txt")
	validatorsToColumns := FilterRowsByValidInput(input)
	validatorToSingleColumn := ValidatorToColumnMapping(validatorsToColumns)
	sum := GetAllDepartureColumnsAndSum(validatorToSingleColumn, input)
	assert.Equal(1001849322119, sum)
}

func ReadFile(t *testing.T, fileName string) string {

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		t.Fatal(err.Error())
	}
	input := string(content)
	input = strings.TrimRight(input, "\n")
	return input
}
