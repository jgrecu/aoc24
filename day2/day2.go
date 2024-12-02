package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseReports(input string) [][]int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	reports := make([][]int, len(lines))

	for i, line := range lines {
		// Split the line and convert each number to an integer
		numStrings := strings.Fields(line)
		reports[i] = make([]int, len(numStrings))

		for j, numStr := range numStrings {
			num, _ := strconv.Atoi(numStr)
			reports[i][j] = num
		}
	}

	return reports
}

func isReportSafe(report []int) bool {
	// Check if levels are increasing or decreasing
	isIncreasing := true
	isDecreasing := true

	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]

		// Check if difference is within ±1 to ±3
		if abs(diff) < 1 || abs(diff) > 3 {
			return false
		}

		// Update increasing/decreasing status
		if diff > 0 {
			isDecreasing = false
		} else if diff < 0 {
			isIncreasing = false
		} else {
			// If difference is exactly 0, it's not considered increasing or decreasing
			return false
		}
	}

	// Must be entirely increasing or entirely decreasing
	return isIncreasing || isDecreasing
}

func isReportSafeWithDampener(report []int) bool {
	// First, check if the report is already safe
	if isReportSafe(report) {
		return true
	}

	// Try removing each level and check for safety
	for i := 0; i < len(report); i++ {
		// Create a new slice without the i-th element
		modifiedReport := make([]int, 0, len(report)-1)
		modifiedReport = append(modifiedReport, report[:i]...)
		modifiedReport = append(modifiedReport, report[i+1:]...)

		// Check if the modified report is safe
		if isReportSafe(modifiedReport) {
			return true
		}
	}

	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func countSafeReports(reports [][]int) int {
	safeReportCount := 0
	for _, report := range reports {
		if isReportSafeWithDampener(report) {
			safeReportCount++
		}
	}
	return safeReportCount
}

func main() {
	reports := parseReports(input)
	safeReports := countSafeReports(reports)
	fmt.Printf("Number of safe reports: %d\n", safeReports)
}
