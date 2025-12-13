package main

import (
	"log"
	"os"
	"strings"
)

func first(grid [][]string) {
	sum := 0
	maxRow := len(grid) - 1
	maxCol := len(grid[0]) - 1

	for i, row := range grid {
		for j, cell := range row {
			if cell != "@" {
				continue
			}

			count := 0
			// top left
			if i != 0 && j != 0 && grid[i-1][j-1] == "@" {
				count++
			}

			// top
			if i != 0 && grid[i-1][j] == "@" {
				count++
			}

			// top right
			if i != 0 && j != maxCol && grid[i-1][j+1] == "@" {
				count++
			}

			// right
			if j != maxCol && grid[i][j+1] == "@" {
				count++
			}

			// bottom right
			if i != maxRow && j != maxCol && grid[i+1][j+1] == "@" {
				count++
			}

			// bottom
			if i != maxRow && grid[i+1][j] == "@" {
				count++
			}

			// bottom left
			if i != maxRow && j != 0 && grid[i+1][j-1] == "@" {
				count++
			}

			// left
			if j != 0 && grid[i][j-1] == "@" {
				count++
			}

			if count < 4 {
				sum++
			}
		}
	}

	log.Println("Answer 1:", sum)
}

func second(grid [][]string) {
	sum := 0

	maxRow := len(grid) - 1
	maxCol := len(grid[0]) - 1

	for {
		removed := false
		for i, row := range grid {
			for j, cell := range row {
				if cell != "@" {
					continue
				}

				count := 0
				// top left
				if i != 0 && j != 0 && grid[i-1][j-1] == "@" {
					count++
				}

				// top
				if i != 0 && grid[i-1][j] == "@" {
					count++
				}

				// top right
				if i != 0 && j != maxCol && grid[i-1][j+1] == "@" {
					count++
				}

				// right
				if j != maxCol && grid[i][j+1] == "@" {
					count++
				}

				// bottom right
				if i != maxRow && j != maxCol && grid[i+1][j+1] == "@" {
					count++
				}

				// bottom
				if i != maxRow && grid[i+1][j] == "@" {
					count++
				}

				// bottom left
				if i != maxRow && j != 0 && grid[i+1][j-1] == "@" {
					count++
				}

				// left
				if j != 0 && grid[i][j-1] == "@" {
					count++
				}

				if count < 4 {
					sum++
					grid[i][j] = "."
					removed = true
				}
			}
		}

		if !removed {
			break
		}
	}

	log.Println("Answer 2:", sum)
}

func main() {
	log.Println("Day 4!")

	data, err := os.ReadFile("./data/04")
	if err != nil {
		log.Fatal(err)
	}

	grid := [][]string{}
	for sr := range strings.SplitSeq(string(data), "\n") {
		if sr == "" {
			continue
		}

		row := strings.Split(sr, "")
		grid = append(grid, row)
	}

	first(grid)
	second(grid)
}
