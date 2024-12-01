package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	// every line in the input is 14 bytes long
	const lineLen = 14

	f, err := os.Open("input.txt")
	check(err)

	buffer := make([]byte, lineLen)
	a := make([]int, 0)
	b := make([]int, 0)

	for {
		_, err = f.Read(buffer)

		if err == io.EOF {
			break
		}

		check(err)

		line := string(buffer)
		left, err := strconv.Atoi(line[:5])
		check(err)
		right, err := strconv.Atoi(line[8:13])
		check(err)

		a = append(a, left)
		b = append(b, right)
	}

	if len(a) != len(b) {
		log.Fatal("input error")
	}

	sort.Ints(a)
	sort.Ints(b)

	distance := 0
	var d int
	for i := range a {
		d = a[i] - b[i]
		distance += max(d, -d)
	}

	// part 1 solution
	fmt.Println("Distance:", distance)

	counter := make(map[int]int, 0)
	for _, v := range b {
		counter[v]++
	}

	similarityScore := 0
	for _, v := range a {
		similarityScore += v * counter[v]
	}

	// part 2 solution
	fmt.Println("Similarity score:", similarityScore)
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
