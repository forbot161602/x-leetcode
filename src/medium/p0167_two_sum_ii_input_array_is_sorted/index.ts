import twoSum from './solution.js';

function main() {
    const solution = twoSum;
    console.log(solution([2, 7, 11, 15], 9));
}

if (import.meta.url === `file://${process.argv[1]}`) {
    main();
}
