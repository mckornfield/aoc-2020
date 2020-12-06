export function determineAnswerCount(inputStr: string) {
    const uniqueChars = new Set(inputStr.split(''))
    uniqueChars.delete('\r')
    uniqueChars.delete('\n')
    uniqueChars.delete(' ')
    return uniqueChars.size
}

export function determineAnswerCountForList(inputs: string[]) {
    let answerInput = '';
    let count = 0;
    for (const line of inputs) {
        if (line && line.length > 0 && line.trim().length > 0) {
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