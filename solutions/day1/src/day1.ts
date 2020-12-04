export function findProductOf2020Sum(vals: number[]): number {
    const VALUE_TO_SUM_TO = 2020;
    const sumPairs = new Map<number, number>()
    for (var i of vals) {
        const possibleVal = sumPairs.get(i)
        if (possibleVal) {
            return possibleVal * i;
        }
        sumPairs.set(VALUE_TO_SUM_TO - i, i);

    }
    return 0;
}

export function findProductOf2020SumOfThreeNumbers(vals: number[]): number {
    const VALUE_TO_SUM_TO = 2020;
    const sumMinusSingleVal = new Map<number, number>()
    const missingValueToPair = new Map<number, number[]>()
    for (let i of vals) {
        const possibleVal = missingValueToPair.get(i)
        if (possibleVal) {
            return possibleVal[0] * possibleVal[1] * i;
        }
        sumMinusSingleVal.set(VALUE_TO_SUM_TO - i, i);
        // Ew
        for (let v of sumMinusSingleVal.values()) {
            const possibleThirdVal = VALUE_TO_SUM_TO - i - v;
            if (possibleThirdVal > 0) {
                missingValueToPair.set(possibleThirdVal, [i, v])
            }
        }

    }
    return 0;
}
