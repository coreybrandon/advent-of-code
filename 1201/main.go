// Sonar Sweep

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func IncreasedDepth(report []int) int {
	count := 0

	// Iterate through measurement reports
	for i := 1; i < len(report); i++ {
		// Compare the current element to the last
		if report[i] > report[i-1] {
			count++
		}
	}
	return count
}

func PartTwo(report []int) int {
	count := 0

	// Iterate through sliding windows
	for i := 3; i < len(report); i++ {
		// Sum the elements in the first and second 'three-measurement sliding window'
		firstSum := report[i-1] + report[i-2] + report[i-3]
		secondSum := report[i] + report[i-1] + report[i-2]
		if secondSum > firstSum {
			count++
		}
	}
	return count
}

func main() {
	file, err := os.Open("puzzle_piece")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	var depth []int

	// Iterate through lines, convert + append to 'depth'
	for scanner.Scan() {
		line := scanner.Text()
		if i, err := strconv.Atoi(line); err == nil {
			depth = append(depth, i)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(IncreasedDepth(depth))
	fmt.Println(PartTwo(depth))
}
