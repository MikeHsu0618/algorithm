package main

// There are N cars, numbered from 0 to N-1. Each of them has some of the M possible optional features,
// numbered from 0 to M-1, for example: voice control, keyless entry, sunroof, blind spot detection, etc.
// The features of a car are described as a string of M characters, where the K-th character being '1' denotes
// that the car has the K-th possible feature and '0' denotes that it does not.
// Two cars are similar if their descriptions differ by at most one feature.
// For example: "01101" and "01001" are similar, because they differ only by feature number 2.
// On the other hand, "01101" and "11110" are not similar, because they differ in feature numbers 0, 3 and 4.
//
// Write a function that, given an array  consisting of N strings, returns an array of integers denoting,
// for every car in , the number of other similar cars.

func similarCars(cars []string) []int {
	counts := make([]int, len(cars)) // array to store the counts for each car
	for i, car := range cars {
		for j, otherCar := range cars {
			if i == j {
				continue // don't compare a car with itself
			}
			if isSimilar(car, otherCar) {
				counts[i]++
			}
		}
	}
	return counts
}

// helper function to check if two cars are similar
func isSimilar(car1, car2 string) bool {
	differenceCount := 0
	for i := 0; i < len(car1); i++ {
		if car1[i] != car2[i] {
			differenceCount++
		}
		if differenceCount > 1 {
			return false // more than one difference, not similar
		}
	}
	return true
}
