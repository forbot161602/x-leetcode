// Runtime beats 100.00%, while memory beats 45.60%.
function twoSum(values: number[], target: number): number[] {
    let lower = 0;
    let upper = values.length - 1;
    while (lower < upper) {
        const value = values[lower]! + values[upper]!;
        if (value === target) {
            break;
        }
        if (value < target) {
            lower++;
        } else {
            upper--;
        }
    }
    return [lower + 1, upper + 1];
}

export default twoSum;
