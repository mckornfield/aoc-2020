export function determineAnswerCount(inputStr: string) {
    // console.log(inputStr);
    const uniqueChars = new Set(inputStr.split(''))
    uniqueChars.delete('\r')
    uniqueChars.delete('\n')
    uniqueChars.delete(' ')
    return uniqueChars.size
}

export function determineAnswerCountPt2(inputs: string[]) {
    const allAnsweredSet = inputs.map(input => new Set(input.split('')))
        .reduce((oneSet, otherSet) => {
            return new Set([...oneSet].filter(item => otherSet.has(item)))
        })

    allAnsweredSet.delete('\r')
    allAnsweredSet.delete('\n')
    allAnsweredSet.delete(' ')
    return allAnsweredSet.size
}

export function determineAnswerCountForList(inputs: string[]) {
    let answerInput = '';
    let count = 0;
    for (const line of inputs) {
        if (line) {
            answerInput += line;
        } else {
            count += determineAnswerCount(answerInput);
            answerInput = '';
        }
    }
    // Final count
    count += determineAnswerCount(answerInput);
    return count;
}

export function determineAnswerCountForListPt2(inputs: string[]) {
    let answerInput = [];
    let count = 0;
    for (const line of inputs) {
        if (line) {
            answerInput.push(line);
        } else {
            count += determineAnswerCountPt2(answerInput);
            answerInput = [];
        }
    }
    // Final count
    count += determineAnswerCountPt2(answerInput);
    return count;
}
