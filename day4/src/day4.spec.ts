import { readFileInputToString } from '../../reader/src/main';
import { countValidPassports, countValidPassportsPt2, isPassportValid, isPassportValidPt2, validateKeyValuePair } from './day4';

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
    expect(validateKeyValuePair('hgt:193cm')).toBe(true)
    expect(validateKeyValuePair('hgt:190cm')).toBe(true)
    expect(validateKeyValuePair('hgt:150cm')).toBe(true)
    expect(validateKeyValuePair('hgt:76in')).toBe(true)
    expect(validateKeyValuePair('hgt:77in')).toBe(false)
    expect(validateKeyValuePair('hgt:149cm')).toBe(false)
    expect(validateKeyValuePair('hgt:194cm')).toBe(false)
    expect(validateKeyValuePair('hgt:190in')).toBe(false)
    expect(validateKeyValuePair('hgt:165in')).toBe(false)
    expect(validateKeyValuePair('hgt:190ix')).toBe(false)
    expect(validateKeyValuePair('hgt:190')).toBe(false)
    expect(validateKeyValuePair('hgt:60')).toBe(false)
})


test('part 2 validate iyr', () => {
    expect(validateKeyValuePair('iyr:2010')).toBe(true)
    expect(validateKeyValuePair('iyr:2015')).toBe(true)
    expect(validateKeyValuePair('iyr:2020')).toBe(true)
    expect(validateKeyValuePair('iyr:1')).toBe(false)
    expect(validateKeyValuePair('iyr:20000')).toBe(false)
    expect(validateKeyValuePair('iyr:2025')).toBe(false)
    expect(validateKeyValuePair('iyr:2030')).toBe(false)
})


test('part 2 validate eyr', () => {
    expect(validateKeyValuePair('eyr:2020')).toBe(true)
    expect(validateKeyValuePair('eyr:2025')).toBe(true)
    expect(validateKeyValuePair('eyr:2030')).toBe(true)
    expect(validateKeyValuePair('eyr:2000')).toBe(false)
    expect(validateKeyValuePair('eyr:3000')).toBe(false)
})


test('part 2 validate hcl', () => {
    expect(validateKeyValuePair('hcl:#123abc')).toBe(true)
    expect(validateKeyValuePair('hcl:#123abz')).toBe(false)
    expect(validateKeyValuePair('hcl:#123abcd')).toBe(false)
    expect(validateKeyValuePair('hcl:#cfa07d')).toBe(true)
})


describe('eye colors', () => test.each(
    ['amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth']
)('part 2 ecl %s', (eyeColor: string) =>
    expect(validateKeyValuePair(`ecl:${eyeColor}`)).toBe(true)))

test('part 2 validate ecl', () => {
    expect(validateKeyValuePair('ecl:brn')).toBe(true)

    expect(validateKeyValuePair('ecl:wat')).toBe(false)
})


test('part 2 validate pid', () => {
    expect(validateKeyValuePair('pid:000000001')).toBe(true)
    expect(validateKeyValuePair('pid:0123456789')).toBe(false)
})

test('check some invalid passports', () => {
    expect(isPassportValidPt2(`eyr:1972 cid:100
    hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926`)).toBe(false)
    expect(isPassportValidPt2(`iyr:2019
    hcl:#602927 eyr:1967 hgt:170cm
    ecl:grn pid:012533040 byr:1946`)).toBe(false)
    expect(isPassportValidPt2(`hcl:dab227 iyr:2012
    ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:27`)).toBe(false)
    expect(isPassportValidPt2(`hgt:59cm ecl:zzz
    eyr:2038 hcl:74454a iyr:2023
    pid:3556412378 byr:2007`)).toBe(false)
})



test('check some valid passports', () => {
    expect(isPassportValidPt2(`pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
    hcl:#623a2f`)).toBe(true)
    expect(isPassportValidPt2(`eyr:2029 ecl:blu cid:129 byr:1989
    iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm`)).toBe(true)
    expect(isPassportValidPt2(`hcl:#888785
    hgt:164cm byr:2001 iyr:2015 cid:88
    pid:545766238 ecl:hzl
    eyr:2022`)).toBe(true)
    expect(isPassportValidPt2(`iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719`)).toBe(true)
})

test('count 4 valid passports', () => {
    const input = `pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719

eyr:1972 cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946

hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007

pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007
`
    expect(countValidPassportsPt2(input)).toBe(5)

})


test('part 2, count all valid passports', () => {
    const input = readFileInputToString('day4');
    const count = countValidPassportsPt2(input);
    expect(count).toBe(186)
})

test('part 2, passport should be valid', () => {
    expect(isPassportValidPt2('byr:1924 cid:321 eyr:2028 hcl:#cfa07d iyr:2010 ecl:amb pid:036669613 hgt:170cm')).toBe(true)
})
