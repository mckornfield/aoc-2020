import { findProductOf2020Sum, findProductOf2020SumOfThreeNumbers } from './day1';

import { readFileInputByLines } from '../../reader/src/main';

test('expect the proper product for the sum input', () => {
  const values: number[] = [
    1721,
    979,
    366,
    299,
    675,
    1456,
  ]
  expect(findProductOf2020Sum(values)).toBe(514579)
})

test('puzzle input part 1', () => {
  const input = readFileInputByLines('day1_puzzle.txt').map(l => parseInt(l));
  expect(findProductOf2020Sum(input)).toBe(252724)
})


test('find product of 3 numbers', () => {
  const input: number[] = [
    1721,
    979,
    366,
    299,
    675,
    1456,
  ]
  expect(findProductOf2020SumOfThreeNumbers(input)).toBe(241861950)
})


test('part 2 input: find product of 3 numbers', () => {
  const input = readFileInputByLines('day1_puzzle.txt').map(l => parseInt(l));
  expect(findProductOf2020SumOfThreeNumbers(input)).toBe(276912720)
})
