import twoSum from './solution.js';

function main() {
    const solution = twoSum;
    console.log(solution([1, 2, 3], 4));
}

if (import.meta.url === `file://${process.argv[1]}`) {
    main();
}
