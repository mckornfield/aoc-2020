import { readFileInputByLinesSkipLastLine } from '../../reader/src/main';
import { OperationType, Direction, parseLine, advanceStateMachineAndReturnAccValue } from './day8'
describe('parse line', () => {

    test.each([
        ['acc +1', OperationType.Accumulate, Direction.Forward, 1],
        ['nop +0', OperationType.Noop, Direction.Forward, 0],
        ['acc -99', OperationType.Accumulate, Direction.Backward, 99],
        ['jmp -4', OperationType.Jump, Direction.Backward, 4],
    ])('parse line %s into %s , %s , %s', (line: string, operationType: OperationType,
        direction: Direction, quantity: number) => {
        const operation = parseLine(line);
        expect(operation.type).toBe(operationType);
        expect(operation.direction).toBe(direction);
        expect(operation.quantity).toBe(quantity);
    })
})

test('advance state machine, no loop', () => {
    const stateInput = `nop +0
acc +1
jmp +4
acc +3
nop +0
acc +1
nop +0
acc +4`
    expect(advanceStateMachineAndReturnAccValue(stateInput.split('\n'))).toBe(5)
})

test('advance state machine, with loop', () => {
    const stateInput = `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`
    expect(advanceStateMachineAndReturnAccValue(stateInput.split('\n'))).toBe(5)
})



test('pt 1', () => {
    const stateInput = readFileInputByLinesSkipLastLine('day8');
    expect(advanceStateMachineAndReturnAccValue(stateInput)).toBe(1594)
})

test('flip jmps to nops and find accumulated value', () => {
    
})
