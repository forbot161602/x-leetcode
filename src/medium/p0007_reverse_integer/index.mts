import reverse from './solution.mjs';

function main() {
    const solution = reverse;
    console.log(solution(123));
}

if (import.meta.url === `file://${process.argv[1]}`) {
    main();
}
