package aoc14

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitwiseMaskSample(t *testing.T) {
	assert := assert.New(t)
	mask := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"
	value := 11
	assert.Equal(73, ApplyMaskToValue(mask, value))
}

func TestBitwiseMaskSampleWithMap(t *testing.T) {
	assert := assert.New(t)
	input := `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`
	memoryMap := CreateMapWithBitMaskUpdates(input)
	assert.Equal(64, memoryMap[8])
	assert.Equal(101, memoryMap[7])
	assert.Equal(165, GetMemoryMapSum(memoryMap))
}

func TestBitwiseMaskPt1(t *testing.T) {
	assert := assert.New(t)
	contents := ReadFile(t, "day14_puzzle.txt")

	memoryMap := CreateMapWithBitMaskUpdates(contents)
	assert.Equal(15018100062885, GetMemoryMapSum(memoryMap))
}

func TestBitwiseMaskPt2GetAddresses(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]int{26, 58, 27, 59}, GetMemoryAddressesFromMask("000000000000000000000000000000X1001X", 42))
}

func TestBitwiseMaskPt2Sample(t *testing.T) {
	assert := assert.New(t)
	input := `mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`
	memoryMap := CreateMapWithBitMaskUpdatesV2(input)
	assert.Equal(208, GetMemoryMapSum(memoryMap))
}

func TestBitwiseMaskPt2(t *testing.T) {
	assert := assert.New(t)
	contents := ReadFile(t, "day14_puzzle.txt")

	memoryMap := CreateMapWithBitMaskUpdatesV2(contents)
	assert.Equal(15018100062885, GetMemoryMapSum(memoryMap))
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
