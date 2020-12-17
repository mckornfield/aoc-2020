package aoc14

import (
	"regexp"
	"strconv"
	"strings"
)

func ApplyMaskToValue(mask string, value int) int {
	keepMaskStr := strings.ReplaceAll(strings.ReplaceAll(mask, "1", "0"), "X", "1")
	keepMask, err := strconv.ParseInt(keepMaskStr, 2, 37)
	if err != nil {
		panic(err.Error())
	}
	overwriteMaskStr := strings.ReplaceAll(mask, "X", "0")
	overwriteMask, err2 := strconv.ParseInt(overwriteMaskStr, 2, 37)
	if err2 != nil {
		panic(err.Error())
	}
	return int(int64(value)&keepMask | overwriteMask)
}

func GetMemoryAddressesFromMask(mask string, memAddress int) []int {
	var firstBuilder strings.Builder
	builders := []*strings.Builder{&firstBuilder}
	for _, r := range mask {
		if r == '0' {
			for idx := range builders {
				builders[idx].WriteRune('X') // Unchanged, old way
			}
		} else if r == 'X' {
			// Copy builders for wild guy and append, but also add 0
			for idx := range builders {
				var splitBuilder strings.Builder
				splitBuilder.WriteString(builders[idx].String())
				builders[idx].WriteRune('0')
				splitBuilder.WriteRune('1')
				builders = append(builders, &splitBuilder)
			}
		} else if r == '1' {
			for idx := range builders {
				builders[idx].WriteRune('1') // Placeholder for an overwrite mask
			}
		}
	}
	memAddresses := []int{}
	for _, builder := range builders {
		memAddresses = append(memAddresses, ApplyMaskToValue(builder.String(), memAddress))
	}
	return memAddresses
}

const maskBeginning = "mask = "
const memBeginning = "mem"

func parseIntExplosively(possibleInt string) int {
	num, err := strconv.Atoi(possibleInt)
	if err != nil {
		panic(err.Error())
	}
	return num
}

func CreateMapWithBitMaskUpdates(input string) map[int]int {
	memoryMap := make(map[int]int)
	mask := ""
	for _, line := range strings.Split(input, "\n") {
		if strings.HasPrefix(line, maskBeginning) {
			mask = strings.Split(line, maskBeginning)[1]
		} else if strings.HasPrefix(line, memBeginning) {
			regexp, err := regexp.Compile(`mem\[(\d+)\] = (\d+)`)
			// fmt.Println(regexp)
			// fmt.Println(line)
			if err != nil {
				panic(err.Error())
			}
			results := regexp.FindStringSubmatch(line)
			address := parseIntExplosively(results[1])
			value := parseIntExplosively(results[2])
			maskedVal := ApplyMaskToValue(mask, value)
			memoryMap[address] = maskedVal
		} else {
			panic("could not match format of line" + line)
		}
	}
	return memoryMap
}

func CreateMapWithBitMaskUpdatesV2(input string) map[int]int {
	memoryMap := make(map[int]int)
	mask := ""
	for _, line := range strings.Split(input, "\n") {
		if strings.HasPrefix(line, maskBeginning) {
			mask = strings.Split(line, maskBeginning)[1]
		} else if strings.HasPrefix(line, memBeginning) {
			regexp, err := regexp.Compile(`mem\[(\d+)\] = (\d+)`)
			// fmt.Println(regexp)
			// fmt.Println(line)
			if err != nil {
				panic(err.Error())
			}
			results := regexp.FindStringSubmatch(line)
			address := parseIntExplosively(results[1])
			value := parseIntExplosively(results[2])
			memoryAddresses := GetMemoryAddressesFromMask(mask, address)
			for _, maskedMemoryAddress := range memoryAddresses {
				memoryMap[maskedMemoryAddress] = value
			}

		} else {
			panic("could not match format of line" + line)
		}
	}
	return memoryMap
}

func GetMemoryMapSum(memoryMap map[int]int) int {
	accumulator := 0
	for _, num := range memoryMap {
		accumulator += num
	}
	return accumulator
}
