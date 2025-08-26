// Runtime beats 83.93%, while memory beats 16.07%.
function minDays(bloomDays: number[], targetBouquets: number, bouquetSize: number): number {
    const length = bloomDays.length;
    if (length < targetBouquets * bouquetSize) {
        return -1;
    }

    let lowerDay = 0;
    let upperDay = Math.max(...bloomDays);
    while (lowerDay < upperDay) {
        const middleDay = Math.floor((lowerDay + upperDay) / 2);
        if (canMakeTargetBouquets(bloomDays, middleDay, targetBouquets, bouquetSize)) {
            upperDay = middleDay;
        } else {
            lowerDay = middleDay + 1;
        }
    }
    return lowerDay;
}

function canMakeTargetBouquets(bloomDays: number[], targetDay: number, targetBouquets: number, bouquetSize: number): boolean {
    let flowerCount = 0;
    let bouquetCount = 0;
    for (const bloomDay of bloomDays) {
        if (bloomDay > targetDay) {
            flowerCount = 0;
            continue;
        }

        flowerCount++;
        if (flowerCount === bouquetSize) {
            bouquetCount++;
            flowerCount = 0;
            if (bouquetCount === targetBouquets) {
                return true;
            }
        }
    }
    return false;
}

export default minDays;
