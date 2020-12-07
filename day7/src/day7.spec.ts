import { readFileInputToString } from '../../reader/src/main';
import { buildBagRulesDag, calculateNumberOfBagsContaining, getBagAndChildren } from './day7';
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
    expect(calculateNumberOfBagsContaining(bagRules,'shiny gold')).toBe(261)
})
