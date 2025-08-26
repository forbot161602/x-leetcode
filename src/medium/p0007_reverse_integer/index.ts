import reverse from './solution.js';

function main() {
    const solution = reverse;
    console.log(solution(123));
}

if (import.meta.url === `file://${process.argv[1]}`) {
    main();
}
