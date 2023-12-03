package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var sum int
	for scanner.Scan() {
		sum += calculateCalibrationValue(replaceNumberWordWithDigit(scanner.Text()))
	}
	fmt.Println(strconv.Itoa(sum))
}

func replaceNumberWordWithDigit(line string) string {
	numberWordToDigit := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	for word, digit := range numberWordToDigit {
		line = strings.ReplaceAll(line, word, digit)
	}
	fmt.Println(line)

	return line
}

func calculateCalibrationValue(line string) int {
	var first, last int
	digitFound := false
	for _, char := range line {
		if unicode.IsDigit(char) {
			last = int(char - '0') // Convert rune to int
			if !digitFound {
				digitFound = true
				first = int(char - '0') // Convert rune to int
			}
		}
	}

	sum := first*10 + last
	// fmt.Println("Line: " + line)
	// fmt.Println(strconv.Itoa(first))
	// fmt.Println(strconv.Itoa(last))
	// fmt.Println(strconv.Itoa(sum))
	return sum
}
