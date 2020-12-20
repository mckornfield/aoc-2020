package aoc18

import (
	"fmt"
	"strconv"
	"strings"
)

func isNumber(c rune) bool {
	_, err := strconv.ParseFloat(string(c), 10)
	return err == nil
}

func ExecuteOperation(currentNumStr *strings.Builder, currentOperation rune, total int) int {
	if currentNumStr.Len() != 0 {
		currentNum, err := strconv.ParseInt(currentNumStr.String(), 10, 32)
		if err != nil {
			panic(err)
		}
		currentNumStr.Reset()
		return ExecuteOperationWithOperand(int(currentNum), currentOperation, total)
	}
	return total
}

func ExecuteOperationWithOperand(currentNum int, currentOperation rune, total int) int {
	if currentOperation == '+' {
		total += currentNum
	} else if currentOperation == '*' {
		total = total * currentNum
	}
	return total
}

func SumOfAllLinesPt1(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		sum += ParseAndSolveMathProblem(line).total
	}
	return sum
}
func ParseAndSolveMathProblem(line string) ExecutionResult {
	total := 0
	var currentNumStr strings.Builder
	currentOperation := '+'
	for idx := 0; idx < len(line); idx++ {

		c := rune(line[idx])
		// fmt.Println(fmt.Sprintf("index %d, current char %c, total %d, current operation %c",
		// 	idx, c, total, currentOperation))
		// Lexing and parsing sounds hard, let's see how this dumb thing works
		if isNumber(c) {
			currentNumStr.WriteRune(c)
		} else if c == ' ' {
			total = ExecuteOperation(&currentNumStr, currentOperation, total)
		} else if c == '+' || c == '*' {
			currentOperation = c
			// Start the recursion BS
		} else if c == '(' {
			execResult := ParseAndSolveMathProblem(line[idx+1:])
			total = ExecuteOperationWithOperand(execResult.total, currentOperation, total)
			// Move index forward for parens
			idx += execResult.endingIndex
		} else if c == ')' {
			total = ExecuteOperation(&currentNumStr, currentOperation, total)
			return ExecutionResult{
				endingIndex: idx + 1,
				total:       total,
			}
		}
	}
	total = ExecuteOperation(&currentNumStr, currentOperation, total)
	return ExecutionResult{endingIndex: len(line) - 1, total: total}
}

func FindClosingParenIndex(line string) int {
	count := 0
	for idx, c := range line {
		if c == '(' {
			count++
		} else if c == ')' {
			count--
		}
		if count == 0 {
			return idx
		}
	}
	panic("No matching paren found for " + line)
}

func FindOpenParenIndex(line string) int {
	fmt.Println(line)
	// Assume line is in correct order and must be reversed
	count := 0
	for idx := len(line) - 1; idx > -1; idx-- {
		c := line[idx]
		if c == '(' {
			count++
		} else if c == ')' {
			count--
		}
		if count == 0 {
			return idx
		}
	}
	panic("No matching paren found for " + line)
}

func InsertParensAroundAddition(line string) string {
	var sb strings.Builder
	openParenIndices := make(map[int]bool)
	closingParenIndices := make(map[int]bool)
	// First determine where to insert parens

	for idx, c := range line {
		if (idx < len(line)-3) && line[idx+2:idx+3] == "+" {
			if c == ')' {
				openParenIndex := FindOpenParenIndex(line[:idx+1]) + 1
				fmt.Println(openParenIndex)
				openParenIndices[openParenIndex] = true
			} else {
				openParenIndices[idx] = true
			}
		}

		if idx > 1 && line[idx-2:idx-1] == "+" {
			if c == '(' {
				// Go on a matchin paren hunt
				closingParenIndex := FindClosingParenIndex(line[idx:]) + idx
				closingParenIndices[closingParenIndex] = true
			} else {
				closingParenIndices[idx] = true
			}
		}

	}

	for idx, c := range line {
		if _, ok := openParenIndices[idx]; ok {
			sb.WriteRune('(')
		}
		sb.WriteRune(c)

		if _, ok := closingParenIndices[idx]; ok {
			sb.WriteRune(')')
		}

	}
	return sb.String()
}
func ParseAndSolveMathProblemPt2(line string) int {
	fixedLine := InsertParensAroundAddition(line)
	return ParseAndSolveMathProblem(fixedLine).total
}

type ExecutionResult struct {
	endingIndex int
	total       int
}
