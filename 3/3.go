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
	sum := 0
	scanner.Scan()

	// Prefill for the first line
	nextLine := scanner.Text()
	var currentLine string

	for scanner.Scan() {
		currentLine, nextLine, sum = processLine(currentLine, nextLine, scanner.Text(), sum)
	}

	// Handle the last line
	_, _, sum = processLine(currentLine, nextLine, "", sum)

	fmt.Println("Sum of part numbers:")
	fmt.Println(sum)
}

func processLine(previousLine string, currentLine string, nextLine string, sum int) (string, string, int) {
	symbolSlices := findSymbolSlices(previousLine, currentLine, nextLine)

	for _, numberSlice := range findNumberSlices(currentLine) {
		if isNumberAdjacentToSymbol(numberSlice, symbolSlices) {
			number, _ := strconv.Atoi(currentLine[numberSlice[0]:numberSlice[1]])
			sum += number
		}
	}
	return currentLine, nextLine, sum
}

func findNumberSlices(currentLine string) [][]int {
	reNumber := regexp.MustCompile(`\d+`)
	numberSlices := reNumber.FindAllStringSubmatchIndex(currentLine, -1)

	return numberSlices
}

func findSymbolSlices(previousLine string, currentLine string, nextLine string) [][]int {
	reSymbol := regexp.MustCompile(`[^\d\.]+`)
	symbolSlices := reSymbol.FindAllStringSubmatchIndex(previousLine, -1)
	symbolSlices = append(symbolSlices, reSymbol.FindAllStringSubmatchIndex(currentLine, -1)...)
	symbolSlices = append(symbolSlices, reSymbol.FindAllStringSubmatchIndex(nextLine, -1)...)

	return symbolSlices
}

func isNumberAdjacentToSymbol(numberSlice []int, symbolSlices [][]int) bool {
	for _, symbolSlice := range symbolSlices {
		if symbolSlice[1] >= numberSlice[0] && symbolSlice[0] <= numberSlice[1] {
			return true
		}
	}
	return false
}
