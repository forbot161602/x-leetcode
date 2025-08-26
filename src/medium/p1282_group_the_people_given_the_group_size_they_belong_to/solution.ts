// Runtime beats 95.08%, while memory beats 52.46%.
function groupThePeople(groupSizes: number[]): number[][] {
    const result: number[][] = [];
    const groups = new Map<number, number[]>();
    for (let i = 0; i < groupSizes.length; i++) {
        const size = groupSizes[i]!;
        let indexes = groups.get(size);
        if (indexes === undefined) {
            indexes = [];
            groups.set(size, indexes);
        }
        indexes.push(i);
        if (indexes.length === size) {
            result.push(indexes);
            groups.delete(size);
        }
    }
    return result;
}

export default groupThePeople;
