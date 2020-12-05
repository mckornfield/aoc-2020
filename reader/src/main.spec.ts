import { readFileInputToString, readFileInputByLines } from './main'
test('read simple file', () => {
    expect(readFileInputByLines('sample')).toStrictEqual(['hello', 'world'])
})

test('read simple file with suffix', () => {
    expect(readFileInputByLines('sample_puzzle.txt')).toStrictEqual(['hello', 'world'])
})
test('read simple file with suffix', () => {
    expect(readFileInputToString('sample_puzzle.txt')).toEqual('hello\nworld\n')
})
