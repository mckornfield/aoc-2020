import { Slope, moveThroughGridForSlopesAndMultiply, moveThroughGrid, parseGridInput } from './day3';
import { readFileInputByLines } from '../../reader/src/main';

const gridInput = `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`
test('test generating grid', () => {
    const grid = parseGridInput(gridInput);
    expect(grid.yMax()).toBe(11)
    expect(grid.get({ x: 0, y: 0 })).toBe('.')
    expect(grid.get({ x: 2, y: 0 })).toBe('#')
    expect(grid.get({ x: 10, y: 10 })).toBe('#')
    expect(grid.get({ x: 21, y: 21 })).toBe('#')
    expect(grid.get({ x: 0, y: 21 })).toBe('.')
})


test('test moving through grid', () => {
    const grid = parseGridInput(gridInput);
    const numberOfTreesEncountered = moveThroughGrid(grid, { right: 3, down: 1 });
    expect(numberOfTreesEncountered).toBe(7);
})


const puzzleGridInput = readFileInputByLines('day3_puzzle.txt').filter((line) => !!line).join('\n')

test('test part one moving through grid', () => {
    const grid = parseGridInput(puzzleGridInput);
    expect(grid.yMax()).toBe(323);
    expect(puzzleGridInput.split('\n')[grid.yMax() - 1]).toBe('.#...#.........#.#.##.#........');
    expect(grid.get({ x: 1, y: 322 })).toBe('#');
    const numberOfTreesEncountered = moveThroughGrid(grid, { right: 3, down: 1 });
    expect(numberOfTreesEncountered).toBe(292);
})


describe('test different slopes', () => {
    test.each([
        [{ right: 1, down: 1 }, 2],
        [{ right: 3, down: 1 }, 7],
        [{ right: 5, down: 1 }, 3],
        [{ right: 7, down: 1 }, 4],
        [{ right: 1, down: 2 }, 2],
    ])('test slope %s hits %i trees', (slope: Slope, treesHit: number) => {
        const grid = parseGridInput(gridInput);
        const numberOfTreesEncountered = moveThroughGrid(grid, slope);
        expect(numberOfTreesEncountered).toBe(treesHit)
    })
})

const slopes = [
    { right: 1, down: 1 },
    { right: 3, down: 1 },
    { right: 5, down: 1 },
    { right: 7, down: 1 },
    { right: 1, down: 2 },
]
test('multiply slopes together sample input', () => {
    const grid = parseGridInput(gridInput);
    const multipliedValues = moveThroughGridForSlopesAndMultiply(grid, slopes);
    expect(multipliedValues).toBe(336);
})

test('multiply numbers', () => {
    const nums = [2, 7, 3, 4, 2];
    const product = nums.reduce((accumulator, currentValue) => accumulator * currentValue, 1);
    expect(product).toBe(336);
})

test('part 2, multiply different slope options together', () => {
    const grid = parseGridInput(puzzleGridInput);
    const multipliedValues = moveThroughGridForSlopesAndMultiply(grid, slopes);
    expect(multipliedValues).toBe(5)
})
