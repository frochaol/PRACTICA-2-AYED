const INF = Number.MAX_SAFE_INTEGER;

function calculateDistance(city1, city2) {
    return Math.sqrt(Math.pow(city1[0] - city2[0], 2) + Math.pow(city1[1] - city2[1], 2));
}

function nearestInsertion(cities) {
    const n = cities.length;
    const distances = new Array(n).fill(null).map(() => new Array(n).fill(0));

    for (let i = 0; i < n; ++i) {
        for (let j = 0; j < n; ++j) {
            distances[i][j] = calculateDistance(cities[i], cities[j]);
        }
    }

    const tour = [0];
    const visited = new Array(n).fill(false);
    visited[0] = true;

    while (tour.length < n) {
        let minIncrease = INF;
        let bestCity = -1;

        for (let i = 0; i < tour.length; ++i) {
            for (let j = 0; j < n; ++j) {
                if (!visited[j]) {
                    const distanceIncrease =
                        distances[tour[i]][j] + distances[j][tour[(i + 1) % tour.length]] - distances[tour[i]][tour[(i + 1) % tour.length]];

                    if (distanceIncrease < minIncrease) {
                        minIncrease = distanceIncrease;
                        bestCity = j;
                        insertionIndex = i;
                    }
                }
            }
        }

        tour.splice(insertionIndex + 1, 0, bestCity);
        visited[bestCity] = true;
    }

    return tour;
}

// Generar datos aleatorios para n ciudades
const n = 5000;
const cities = [];
for (let i = 0; i < n; ++i) {
    const x = Math.floor(Math.random() * 100);
    const y = Math.floor(Math.random() * 100);
    cities.push([x, y]);
}

const startTime = new Date().getTime();
const tour = nearestInsertion(cities);
const endTime = new Date().getTime();
const executionTimeInSeconds = (endTime - startTime) / 1000;

// console.log("Recorrido del agente viajero:", tour.join(" -> "));
console.log("Ejecución para:", n, "ciudades");
console.log("Tiempo de ejecución:", executionTimeInSeconds, "segundos");
