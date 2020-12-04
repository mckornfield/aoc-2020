import { isPasswordValid, PasswordInput, parseInput, isPasswordValidPt2 } from './password-verifier'
import { readFileInputByLines } from '../../reader/src/main';
const samplePasswords: [string, boolean][] = [
    ['1-3 a: abcde', true],
    ['1-3 b: cdefg', false],
    ['2-9 c: ccccccccc', true]
]
describe("test set of passwords", () => {
    test.each(samplePasswords)("test input '%s' validity is %s ",
        (input: string, expectedResult: boolean) => {
            expect(isPasswordValid(input)).toEqual(expectedResult);
        }
    )
})

test('7-10 z: gzjtmtcrzv is parsed appropriately', () => {
    const input = '7-10 z: gzjtmtcrzv';
    const pwInput: PasswordInput = parseInput(input);
    expect(pwInput.minimumCount).toBe(7)
    expect(pwInput.maximumCount).toBe(10)
    expect(pwInput.requiredChar).toBe("z")
    expect(pwInput.password).toBe("gzjtmtcrzv")
})


test('only 2 valid pws in list', () => {
    const samplePwsOnly: string[] = samplePasswords.map(l => l[0]);
    expect(samplePwsOnly.filter(input => isPasswordValid(input)).length).toBe(2);
})

test('Pt 1. only 560 valid pws in list', () => {
    const puzzlePws: string[] = readFileInputByLines('day2_puzzle.txt');
    expect(puzzlePws.filter(input => isPasswordValid(input)).length).toBe(560);
})


describe("test set of passwords for pt 2", () => {
    test.each([
        ['1-3 a: abcde', true],
        ['1-3 b: cdefg', false],
        ['2-9 c: ccccccccc', false]
    ])("test input '%s' validity is %s ",
        (input: string, expectedResult: boolean) => {
            expect(isPasswordValidPt2(input)).toEqual(expectedResult);
        }
    )
})


test('Pt 2. only N valid pws in list', () => {
    const puzzlePws: string[] = readFileInputByLines('day2_puzzle.txt');
    expect(puzzlePws.filter(input => isPasswordValidPt2(input)).length).toBe(303);
})
