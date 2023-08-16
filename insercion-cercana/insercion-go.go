package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const INF = math.MaxInt32

type City struct {
	X int
	Y int
}

func calculateDistance(city1, city2 City) float64 {
	return math.Sqrt(math.Pow(float64(city1.X-city2.X), 2) + math.Pow(float64(city1.Y-city2.Y), 2))
}

func nearestInsertion(cities []City) []int {
	n := len(cities)
	distances := make([][]float64, n)
	for i := range distances {
		distances[i] = make([]float64, n)
		for j := range distances[i] {
			distances[i][j] = calculateDistance(cities[i], cities[j])
		}
	}

	tour := []int{0}
	visited := make([]bool, n)
	visited[0] = true

	for len(tour) < n {
		minIncrease := float64(INF)
		var bestCity, insertionIndex int

		for i, currentCity := range tour {
			if visited[currentCity] {
				for j, visitedCity := range visited {
					if !visitedCity {
						distanceIncrease := distances[currentCity][j] + distances[j][tour[(i+1)%len(tour)]] - distances[currentCity][tour[(i+1)%len(tour)]]

						if distanceIncrease < minIncrease {
							minIncrease = distanceIncrease
							bestCity = j
							insertionIndex = i
						}
					}
				}
			}
		}

		tour = append(tour[:insertionIndex+1], append([]int{bestCity}, tour[insertionIndex+1:]...)...)
		visited[bestCity] = true
	}

	return tour
}

func main() {
	rand.Seed(time.Now().UnixNano())

	n := 5000
	cities := make([]City, n)
	for i := 0; i < n; i++ {
		x := rand.Intn(100)
		y := rand.Intn(100)
		cities[i] = City{X: x, Y: y}
	}

	startTime := time.Now()
	nearestInsertion(cities)
	endTime := time.Now()
	executionTime := endTime.Sub(startTime).Seconds()

	fmt.Println("Ejecución para programa optimizado para:", n, "ciudades")
	// fmt.Println("Recorrido del agente viajero:", tour)
	fmt.Println("Tiempo de ejecución:", executionTime, "segundos")
}
