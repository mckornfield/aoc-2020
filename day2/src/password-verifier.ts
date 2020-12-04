export function isPasswordValid(input: string): boolean {
    const { minimumCount, maximumCount, password, requiredChar } = parseInput(input);
    const requiredCharOccurences = password.split(requiredChar).length - 1
    return requiredCharOccurences >= minimumCount && requiredCharOccurences <= maximumCount;
}


export function isPasswordValidPt2(input: string): boolean {
    const { minimumCount, maximumCount, password, requiredChar } = parseInput(input);
    // Caution: 1 indexing and XOR
    return (password[minimumCount - 1] == requiredChar && !(password[maximumCount - 1] == requiredChar)) ||
        (!(password[minimumCount - 1] == requiredChar) && password[maximumCount - 1] == requiredChar);
}

export function parseInput(input: string): PasswordInput {
    const regexpPw: RegExp = /^(\d+)-(\d+) (\w): (\w+)/
    const match = input.match(regexpPw);
    if (!match) {
        throw new Error(`Invalid format ${input}`)
    }
    return {
        minimumCount: parseInt(match[1]),
        maximumCount: parseInt(match[2]),
        requiredChar: match[3],
        password: match[4]
    }
}

export interface PasswordInput {
    minimumCount: number;
    maximumCount: number;
    requiredChar: string;
    password: string;
}
