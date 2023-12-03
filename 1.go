package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var sum int
	for scanner.Scan() {
		sum += calculateCalibrationValue(scanner.Text())
	}
	fmt.Println(strconv.Itoa(sum))
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
