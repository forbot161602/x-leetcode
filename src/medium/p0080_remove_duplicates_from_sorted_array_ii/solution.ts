// Runtime beats 95.41%, while memory beats 16.07%.
function removeDuplicates(values: number[]): number {
    if (values.length === 1) {
        return 1;
    }

    let index = 0;
    let count = 1;
    let current = values[0];
    for (let i = 1; i < values.length; i++) {
        const value = values[i]!;
        if (value !== current) {
            index++;
            count = 1;
            current = value;
            values[index] = value;
            continue;
        }
        if (count === 1) {
            index++;
            count++;
            values[index] = value;
        }
    }
    return index + 1;
}

export default removeDuplicates;
