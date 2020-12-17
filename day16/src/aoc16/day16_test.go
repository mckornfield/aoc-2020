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

func ReadFile(t *testing.T, fileName string) string {

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		t.Fatal(err.Error())
	}
	input := string(content)
	input = strings.TrimRight(input, "\n")
	return input
}
