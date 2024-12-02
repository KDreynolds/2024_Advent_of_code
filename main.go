package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read input from stdin
	scanner := bufio.NewScanner(os.Stdin)
	var leftList, rightList []int

	// Parse input into two lists
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		
		numbers := strings.Fields(line)
		if len(numbers) != 2 {
			continue
		}
		
		left, _ := strconv.Atoi(numbers[0])
		right, _ := strconv.Atoi(numbers[1])
		
		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	// Create a map to count occurrences in rightList
	rightCounts := make(map[int]int)
	for _, num := range rightList {
		rightCounts[num]++
	}

	// Calculate similarity score
	similarityScore := 0
	for _, leftNum := range leftList {
		count := rightCounts[leftNum]
		similarityScore += leftNum * count
	}

	fmt.Println(similarityScore)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}