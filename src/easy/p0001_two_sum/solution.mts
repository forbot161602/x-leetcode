// Runtime beats 100.00%, while memory beats 11.10%.
function twoSum(numbers: number[], target: number): number[] {
    const compIndexes: Record<number, number> = {};
    for (const [index, number] of numbers.entries()) {
        const compIndex = compIndexes[number];
        if (compIndex !== undefined) {
            return [compIndex, index];
        }
        compIndexes[target - number] = index;
    }
    return [];
}

export default twoSum;
