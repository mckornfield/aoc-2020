import { readFileInputByLines } from './main'
test('read simple file', () => {
    expect(readFileInputByLines('sample.txt')).toStrictEqual(['hello', 'world'])
})
