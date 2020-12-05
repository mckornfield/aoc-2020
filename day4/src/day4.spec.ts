import { countValidPassports, isPassportValid } from './day4'
import { readFileInputToString } from '../../reader/src/main';

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
