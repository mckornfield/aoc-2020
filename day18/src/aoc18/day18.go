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

// 5 * 3 + 1 + 2 * 3 + 4 * 6 * 7
// 5 * 6 * 7 * 6 * 7

// 5 * 3 * 7 * 6 * 7
//  5 * 6 * 3 + 4 * 6 * 7

func PerformAdditionOnly(operands []int, operators []rune) []int {
	remainingNumbers := []int{}
	runningTotal := operands[0]
	// fmt.Println(operands, operators)
	for idx, operand := range operands[1:] {
		operator := operators[idx]
		if operator == '+' {
			runningTotal += operand
		} else {
			remainingNumbers = append(remainingNumbers, runningTotal)
			runningTotal = operand
		}
	}
	remainingNumbers = append(remainingNumbers, runningTotal)
	return remainingNumbers
}

func ExecuteOperationsWithAdditionFirst(operands []int, operators []rune) int {
	remainingNumbers := PerformAdditionOnly(operands, operators)
	// Multiply remaining numbers
	total := 1
	for _, remainingNumbers := range remainingNumbers {
		total *= remainingNumbers
	}
	return int(total)
}

func AppendAndResetIfApplicable(currentNumStr *strings.Builder, operands []int) []int {
	if currentNumStr.Len() > 0 {
		currentNum, err := strconv.ParseInt(currentNumStr.String(), 10, 32)
		if err != nil {
			panic(err)
		}
		operands = append(operands, int(currentNum))
		currentNumStr.Reset()
	}
	return operands
}

func ParseAndSolveMathProblemPt2(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		sum += ParseAndSolveMathProblemPt2Inner(line).total
	}
	return sum
}

func ParseAndSolveMathProblemPt2Inner(line string) ExecutionResult {
	total := 0
	var currentNumStr strings.Builder
	operands := []int{}
	operators := []rune{}
	// Tokenize first
	for idx := 0; idx < len(line); idx++ {
		c := rune(line[idx])

		// fmt.Println(fmt.Sprintf("index %d, rune '%c'", idx, c))
		// fmt.Println(fmt.Sprintf("index %d, current char %c, total %d, current operation %c",
		// 	idx, c, total, currentOperation))
		// Lexing and parsing sounds hard, let's see how this dumb thing works
		if isNumber(c) {
			currentNumStr.WriteRune(c)
		} else if c == ' ' {
			operands = AppendAndResetIfApplicable(&currentNumStr, operands)
		} else if c == '+' || c == '*' {
			operators = append(operators, c)
			// Start the recursion BS
		} else if c == '(' {
			execResult := ParseAndSolveMathProblemPt2Inner(line[idx+1:])
			operands = append(operands, execResult.total)
			// Move index forward for parens
			idx += execResult.endingIndex
		} else if c == ')' {
			operands = AppendAndResetIfApplicable(&currentNumStr, operands)
			total = ExecuteOperationsWithAdditionFirst(operands, operators)
			return ExecutionResult{
				endingIndex: idx + 1,
				total:       total,
			}
		}
		// fmt.Println(fmt.Sprintf("operands %v", operands))
		// fmt.Println(fmt.Sprintf("operands %v", RuneSliceToStringSlice(operators)))
	}
	operands = AppendAndResetIfApplicable(&currentNumStr, operands)
	total = ExecuteOperationsWithAdditionFirst(operands, operators)
	return ExecutionResult{endingIndex: len(line) - 1, total: total}
}

func RuneSliceToStringSlice(runes []rune) []string {
	vals := []string{}
	for _, c := range runes {
		vals = append(vals, string(c))
	}
	return vals
}

type ExecutionResult struct {
	endingIndex int
	total       int
}
