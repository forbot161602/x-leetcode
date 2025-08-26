import removeDuplicates from './solution.js';

function main() {
    const solution = removeDuplicates;
    console.log(solution([0, 0, 1, 1, 1, 1, 2, 3, 3]));
}

if (import.meta.url === `file://${process.argv[1]}`) {
    main();
}
