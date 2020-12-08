import { readFileInputToString } from '../../reader/src/main';
import { buildBagDagPt2, buildBagRulesDag, calculateNumberOfBagsContaining, getBagAndChildren, getBagAndChildrenWithCount, calculateNumberOfBagsContainedBy } from './day7';
test('how many bags can hold a shiny bag', () => {
    const bagRules = `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`
    expect(calculateNumberOfBagsContaining(bagRules, "shiny gold")).toBe(4)
})
test('build bag tree', () => {
    const bagRules = `shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`
    const bagDag = buildBagRulesDag(bagRules)
    expect(bagDag.get('shiny gold')).toContain('dark olive')
    expect(bagDag.get('shiny gold')).toContain('vibrant plum')
    expect(bagDag.get('shiny gold')).toContain('dotted black')
    expect(bagDag.get('shiny gold')).toContain('faded blue')
    expect(bagDag.get('dark olive')).toContain('faded blue')
    expect(bagDag.get('dark olive')).toContain('dotted black')
    expect(bagDag.get('vibrant plum')).toContain('dotted black')
    expect(bagDag.get('vibrant plum')).toContain('dotted black')
    expect(bagDag.get('faded blue')?.size).toBe(0)
    expect(bagDag.get('dotted black')?.size).toBe(0)
})

test('get bag and children, 2 children', () => {
    const bagRule = 'shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.'
    const bagAndChildren = getBagAndChildren(bagRule);
    expect(bagAndChildren.outerBag).toBe('shiny gold')
    expect(bagAndChildren.children).toStrictEqual(['dark olive', 'vibrant plum'])
})

test('get bag and children, 4 children', () => {
    const bagRule = 'muted bronze bags contain 5 bright tomato bags, 5 light red bags, 2 shiny yellow bags, 2 dim teal bags.'
    const bagAndChildren = getBagAndChildren(bagRule);
    expect(bagAndChildren.outerBag).toBe('muted bronze')
    expect(bagAndChildren.children).toStrictEqual(['bright tomato', 'light red', 'shiny yellow', 'dim teal'])
})


test('get bag and children, 1 child', () => {
    const bagRule = 'bright plum bags contain 2 dim violet bags.'
    const bagAndChildren = getBagAndChildren(bagRule);
    expect(bagAndChildren.outerBag).toBe('bright plum')
    expect(bagAndChildren.children).toStrictEqual(['dim violet'])
})


test('get bag and children, 0 children', () => {
    const bagRule = 'dull white bags contain no other bags.'
    const bagAndChildren = getBagAndChildren(bagRule);
    expect(bagAndChildren.outerBag).toBe('dull white')
    expect(bagAndChildren.children).toStrictEqual([])
})


test('part one oh boy', () => {
    const bagRules = readFileInputToString('day7');
    expect(calculateNumberOfBagsContaining(bagRules, 'shiny gold')).toBe(261)
})

test('part two parse bag rule with counts', () => {
    const bagRule = `light red bags contain 1 bright white bag, 2 muted yellow bags, 6 dotted black bags.`
    const bagAndChildrenWithCount = getBagAndChildrenWithCount(bagRule);
    expect(bagAndChildrenWithCount.outerBag).toBe('light red');
    expect(bagAndChildrenWithCount.children.length).toBe(3)
    expect(bagAndChildrenWithCount.children[0]).toStrictEqual({ count: 1, bag: 'bright white' });
    expect(bagAndChildrenWithCount.children[1]).toStrictEqual({ count: 2, bag: 'muted yellow' });
    expect(bagAndChildrenWithCount.children[2]).toStrictEqual({ count: 6, bag: 'dotted black' });
    // expect(bagAndChildrenWithCount.outerBag).toBe('light red');
})

test('part two get bag data structure', () => {
    const bagRules = `shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.`
    const bagData = buildBagDagPt2(bagRules);
    expect(bagData.get('shiny gold')).toContainEqual({ bag: 'dark red', count: 2 })
    expect(bagData.get('dark blue')).toContainEqual({ bag: 'dark violet', count: 2 })
    expect(bagData.get('dark violet')?.size).toBe(0)
})

test('part two get bag counts', () => {
    const bagRules = `shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.`
    expect(calculateNumberOfBagsContainedBy(bagRules, 'shiny gold')).toBe(126);
})
test('part two get bag counts first example', () => {
    const bagRules = `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`
    expect(calculateNumberOfBagsContainedBy(bagRules, 'shiny gold')).toBe(32);
})

test('part two wow wtf I suck at recursion', () => {
    const bagRules = readFileInputToString('day7');
    expect(calculateNumberOfBagsContainedBy(bagRules, 'shiny gold')).toBe(3765);
})
