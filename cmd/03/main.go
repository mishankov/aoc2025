package main

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func first(banks [][]int) {
	sum := 0

	for _, bank := range banks {
		firstNum := 0
		firstNumIndex := 0
		secondNum := 0

		for index, num := range bank[:len(bank)-1] {
			if num > firstNum {
				firstNum = num
				firstNumIndex = index

				if num == 9 {
					break
				}
			}
		}

		for _, num := range bank[firstNumIndex+1:] {
			if num > secondNum {
				secondNum = num
			}

			if num == 9 {
				break
			}
		}

		sum += firstNum*10 + secondNum
	}

	log.Println("Answer 1:", sum)
}

func second(banks [][]int) {
	sum := 0

	for _, bank := range banks {
		nums := [12]int{}
		leftIndexWall := 0
		rightIndexWall := len(bank) - 11

		for i := range nums {
			bestNum := 0
			originalLeftIndexWall := leftIndexWall
			for index, num := range bank[leftIndexWall:rightIndexWall] {
				if num > bestNum {
					bestNum = num
					leftIndexWall = originalLeftIndexWall + index + 1
				}

				if num == 9 {
					break
				}
			}

			rightIndexWall++
			nums[i] = bestNum
		}

		for index, num := range nums {
			sum += num * int(math.Pow(10, float64(11-index)))
		}
	}

	log.Println("Answer 2:", sum)
}

func main() {
	log.Println("Day 3!")

	data, err := os.ReadFile("./data/03")
	if err != nil {
		log.Fatal(err)
	}

	banks := [][]int{}
	for sb := range strings.SplitSeq(string(data), "\n") {
		if sb == "" {
			continue
		}

		bank := []int{}
		for _, b := range sb {
			bn, err := strconv.Atoi(string(b))
			if err != nil {
				log.Fatal(err)
			}
			bank = append(bank, bn)
		}
		banks = append(banks, bank)
	}

	first(banks)
	second(banks)
}
