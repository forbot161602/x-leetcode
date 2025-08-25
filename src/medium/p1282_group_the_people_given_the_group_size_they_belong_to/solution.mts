// Runtime beats 87.88%, while memory beats 66.67%.
function groupThePeople(groupSizes: number[]): number[][] {
    const result: number[][] = [];
    const groups: Record<number, number[]> = {};
    for (let i = 0; i < groupSizes.length; i++) {
        const size = groupSizes[i]!;
        let group = groups[size];
        if (group === undefined) {
            group = [];
            groups[size] = group;
        }
        group.push(i);
        if (group.length === size) {
            result.push(group);
            delete groups[size];
        }
    }
    return result;
}

export default groupThePeople;
