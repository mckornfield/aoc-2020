function traverseChildren(children: Set<string>, dag: Map<string, Set<string>>): Set<string> {
    const childrenToCheck = [...children]; // Use as queue
    const checkedChildren = new Set<string>();
    while (childrenToCheck.length > 0) {
        const childToCheck = childrenToCheck.pop()
        if (childToCheck && !checkedChildren.has(childToCheck)) {
            const possibleGrandChildren = dag.get(childToCheck);
            if (possibleGrandChildren) {
                childrenToCheck.push(...possibleGrandChildren);
            }

            checkedChildren.add(childToCheck);
        }
    }
    return checkedChildren;
}
export function buildBagRulesDag(bagRules: string): Map<string, Set<string>> {
    const dag = new Map<string, Set<string>>()
    bagRules.split('\n').forEach(line => {
        if (line.trim()) {
            const bagAndChildren = getBagAndChildren(line);
            const children = new Set(bagAndChildren.children);
            // const checkedChildren = traverseChildren(childrenToCheck, dag);
            dag.set(bagAndChildren.outerBag, children);
        }
    })
    // Find all children after one last go
    for (const entry of dag.entries()) {
        const [bag, childrenToCheck] = entry;
        const children = traverseChildren(childrenToCheck, dag);
        dag.set(bag, children);
    }
    return dag
}

export function calculateNumberOfBagsContaining(bagRules: string, bagToFind: string): number {
    const dag = buildBagRulesDag(bagRules);
    if (dag.has(bagToFind)) {
        let counter = 0;
        for (const childSet of dag.values()) {
            if (childSet.has(bagToFind)) {
                counter += 1;
            }
        }
        return counter;
    }
    throw new Error(`Could not find bag ${bagToFind} in dag ${dag}`)
}


export function getBagAndChildren(bagRule: string): BagAndChildren {
    const bagMatch = bagRule.match(/^(\w+ \w+) bags contain/);
    if (!bagMatch) {
        throw new Error(`No outerBag found for line ${bagRule}`)
    }
    const outerBag = bagMatch[1];

    const children = bagRule.includes('no other') ?
        [] :
        bagRule.split(' contain')[1]
            .split(',')
            .map(s => s.split(' ').slice(2, 4).join(' '));


    return {
        outerBag,
        children
    }
}


export function buildBagDagPt2(bagRules: string): Map<string, Set<ChildWithCount>> {
    const dag = new Map<string, Set<ChildWithCount>>()
    bagRules.split('\n').forEach(line => {
        if (line.trim()) {
            const bagAndChildren = getBagAndChildrenWithCount(line);
            const children = new Set(bagAndChildren.children);
            // const checkedChildren = traverseChildren(childrenToCheck, dag);
            dag.set(bagAndChildren.outerBag, children);
        }
    })
    return dag
}

export function getCountFromBags(bagDag: Map<string, Set<ChildWithCount>>, currentChild: ChildWithCount, runningCount: number): number {
    const grandChildrenSet = bagDag.get(currentChild.bag);
    if (grandChildrenSet?.size && grandChildrenSet?.size > 0) {
        return currentChild.count + [...grandChildrenSet]
            .map(grandChild => getCountFromBags(bagDag, grandChild, runningCount) * currentChild.count)
            .reduce((a, b) => a + b)
    } else {
        return runningCount * currentChild.count;
    }
}

export function calculateNumberOfBagsContainedBy(bagRules: string, bag: string): number {
    const dag = buildBagDagPt2(bagRules);
    const fakeStartingChild = {
        bag,
        count: 1
    }
    let count = getCountFromBags(dag, fakeStartingChild, 1) - 1; // Need to not include the bag itself
    return count;
}

export function getBagAndChildrenWithCount(bagRule: string): BagAndChildrenWithCounts {
    const bagMatch = bagRule.match(/^(\w+ \w+) bags contain/);
    if (!bagMatch) {
        throw new Error(`No outerBag found for line ${bagRule}`)
    }
    const outerBag = bagMatch[1];

    const children = bagRule.includes('no other') ?
        [] :
        bagRule.split(' contain')[1]
            .split(',')
            .map(s => s.split(' ').slice(1, 4).join(' '));

    const childrenWithCounts = children.map(s => {
        const splitBySpaces = s.split(" ");
        const count = parseInt(s.split(" ")[0]);
        const bag = splitBySpaces.slice(1).join(" ");
        return {
            count,
            bag
        }
    })

    return {
        outerBag,
        children: childrenWithCounts
    }
}

interface BagAndChildren {
    outerBag: string,
    children: string[]
}


interface ChildWithCount {
    bag: string
    count: number
}
interface BagAndChildrenWithCounts {
    outerBag: string,
    children: ChildWithCount[]
}
