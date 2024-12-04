package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	enabled := true // Start with multiplications enabled
	
	// Regex patterns
	mulRe := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	doRe := regexp.MustCompile(`do\(\)`)
	dontRe := regexp.MustCompile(`don't\(\)`)
	
	for scanner.Scan() {
		line := scanner.Text()
		pos := 0
		
		for pos < len(line) {
			// Find next occurrence of each pattern
			mulMatch := mulRe.FindStringSubmatch(line[pos:])
			mulIndex := mulRe.FindStringIndex(line[pos:])
			doMatch := doRe.FindStringIndex(line[pos:])
			dontMatch := dontRe.FindStringIndex(line[pos:])
			
			// Adjust matches to account for current position
			if mulIndex != nil {
				mulIndex[0] += pos
				mulIndex[1] += pos
			}
			if doMatch != nil {
				doMatch[0] += pos
				doMatch[1] += pos
			}
			if dontMatch != nil {
				dontMatch[0] += pos
				dontMatch[1] += pos
			}
			
			// Find the earliest match
			nextPos := len(line)
			if mulIndex != nil {
				nextPos = mulIndex[0]
			}
			if doMatch != nil && doMatch[0] < nextPos {
				nextPos = doMatch[0]
			}
			if dontMatch != nil && dontMatch[0] < nextPos {
				nextPos = dontMatch[0]
			}
			
			// Process the earliest match
			if nextPos == len(line) {
				break
			}
			
			if doMatch != nil && doMatch[0] == nextPos {
				enabled = true
				pos = doMatch[1]
			} else if dontMatch != nil && dontMatch[0] == nextPos {
				enabled = false
				pos = dontMatch[1]
			} else if mulIndex != nil && mulIndex[0] == nextPos && mulMatch != nil {
				if enabled {
					x, _ := strconv.Atoi(mulMatch[1])
					y, _ := strconv.Atoi(mulMatch[2])
					sum += x * y
				}
				pos = mulIndex[1]
			}
		}
	}
	
	fmt.Println(sum)
}