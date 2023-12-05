package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	maxAllowedCounts := map[string]int{
		"blue":  14,
		"green": 13,
		"red":   12,
	}
	sumOfIds := 0
	totalPower := 0
	for scanner.Scan() {
		skipId := false
		line := scanner.Text()
		// Initialize each color with 1 as we're multiplying these values later
		maxCountsInGame := map[string]int{
			"blue":  1,
			"green": 1,
			"red":   1,
		}
		slices := strings.Split(line, ": ")
		id, _ := strconv.Atoi(strings.Split(slices[0], " ")[1])
		for _, subset := range strings.Split(slices[1], "; ") {
			for _, set := range strings.Split(subset, ", ") {
				setSlices := strings.Split(set, " ")
				count, _ := strconv.Atoi(setSlices[0])
				color := setSlices[1]
				if count > maxCountsInGame[color] {
					maxCountsInGame[color] = count
				}
			}
		}

		power := 1
		for color, count := range maxCountsInGame {
			power *= count

			if count > maxAllowedCounts[color] {
				skipId = true
			}
		}
		totalPower += power

		if !skipId {
			sumOfIds += id
		}
	}
	fmt.Println("Sum of IDs:")
	fmt.Println(sumOfIds)
	fmt.Println("Total power:")
	fmt.Println(totalPower)
}
