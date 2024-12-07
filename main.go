package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Read input into a 2D grid
	var grid []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	
	rows := len(grid)
	if rows == 0 {
		return
	}
	cols := len(grid[0])
	
	count := 0
	
	// Check each possible center point
	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			if grid[i][j] == 'A' {
				// Check all possible combinations of MAS in X pattern
				// Upper-left to lower-right and upper-right to lower-left
				if isValidXMAS(grid, i, j, rows, cols) {
					count++
				}
			}
		}
	}
	
	fmt.Println(count)
}

func isValidXMAS(grid []string, i, j, rows, cols int) bool {
	// Check if we can form valid MAS strings in both diagonals
	ul := string([]byte{grid[i-1][j-1], grid[i][j], grid[i+1][j+1]})  // upper-left to lower-right
	ur := string([]byte{grid[i-1][j+1], grid[i][j], grid[i+1][j-1]})  // upper-right to lower-left
	
	return isValidArm(ul) && isValidArm(ur)
}

func isValidArm(s string) bool {
	return s == "MAS" || s == "SAM"
}