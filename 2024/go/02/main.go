package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func isSafe(nums []int) bool {
	isIncreasing := nums[0] < nums[1]
	isDecreasing := nums[0] > nums[1]

	var diff float64
	for i := 1; i < len(nums); i++ {
		if (nums[i] > nums[i-1] && isDecreasing) || (nums[i] < nums[i-1] && isIncreasing) {
			return false
		}

		diff = math.Abs(float64(nums[i] - nums[i-1]))
		if diff > 3 || diff < 1 {
			return false
		}
	}

	return true
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	safeReports := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		var nums []int

		for _, n := range strings.Split(line, " ") {
			nint, err := strconv.Atoi(n)
			if err != nil {
				log.Fatal(err)
			}
			nums = append(nums, nint)
		}

		if isSafe(nums) {
			safeReports++
		}
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("safe reports: %d\n", safeReports)
}
