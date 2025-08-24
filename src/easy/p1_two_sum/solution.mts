// Runtime beats 65.71%, while memory beats 7.21%.
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
