package aoc18

import (
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

type ExecutionResult struct {
	endingIndex int
	total       int
}
