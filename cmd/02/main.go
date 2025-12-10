package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func first(ranges [][]int) {
	sum := 0
	for _, r := range ranges {
		for i := r[0]; i <= r[1]; i++ {
			s := strconv.Itoa(i)

			if len(s)%2 != 0 {
				continue
			}

			if s[0:len(s)/2] == s[len(s)/2:] {
				sum += i
			}
		}
	}

	log.Println("Answer 1:", sum)
}

func second(ranges [][]int) {
	sum := 0
	for _, r := range ranges {
		for num := r[0]; num <= r[1]; num++ {
			s := strconv.Itoa(num)
			for i := range len(s) / 2 {
				c := strings.Count(s, s[0:i+1])
				if c*(i+1) == len(s) {
					sum += num
					break
				}
			}
		}
	}

	log.Println("Answer 2:", sum)
}

func main() {
	log.Println("Day 2!")

	data, err := os.ReadFile("./data/02")
	if err != nil {
		log.Fatal(err)
	}

	ranges := [][]int{}
	for s := range strings.SplitSeq(string(data), ",") {
		sr := strings.Split(s, "-")
		n1, err := strconv.Atoi(strings.TrimSpace(sr[0]))
		if err != nil {
			log.Fatal(err)
		}
		n2, err := strconv.Atoi(strings.TrimSpace(sr[1]))
		if err != nil {
			log.Fatal(err)
		}

		ranges = append(ranges, []int{n1, n2})
	}

	first(ranges)
	second(ranges)
}
