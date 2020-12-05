import { readFileInputToString } from '../../reader/src/main';
import { countValidPassports, isPassportValid, validateKeyValuePair } from './day4';

test('is valid passport', () => {
    const isValid = isPassportValid(`ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
    byr:1937 iyr:2017 cid:147 hgt:183cm`)
    expect(isValid).toBe(true)
})
test('is invalid passport', () => {
    const isValid = isPassportValid(`iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
    hcl:#cfa07d byr:1929`)
    expect(isValid).toBe(false)
})


test('count 2 valid passports', () => {
    const inputStr = `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`
    const numPassports = countValidPassports(inputStr)
    expect(numPassports).toBe(2)
})

test('part 1, count all valid passports', () => {
    const input = readFileInputToString('day4');
    const count = countValidPassports(input);
    expect(count).toBe(242)
})

test('part 2 validate byr', () => {
    expect(validateKeyValuePair('byr:2002')).toBe(true)
    expect(validateKeyValuePair('byr:2003')).toBe(false)
    expect(validateKeyValuePair('byr:1919')).toBe(false)
    expect(validateKeyValuePair('byr:1920')).toBe(true)
    expect(validateKeyValuePair('byr:1955')).toBe(true)
})


test('part 2 validate hgt', () => {
    expect(validateKeyValuePair('hgt:60in')).toBe(true)
    expect(validateKeyValuePair('hgt:190cm')).toBe(true)
    expect(validateKeyValuePair('hgt:190in')).toBe(false)
    expect(validateKeyValuePair('hgt:190')).toBe(false)
})


test('part 2 validate iyr', () => {
    expect(validateKeyValuePair('iyr:2010')).toBe(true)
    expect(validateKeyValuePair('iyr:2020')).toBe(true)
    expect(validateKeyValuePair('iyr:2025')).toBe(false)
    expect(validateKeyValuePair('iyr:2030')).toBe(false)
})


test('part 2 validate eyr', () => {
    expect(validateKeyValuePair('eyr:2020')).toBe(true)
    expect(validateKeyValuePair('eyr:2030')).toBe(true)
    expect(validateKeyValuePair('eyr:2000')).toBe(false)
    expect(validateKeyValuePair('eyr:3000')).toBe(false)
})


test('part 2 validate hcl', () => {
    expect(validateKeyValuePair('hcl:#123abc')).toBe(true)
    expect(validateKeyValuePair('hcl:#123abz')).toBe(false)
    expect(validateKeyValuePair('hcl:#123abcd')).toBe(false)
    expect(validateKeyValuePair('hcl:123abc')).toBe(false)
})


test('part 2 validate ecl', () => {
    expect(validateKeyValuePair('ecl:brn')).toBe(true)
    expect(validateKeyValuePair('ecl:wat')).toBe(false)
})


test('part 2 validate pid', () => {
    expect(validateKeyValuePair('pid:000000001')).toBe(true)
    expect(validateKeyValuePair('pid:0123456789')).toBe(false)
})
