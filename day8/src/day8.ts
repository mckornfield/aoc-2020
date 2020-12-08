export enum OperationType {
    Accumulate = 'Accumulate',
    Jump = 'Jump',
    Noop = 'Noop',
}

export enum Direction {
    Forward = 'Forward',
    Backward = 'Backward',
}

interface Operation {
    type: OperationType,
    direction: Direction,
    quantity: number,
}


const OPERATION_TYPE_MAPPING = new Map<string, OperationType>(
    [
        ['acc', OperationType.Accumulate],
        ['jmp', OperationType.Jump],
        ['nop', OperationType.Noop],
    ])

export function parseLine(line: string): Operation {
    const match = line.match(/(\w+) ([+-])(\d+)/)
    if (match) {
        const type = OPERATION_TYPE_MAPPING.get(match[1]) || OperationType.Noop
        const direction = match[2] === "+" ? Direction.Forward : Direction.Backward
        const quantity = parseInt(match[3])
        return {
            type,
            direction,
            quantity
        }
    }
    throw new Error(`No match for line '${line}'`)
}

export function advanceStateMachineAndReturnAccValue(input: string[]): number {
    let acc = 0;
    const visitedInstructions = new Set<number>();
    const instructions = input.map(parseLine);
    let currentIndex = 0;
    while (!visitedInstructions.has(currentIndex) && currentIndex < instructions.length) {
        visitedInstructions.add(currentIndex)
        const { direction, quantity, type } = instructions[currentIndex];
        // console.log(instructions[currentIndex])
        if (type == OperationType.Noop) {
            currentIndex += 1;
            continue;
        }
        const isForward = direction == Direction.Forward;
        const quantityWithSign = isForward ? quantity : -quantity;
        switch (type) {
            case (OperationType.Jump):
                currentIndex += quantityWithSign
                break;
            case (OperationType.Accumulate):
                const additionWithSign = quantityWithSign;
                acc += additionWithSign;
                currentIndex += 1;
        }

    }
    return acc;
}


interface StateMachineRun {
    acc: number,
    programCompletedNormally: boolean,
}

export function permuteStateOperationsUntilSuccessfulRun(input: string[]): number {
    const instructions = input.map(parseLine);
    for (let i = 0; i < instructions.length; i++) {
        const { type, direction, quantity } = instructions[i];
        if (type == OperationType.Accumulate) {
            continue; // Don't permute accumulate indices
        }
        const permutedInstructions = [...instructions];
        permutedInstructions[i] = {
            type: type == OperationType.Jump ? OperationType.Noop : OperationType.Jump,
            direction,
            quantity
        }
        const result = runStateMachine(permutedInstructions)
        if (result.programCompletedNormally) {
            return result.acc;
        }
    }
    throw new Error("Could not permute program to work")
}

export function runStateMachine(instructions: Operation[]): StateMachineRun {
    let acc = 0;
    const visitedInstructions = new Set<number>();
    let currentIndex = 0;
    while (!visitedInstructions.has(currentIndex) && currentIndex < instructions.length) {
        visitedInstructions.add(currentIndex)
        const { direction, quantity, type } = instructions[currentIndex];
        if (type == OperationType.Noop) {
            currentIndex += 1;
            continue;
        }
        const isForward = direction == Direction.Forward;
        const quantityWithSign = isForward ? quantity : -quantity;
        switch (type) {
            case (OperationType.Jump):
                currentIndex += quantityWithSign
                break;
            case (OperationType.Accumulate):
                const additionWithSign = quantityWithSign;
                acc += additionWithSign;
                currentIndex += 1;
        }

    }
    return {
        acc,
        programCompletedNormally: instructions.length == currentIndex,
    }
}
