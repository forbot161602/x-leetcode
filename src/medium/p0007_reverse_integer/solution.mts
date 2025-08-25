// Runtime beats 87.11%, while memory beats 7.15%.
function reverse(value: number): number {
    const isNonnegative = value >= 0;
    let reversed = 0;
    let quotient = isNonnegative ? value : -value;
    let remainder = 0;
    while (quotient !== 0) {
        remainder = quotient % 10;
        reversed = reversed * 10 + remainder;
        quotient = Math.floor(quotient / 10);
        if ((isNonnegative && reversed > MAX_INT32) || (!isNonnegative && -reversed < MIN_INT32)) {
            return 0;
        }
    }
    return isNonnegative ? reversed : -reversed;
}

const MAX_INT32 = 2147483647;
const MIN_INT32 = -2147483648;

export default reverse;
