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
	sum := 0
lines:
	for scanner.Scan() {
		line := scanner.Text()
		slices := strings.Split(line, ": ")
		id, _ := strconv.Atoi(strings.Split(slices[0], " ")[1])
		for _, subset := range strings.Split(slices[1], "; ") {
			for _, set := range strings.Split(subset, ", ") {
				setSlices := strings.Split(set, " ")
				count, _ := strconv.Atoi(setSlices[0])
				color := setSlices[1]
				maxCounts := map[string]int{
					"blue":  14,
					"green": 13,
					"red":   12,
				}

				if count > maxCounts[color] {
					continue lines
				}
			}
		}
		sum += id
	}
	fmt.Println(sum)
}
