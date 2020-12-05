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

export function countValidPassports(passportStrBlocks: string): number {
    let currentPassportSegments = [];
    let count = 0;
    // console.log(passportStrBlocks)
    for (const line of passportStrBlocks.split('\n')) {
        if (line.includes(':')) {
            currentPassportSegments.push(line);
        } else {
            // console.log(currentPassportSegments)
            const isValid = isPassportValid(currentPassportSegments.join(' '))
            if (isValid) {
                count += 1;
            }
            currentPassportSegments = []
        }
    }
    return count;
}

export function isPassportValid(passportStr: string): boolean {
    for (let key of requiredKeys) {
        if (!passportStr.includes(key)) {
            return false;
        }
    }
    return true

}
