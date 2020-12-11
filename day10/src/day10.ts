export function getOneTimesThreeVoltageDifferences(input: string) {
    const voltages = input.split('\n').map(s => parseInt(s)).sort((a, b) => a - b);
    // conssole.log(voltages);
    // const last = voltages[voltages.length - 1];
    // const myDeviceRating = last + 3;
    let oneJumps = 0;
    let threeJumps = 1; // My device is the last 3 jump
    voltages.unshift(0); // Starting outlet is 0
    for (let i = 0; i < voltages.length; i++) {
        const difference = voltages[i + 1] - voltages[i];
        if (difference == 1) {
            oneJumps++;
        } else if (difference == 3) {
            threeJumps++;
        }
    }
    // console.log(`one jumps ${oneJumps} three jumps ${threeJumps}`);
    return oneJumps * threeJumps;
}

function findBranches(voltages: number[], index: number, memoizer: Map<number, number>): number {
    let result = 0;
    // Second to last
    if (memoizer.has(index)) {
        return memoizer.get(index) || 1;
    }
    if (index >= voltages.length - 2) {
        return 1;
    } else if (index == voltages.length - 3) {
        const nextDifference = voltages[index + 2] - voltages[index];
        return nextDifference < 4 ? 2 : 1;
    } else {
        const canSkipFirst = voltages[index + 2] - voltages[index] < 4;
        const canSkipSecond = voltages[index + 3] - voltages[index] < 4;
        if (canSkipSecond && canSkipFirst) {
            result = findBranches(voltages, index + 1, memoizer) +
                findBranches(voltages, index + 2, memoizer) +
                findBranches(voltages, index + 3, memoizer);
        } else if (canSkipFirst) {
            result = findBranches(voltages, index + 1, memoizer) +
                findBranches(voltages, index + 2, memoizer);
        } else {
            result = findBranches(voltages, index + 1, memoizer);
        }
    }
    memoizer.set(index, result);
    // console.log(`index=${index},voltage=${voltages[index]},result=${result}`);
    return result;

}
export function determineAllChargerConfigs(input: string) {
    const voltages = input.split('\n').map(s => parseInt(s)).sort((a, b) => a - b);
    let options = 1; // Always one way to go
    const optionsMapping = new Map<number, number[]>();
    // const branches = findBranches(voltages, 0);
    voltages.unshift(0); // Starting outlet is 0
    const paths = [voltages[0]];
    const memoizer = new Map<number, number>();
    const count = findBranches(voltages, 0, memoizer);
    // for (let i = 0; i < voltages.length - 1; i++) {
    // Assume contiguous, only check i+2 and i+3
    // const possibleOptions = [];
    // possibleOptions.push(voltages[i + 1]); // Assume first difference works
    // const difference2 = voltages[i + 2] - voltages[i];
    // if (difference2 < 4) {
    //     possibleOptions.push(voltages[i + 2]);
    // }
    // const difference3 = voltages[i + 3] - voltages[i];
    // if (difference3 < 4) {
    //     possibleOptions.push(voltages[i + 3]);
    // }
    // optionsMapping.set(voltages[i], possibleOptions);

    // }
    // console.log(optionsMapping);
    return count;
}
