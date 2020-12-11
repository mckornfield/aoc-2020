import { readFileInputToString } from '../../reader/src/main';
import { determineAllChargerConfigs, getOneTimesThreeVoltageDifferences } from './day10';
test('find 1 to 3 volt differences', () => {
    const voltageList = `16
10
15
5
1
11
7
19
6
12
4`;
    expect(getOneTimesThreeVoltageDifferences(voltageList)).toBe(35);
});
test('find 1 to 3 volt differences, second example', () => {
    const voltageList = `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`;
    expect(getOneTimesThreeVoltageDifferences(voltageList)).toBe(220);
});

test('find 1 to 3 volt differences, second example', () => {
    const input = readFileInputToString('day10');
    expect(getOneTimesThreeVoltageDifferences(input)).toBe(2775);
});


test('find all possible combos', () => {
    const voltageList = `16
10
15
5
1
11
7
19
6
12
4`;
    expect(determineAllChargerConfigs(voltageList)).toBe(8);
});


test('find all possible combos second example', () => {
    const voltageList = `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`;
    expect(determineAllChargerConfigs(voltageList)).toBe(19208);
});


test('part 2 lez go', () => {
    const input = readFileInputToString('day10');
    expect(determineAllChargerConfigs(input)).toBe(518344341716992);
});
