package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	left, right, err := getInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sort.Ints(left)
	sort.Ints(right)

	distance := calculateDistance(left, right)
	fmt.Println("Distance:", distance)

	similarityScore := calculateSimilarity(left, right)
	fmt.Println("Similarity score:", similarityScore)
}

func calculateDistance(left, right []int) int {
	var distance, tmpd int

	for i := range left {
		tmpd = left[i] - right[i]
		distance += max(tmpd, -tmpd)
	}

	return distance
}

func calculateSimilarity(left, right []int) int {
	counter := make(map[int]int, 0)
	for _, v := range right {
		counter[v]++
	}

	similarityScore := 0
	for _, v := range left {
		similarityScore += v * counter[v]
	}

	return similarityScore
}

func getInput(filename string) ([]int, []int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	var leftList, rightList []int
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		var left, right int
		_, err := fmt.Sscan(line, &left, &right)
		if err != nil {
			return nil, nil, err
		}

		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	if err = scanner.Err(); err != nil {
		return nil, nil, err
	}

	return leftList, rightList, nil
}
