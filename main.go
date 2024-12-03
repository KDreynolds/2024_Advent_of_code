package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	safeCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		
		// Parse numbers from line
		var numbers []int
		for _, numStr := range strings.Fields(line) {
			num, _ := strconv.Atoi(numStr)
			numbers = append(numbers, num)
		}
		
		if isSequenceSafeWithDampener(numbers) {
			safeCount++
		}
	}

	fmt.Println(safeCount)
}

func isSequenceSafeWithDampener(nums []int) bool {
	// First check if sequence is safe without removing any number
	if isSequenceSafe(nums) {
		return true
	}
	
	// Try removing each number one at a time
	for i := 0; i < len(nums); i++ {
		// Create new slice without number at index i
		dampened := make([]int, 0, len(nums)-1)
		dampened = append(dampened, nums[:i]...)
		dampened = append(dampened, nums[i+1:]...)
		
		if isSequenceSafe(dampened) {
			return true
		}
	}
	
	return false
}

func isSequenceSafe(nums []int) bool {
	if len(nums) < 2 {
		return true
	}

	// Check first difference to determine if sequence should increase or decrease
	increasing := nums[1] > nums[0]
	
	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		
		// Check if difference is between 1 and 3
		if abs(diff) < 1 || abs(diff) > 3 {
			return false
		}
		
		// Check if sequence maintains direction (increasing or decreasing)
		if increasing && diff < 0 || !increasing && diff > 0 {
			return false
		}
	}
	
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}