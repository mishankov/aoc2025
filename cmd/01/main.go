package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func first(rotations []string) {
	position := 50
	nullCount := 0
	for _, rotation := range rotations {
		if len(rotation) == 0 {
			continue
		}

		rotDir := 0
		switch rotation[0] {
		case 'L':
			rotDir = -1
		case 'R':
			rotDir = 1
		}

		amount, err := strconv.Atoi(rotation[1:])
		if err != nil {
			log.Fatal(err)
		}

		position = position + rotDir*amount
		if position > 99 {
			position = position % 100
		}

		if position < 0 {
			position = position % 100
			if position < 0 {
				position += 100
			}
		}

		if position == 0 {
			nullCount++
		}
	}

	log.Println("Answer 1:", nullCount)
}

func second(rotations []string) {
	position := 50
	nullCount := 0
	for _, rotation := range rotations {
		if len(rotation) == 0 {
			continue
		}

		rotDir := 0
		switch rotation[0] {
		case 'L':
			rotDir = -1
		case 'R':
			rotDir = 1
		}

		amount, err := strconv.Atoi(rotation[1:])
		if err != nil {
			log.Fatal(err)
		}

		if rotDir > 0 {
			nullCount += (position + rotDir*amount) / 100
		}

		if rotDir < 0 {
			p := 0
			if position != 0 {
				p = position - 100
			}
			nullCount += -(p + rotDir*amount) / 100
		}

		position = position + rotDir*amount
		if position > 99 {
			position = position % 100
		}

		if position < 0 {
			position = position % 100
			if position < 0 {
				position += 100
			}
		}
	}

	log.Println("Answer 2:", nullCount)
}

func main() {
	log.Println("Day 1!")

	data, err := os.ReadFile("./data/01")
	if err != nil {
		log.Fatal(err)
	}

	rotations := strings.Split(string(data), "\n")

	first(rotations)
	second(rotations)
}
