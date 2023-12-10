package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var sumOfPartNumbers, sumOfGearRatios int
	scanner.Scan()

	// Prefill for the first line
	nextLine := scanner.Text()
	var currentLine string

	for scanner.Scan() {
		currentLine, nextLine, sumOfPartNumbers, sumOfGearRatios = processLine(currentLine, nextLine, scanner.Text(), sumOfPartNumbers, sumOfGearRatios)
	}

	// Handle the last line
	_, _, sumOfPartNumbers, sumOfGearRatios = processLine(currentLine, nextLine, "", sumOfPartNumbers, sumOfGearRatios)

	fmt.Println("Sum of part numbers:")
	fmt.Println(sumOfPartNumbers)
	fmt.Println("Sum of gear ratios:")
	fmt.Println(sumOfGearRatios)
}

func processLine(previousLine string, currentLine string, nextLine string, sumOfPartNumbers int, sumOfGearRatios int) (string, string, int, int) {
	// Sum up part numbers
	symbolSlices := findSlices(`[^\d\.]+`, previousLine, currentLine, nextLine)

	for _, numberSlice := range findSlices(`\d+`, currentLine) {
		if isNumberAdjacentToSymbol(numberSlice, symbolSlices) {
			number, _ := strconv.Atoi(currentLine[numberSlice[0]:numberSlice[1]])
			sumOfPartNumbers += number
		}
	}

	// Sum up gear ratios

	return currentLine, nextLine, sumOfPartNumbers, sumOfGearRatios
}

func findSlices(regex string, lines ...string) (slices [][]int) {
	re := regexp.MustCompile(regex)
	for _, line := range lines {
		slices = append(slices, re.FindAllStringSubmatchIndex(line, -1)...)
	}

	return slices
}

func isNumberAdjacentToSymbol(numberSlice []int, symbolSlices [][]int) bool {
	for _, symbolSlice := range symbolSlices {
		if symbolSlice[1] >= numberSlice[0] && symbolSlice[0] <= numberSlice[1] {
			return true
		}
	}
	return false
}
