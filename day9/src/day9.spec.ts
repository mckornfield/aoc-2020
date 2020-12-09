import { readFileInputByLinesSkipLastLine } from '../../reader/src/main';
import { findContiguousSumAndSumMinAndMax, findContiguousSumThatReachesNumber, findFirstInvalidNumber, numberIsValidBasedOnPrevious } from './day9';
test('sum of previous 5 numbers, first mismatch', () => {
    const listOfNumbers = `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`;
    expect(findFirstInvalidNumber(listOfNumbers.split('\n'), 5)).toBe(127);
});

test('pt 1 let"s give it a shot"', () => {
    const lines = readFileInputByLinesSkipLastLine('day9');
    expect(findFirstInvalidNumber(lines)).toBe(133015568);
});

const range = (start: number, size: number) => [...Array(size).keys()].map(i => i + start);

const sampleNumberRange = range(1, 25);

describe('test if numbers are valid based on previous 25', () => {
    test.each([
        [26, true],
        [49, true],
        [100, false],
        [50, false],
    ])('test %s is considered valid: %s', (numberToCheck: number, isValid: boolean) => {
        expect(numberIsValidBasedOnPrevious(sampleNumberRange, numberToCheck)).toBe(isValid);
    });
});

describe('test if numbers are valid based on previous 25, slighlty different range', () => {
    test.each([
        [26, true],
        [65, false],
        [64, true],
        [66, true],
    ])('test %s is considered valid: %s', (numberToCheck: number, isValid: boolean) => {
        const updatedNumberRange = [...sampleNumberRange].filter(i => i != 20);
        updatedNumberRange.push(45);
        expect(numberIsValidBasedOnPrevious(updatedNumberRange, numberToCheck)).toBe(isValid);
    });
});

test('try and activate the pt2 sum snake', () => {
    const input = `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`.split('\n');
    const desiredSum = 127;
    expect(findContiguousSumThatReachesNumber(input, desiredSum)).toStrictEqual([15, 25, 47, 40,]);
    expect(findContiguousSumAndSumMinAndMax(input, desiredSum)).toBe(62);
});


test('unleash the sum snake on puzzle input', () => {
    const lines = readFileInputByLinesSkipLastLine('day9');
    expect(findContiguousSumAndSumMinAndMax(lines, 133015568)).toBe(16107959);
});
