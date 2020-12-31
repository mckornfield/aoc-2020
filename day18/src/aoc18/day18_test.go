package aoc18

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeirdMathSolverPt1Sample(t *testing.T) {
	assert := assert.New(t)
	line := "1 + 2 * 3 + 4 * 5 + 6"
	assert.Equal(71, ParseAndSolveMathProblem(line).total)
	assert.Equal(20, ParseAndSolveMathProblem(line).endingIndex)
}

func TestWeirdMathSolverPt1SampleWithParens(t *testing.T) {
	assert := assert.New(t)
	line := "1 + (2 * 3) + (4 * (5 + 6))"
	assert.Equal(51, ParseAndSolveMathProblem(line).total)
	assert.Equal(26, ParseAndSolveMathProblem("2 * 3 + (4 * 5)").total)
	assert.Equal(437, ParseAndSolveMathProblem("5 + (8 * 3 + 9 + 3 * 4 * 3)").total)
	assert.Equal(12240, ParseAndSolveMathProblem("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))").total)
	assert.Equal(13632, ParseAndSolveMathProblem("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2").total)
}

func TestSumEmAllSamplePt1(t *testing.T) {
	assert := assert.New(t)

	lines := `"2 * 3 + (4 * 5)")
"5 + (8 * 3 + 9 + 3 * 4 * 3)")
"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))")
"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2")`

	assert.Equal(26+437+12240+13632, SumOfAllLinesPt1(lines))

}

func TestPt1(t *testing.T) {
	assert := assert.New(t)

	lines := ReadFile(t, "day18_puzzle.txt")

	assert.Equal(36382392389406, SumOfAllLinesPt1(lines))

}

// 1 + 2 * 3 + 4 * 6
// 3 * 7 * 6
// 126
func TestExecuteOperationsWithAdditionFirst(t *testing.T) {
	assert := assert.New(t)
	operands := []int{1, 2, 3, 4, 6}
	operators := []rune{'+', '*', '+', '*'}
	assert.Equal(126, ExecuteOperationsWithAdditionFirst(operands, operators))
	assert.Equal([]int{3, 7, 6}, PerformAdditionOnly(operands, operators))
}

// 5 * 3 + 1 + 2 * 3 + 4 * 6 * 7
// 5 * 6 * 7 * 6 * 7

func TestExecuteOperationsWithAdditionFirstSecondExample(t *testing.T) {
	assert := assert.New(t)
	operands := []int{5, 3, 1, 2, 3, 4, 6, 7}
	operators := []rune{'*', '+', '+', '*', '+', '*', '*'}
	assert.Equal(8820, ExecuteOperationsWithAdditionFirst(operands, operators))
	assert.Equal([]int{5, 6, 7, 6, 7}, PerformAdditionOnly(operands, operators))
}

func TestPt2SampleWithAdditionFirst(t *testing.T) {
	assert := assert.New(t)

	line := "2 * 3 + (4 * 5)"
	assert.Equal(46, ParseAndSolveMathProblemPt2(line))
	assert.Equal(1445, ParseAndSolveMathProblemPt2("5 + (8 * 3 + 9 + 3 * 4 * 3)"))
	assert.Equal(669060, ParseAndSolveMathProblemPt2("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"))

	line = "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"
	assert.Equal(23340, ParseAndSolveMathProblemPt2(line))
}

func TestPt2(t *testing.T) {
	assert := assert.New(t)

	lines := ReadFile(t, "day18_puzzle.txt")

	assert.Equal(36382392389406, ParseAndSolveMathProblemPt2(lines))

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
