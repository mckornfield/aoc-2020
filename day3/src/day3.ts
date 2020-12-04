export function parseGridInput(gridInput: string): Grid {
    const rows = gridInput.split('\n');
    const xMax = rows[0].length
    const yMax = rows.length
    return {
        get: function (x_y: Pair) {
            let { x, y } = x_y;
            while (x >= xMax) {
                x -= xMax;
            }
            while (y >= yMax) {
                y -= yMax;
            }
            // console.log(`y is ${y} x is ${x}`)
            return rows[y][x];
        },
        yMax: () => yMax,
    }
}

export function moveThroughGrid(grid: Grid, slope: Slope): number {
    let x = slope.right;
    let y = slope.down;
    let count = 0;
    while (y < grid.yMax()) {
        // Check for a tree
        let point = grid.get({ y, x })
        if (point == '#') {
            // console.log(`Tree found at ${x},${y}`)
            count++;
        }
        // Advance
        x += slope.right
        y += slope.down
    }

    return count;

}

export function moveThroughGridForSlopesAndMultiply(grid: Grid, slopes: Slope[]): number {
    const treesHitList = slopes
        .map(slope => moveThroughGrid(grid, slope));
    console.log(treesHitList);
    const multipliedTreesHit =
        treesHitList.reduce((accumulator, currentValue) => accumulator * currentValue, 1);;
    return multipliedTreesHit
}

// Measure from the top down
interface Grid {
    get(x_y: Pair): string
    yMax(): number
}

interface Pair {
    x: number,
    y: number
}

export interface Slope {
    right: number,
    down: number
}
