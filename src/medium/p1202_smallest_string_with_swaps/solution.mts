// Runtime beats 60.87%, while memory beats 30.44%.
function smallestStringWithSwaps(text: string, pairs: number[][]): string {
    const length = text.length;
    if (length === 1 || pairs.length === 0) {
        return text;
    }

    const roots = Array.from({length}, (_, i) => i);
    const sizes = Array<number>(length).fill(1);
    for (const [index1, index2] of pairs) {
        UniteRoots(roots, sizes, index1!, index2!);
    }

    const groups: Record<number, number[]> = {};
    for (let index = 0; index < length; index++) {
        const root = findRoot(roots, index);
        groups[root] ??= [];
        groups[root].push(index);
    }

    const result = Array<string>(length);
    for (const indexes of Object.values(groups)) {
        indexes.sort((i, j) => i - j);
        const chars = indexes.map(i => text[i]).sort();
        for (let i = 0; i < indexes.length; i++) {
            result[indexes[i]!] = chars[i]!;
        }
    }
    return result.join('');
}

function findRoot(roots: number[], index: number): number {
    if (roots[index] !== index) {
        roots[index] = findRoot(roots, roots[index]!);
    }
    return roots[index];
}

function UniteRoots(roots: number[], sizes: number[], index1: number, index2: number): void {
    let root1 = findRoot(roots, index1);
    let root2 = findRoot(roots, index2);
    if (root1 === root2) {
        return;
    }
    if (sizes[root1]! < sizes[root2]!) {
        [root1, root2] = [root2, root1];
    }
    roots[root2] = root1;
    sizes[root1]! += sizes[root2]!;
}

export default smallestStringWithSwaps;
