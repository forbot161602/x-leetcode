// Runtime beats 100.00%, while memory beats 11.10%.
function twoSum(values: number[], target: number): number[] {
    const compIndexes = new Map<number, number>();
    for (const [index, value] of values.entries()) {
        const compIndex = compIndexes.get(value);
        if (compIndex !== undefined) {
            return [compIndex, index];
        }
        compIndexes.set(target - value, index);
    }
    return [];
}

export default twoSum;
