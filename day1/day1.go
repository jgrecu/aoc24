package main

import (
	"fmt"
	"sort"
	"strings"
)

func parseInput(input string) ([]int, []int) {
	// Split the lines by '\n' after we trim other spaces.
	lines := strings.Split(strings.TrimSpace(input), "\n")

	// Prepare the slices.
	leftList := make([]int, len(lines))
	rightList := make([]int, len(lines))

	// Parse the lines into left and right lists.
	for i, line := range lines {
		fmt.Sscanf(line, "%d %d", &leftList[i], &rightList[i])
	}

	return leftList, rightList
}

func calculateDistance(leftList, rightList []int) int {
	// To calculate the distance we need ordered lists first.
	sort.Ints(leftList)
	sort.Ints(rightList)

	totalDistance := 0

	// Calculate the sum of the diference between list values.
	for i := 0; i < len(leftList); i++ {
		totalDistance += abs(leftList[i] - rightList[i])
	}

	return totalDistance
}

// abs is a simple function to return the absolute of an int
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func calculateSimilarityScore(leftList, rightList []int) int {
	// Count occurrences of each number in the right list
	rightCounts := make(map[int]int)
	for _, num := range rightList {
		rightCounts[num]++
	}

	// Calculate similarity score
	similarityScore := 0
	for _, leftNum := range leftList {
		// Multiply left number by its count in the right list
		similarityScore += leftNum * rightCounts[leftNum]
	}

	return similarityScore
}

func main() {
	leftList, rightList := parseInput(input)
	distance := calculateDistance(leftList, rightList)
	score := calculateSimilarityScore(leftList, rightList)

	fmt.Println("Total distance:", distance)
	fmt.Println("Similarity score:", score)
}
