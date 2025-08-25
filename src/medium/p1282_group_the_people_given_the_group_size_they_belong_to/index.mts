import groupThePeople from './solution.mjs';

function main() {
    const solution = groupThePeople;
    console.log(solution([2, 1, 3, 3, 3, 2]));
}

if (import.meta.url === `file://${process.argv[1]}`) {
    main();
}
