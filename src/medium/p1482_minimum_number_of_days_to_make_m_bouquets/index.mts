import minDays from './solution.mjs';

function main() {
    const solution = minDays;
    console.log(solution([7, 7, 7, 7, 12, 7, 7], 2, 3));
}

if (import.meta.url === `file://${process.argv[1]}`) {
    main();
}
