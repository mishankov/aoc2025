package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func first(ranges [][2]int, ids []int) {
	sum := 0

	for _, id := range ids {
		for _, r := range ranges {
			if id >= r[0] && id <= r[1] {
				sum++
				break
			}
		}
	}

	log.Println("Answer 1:", sum)
}

func second(ranges [][2]int) {
	sum := 0

	for {
		merged := false
		newRanges := [][2]int{}

		for _, r := range ranges {
			processed := false
			for i, nr := range newRanges {
				// r - [], nr - {}
				// {[}]
				if r[0] >= nr[0] && r[0] <= nr[1] && r[1] > nr[1] {
					nr[1] = r[1]
					newRanges[i] = nr
					merged = true
					processed = true
					break
				}

				// [{]}
				if r[0] < nr[0] && r[1] >= nr[0] && r[1] <= nr[1] {
					nr[0] = r[0]
					newRanges[i] = nr
					merged = true
					processed = true
					break
				}

				// {[]}
				if r[0] >= nr[0] && r[0] <= nr[1] && r[1] >= nr[0] && r[1] <= nr[1] {
					merged = true
					processed = true
					break
				}

				// [{}]
				if r[0] <= nr[0] && r[1] >= nr[1] {
					newRanges[i] = r
					merged = true
					processed = true
					break
				}
			}

			if !processed {
				newRanges = append(newRanges, r)
			}

		}
		ranges = newRanges

		if !merged {
			break
		}
	}

	for _, r := range ranges {
		sum += r[1] - r[0] + 1
	}

	log.Println("Answer 2:", sum)
}

func main() {
	log.Println("Day 5!")

	data, err := os.ReadFile("./data/05")
	if err != nil {
		log.Fatal(err)
	}

	rangeData := strings.Split(string(data), "\n\n")[0]
	idData := strings.Split(string(data), "\n\n")[1]

	ranges := [][2]int{}
	for r := range strings.SplitSeq(rangeData, "\n") {
		if r == "" {
			continue
		}
		r := strings.TrimSpace(r)

		leftNum, err := strconv.Atoi(strings.Split(r, "-")[0])
		if err != nil {
			log.Fatal(err)
		}

		rightNum, err := strconv.Atoi(strings.Split(r, "-")[1])
		if err != nil {
			log.Fatal(err)
		}

		ranges = append(ranges, [2]int{leftNum, rightNum})
	}

	ids := []int{}
	for r := range strings.SplitSeq(idData, "\n") {
		if r == "" {
			continue
		}
		r := strings.TrimSpace(r)

		id, err := strconv.Atoi(r)
		if err != nil {
			log.Fatal(err)
		}

		ids = append(ids, id)
	}

	first(ranges, ids)
	second(ranges)
}
