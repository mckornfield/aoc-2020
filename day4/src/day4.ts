const requiredKeys = [
    "byr:",
    "iyr:",
    "eyr:",
    "hgt:",
    "hcl:",
    "ecl:",
    "pid:",
    // "cid",
]

export const validateKeyValuePair = (keyValuePair: string) => {
    const [key, value] = keyValuePair.split(':');
    switch (key) {
        case ('byr'):
            const birthYear = parseInt(value);
            return birthYear >= 1920 && birthYear <= 2002

        case ('iyr'):
            const issueYear = parseInt(value);
            return issueYear >= 2010 && issueYear <= 2020
        case ('eyr'):
            const expirationYear = parseInt(value);
            return expirationYear >= 2020 && expirationYear <= 2030
        case ('hgt'):
            const regex = /(\d+)(\w+)/;
            const matchResult = value.match(regex)

            if (!matchResult) {
                return false;
            }
            const numSegment = parseInt(matchResult[1])
            const unitSegment = matchResult[2]
            if (unitSegment == "cm") {
                return numSegment >= 150 && numSegment <= 193;
            } else if (unitSegment == "in") {
                return numSegment >= 59 && numSegment <= 76;
            }
            return false;
        case ('hcl'):
            return value.length == 7 && /#[0-9a-f]{6}/.test(value);
        case ('ecl'):
            const eyeColors = ['amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth']
            return eyeColors.includes(value)
        case ('pid'):
            return value.length == 9 && /\d{9}/.test(value)
    }
    return false;
}

export function countValidPassports(passportStrBlocks: string): number {
    let currentPassportSegments = [];
    let count = 0;
    for (const line of passportStrBlocks.split('\n')) {
        if (line.includes(':')) {
            currentPassportSegments.push(line);
        } else {
            const isValid = isPassportValid(currentPassportSegments.join(' '))
            if (isValid) {
                count += 1;
            }
            currentPassportSegments = []
        }
    }
    return count;
}

export function countValidPassportsPt2(passportStrBlocks: string): number {
    let currentPassportSegments = [];
    let count = 0;
    for (const line of passportStrBlocks.split('\n')) {
        if (line.includes(':')) {
            currentPassportSegments.push(line);
        } else {
            const isValid = isPassportValidPt2(currentPassportSegments.join(' '))
            if (isValid) {
                count += 1;
            }
            currentPassportSegments = []
        }
    }
    return count;
}

const NUM_REQUIRED_CHECKS = 7;
export function isPassportValidPt2(passportStr: string): boolean {

    const validChecks = passportStr.split(/\s+/).filter(
        validateKeyValuePair
    ).length
    // console.log(validChecks)
    // if (validChecks == 6) {
    // console.log(`6 Checks for ${passportStr}`)
    return validChecks == NUM_REQUIRED_CHECKS;

}

export function isPassportValid(passportStr: string): boolean {
    for (let key of requiredKeys) {
        if (passportStr.split(key).length != 2) {
            return false;
        }
    }
    return true

}
