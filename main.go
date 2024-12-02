package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	// Sort both lists
	sort.Ints(leftList)
	sort.Ints(rightList)

	// Calculate total distance
	totalDistance := 0
	for i := 0; i < len(leftList); i++ {
		distance := abs(leftList[i] - rightList[i])
		totalDistance += distance
	}

	fmt.Println(totalDistance)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}