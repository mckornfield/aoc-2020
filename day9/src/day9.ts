export function numberIsValidBasedOnPrevious(sampleNumberRange: number[], numberToCheck: number): boolean {
    const sumSet = createSumSet(sampleNumberRange);
    return sumSet.isNumberValid(numberToCheck);
}

export function findFirstInvalidNumber(listOfNumbersAsStr: string[], initialRangeSize: number = 25): number {
    const listOfNumbers = listOfNumbersAsStr.map(s => parseInt(s));
    const sumSet = createSumSet(listOfNumbers.slice(0, initialRangeSize));
    // console.log(listOfNumbersAsStr.slice(0, initialRangeSize));
    // console.log(listOfNumbers.slice(0, initialRangeSize));
    for (let i = initialRangeSize; i < listOfNumbers.length; i++) {
        const num = listOfNumbers[i];
        // console.log(num);
        if (sumSet.isNumberValid(num)) {
            sumSet.removeOldestSum();
            sumSet.addNumber(num);
        } else {
            return num;
        }
    }

    return -1;
}

const calculateSumSet = (currentNumbers: number[], newNumber: number) => {
    const sumSetForNumber = new Set<number>(currentNumbers.map(num => num + newNumber));
    // console.log(currentNumbers);
    return sumSetForNumber;
};

function createSumSet(numberRange: number[]): GroupOfSums {
    const sums: Set<number>[] = [];
    const currentNumbers = [...numberRange];
    for (let i = 0; i < numberRange.length; i++) {
        // Have to exclude the current number for this first setup
        const currentNumber = numberRange[i];
        const currentNumbersMinusCurrentOne = currentNumbers.slice(0, i).concat(currentNumbers.slice(i + 1));
        const sumSetForNumber = calculateSumSet(currentNumbersMinusCurrentOne, currentNumber);
        sums.push(sumSetForNumber);
    }
    return {
        removeOldestSum: () => { sums.shift(); currentNumbers.shift(); },

        addNumber: (newNumber: number) => {
            const sumSetForNumber = calculateSumSet(currentNumbers, newNumber);
            sums.push(sumSetForNumber);
            currentNumbers.push(newNumber);
        },
        isNumberValid: (newNumber: number) => sums.some(s => s.has(newNumber)),
    };
}

interface GroupOfSums {
    removeOldestSum: () => void,
    addNumber: (value: number) => void,
    isNumberValid: (value: number) => boolean;
}

export function findContiguousSumAndSumMinAndMax(input: string[], desiredSum: number): number {
    const sumRange = findContiguousSumThatReachesNumber(input, desiredSum);
    return Math.max(...sumRange) + Math.min(...sumRange);
}
export function findContiguousSumThatReachesNumber(input: string[], desiredSum: number): number[] {
    const numbers = input.map(s => parseInt(s));
    const sums: number[] = [];
    let sum = 0;
    const sumSnake: SumSnake = {
        shift: () => { sum -= sums.shift() || 0; },
        push: (elem: number) => { sums.push(elem); sum += elem; },
        getSum: () => sum,
        getNumbers: () => [...sums]
    };

    for (const num of numbers) {
        // console.log(sumSnake.getNumbers());
        // console.log(sumSnake.getSum());
        sumSnake.push(num);
        if (sumSnake.getSum() == desiredSum) {
            return sumSnake.getNumbers();
        } else if (sumSnake.getSum() > desiredSum) {
            while (sumSnake.getSum() > desiredSum) {
                sumSnake.shift();
                if (sumSnake.getSum() == desiredSum) {
                    return sumSnake.getNumbers();
                }
            }
        }
    }
    throw new Error("No sum snake found :(");

}

interface SumSnake {
    shift: () => void;
    push: (elem: number) => void;
    getSum: () => number;
    getNumbers: () => number[];
}
