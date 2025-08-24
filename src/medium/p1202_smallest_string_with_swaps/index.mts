import smallestStringWithSwaps from './solution.mjs';

function main() {
    const solution = smallestStringWithSwaps;
    console.log(solution('dcab', [[0, 3], [1, 2]]));
}

if (import.meta.url === `file://${process.argv[1]}`) {
    main();
}
