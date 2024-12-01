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
	for i := range len(a) {
		d = a[i] - b[i]
		distance += max(d, -d)
	}

	fmt.Println("Distance:", distance)
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
