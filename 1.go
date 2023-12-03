package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func getNumberToDigit() [9]string {
	return [9]string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var sum int
	for scanner.Scan() {
		sum += calculateCalibrationValue(replaceNumberWordWithDigit(scanner.Text()))
	}
	fmt.Println(strconv.Itoa(sum))
}

func replaceNumberWordWithDigit(line string) string {
	for idx, word := range getNumberToDigit() {
		// fmt.Println(word, idx)
		line = strings.ReplaceAll(line, word, strconv.Itoa(idx+1))
	}
	// fmt.Println(line)
	// time.Sleep(time.Second)

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
	fmt.Println(line)
	// fmt.Println(strconv.Itoa(first))
	// fmt.Println(strconv.Itoa(last))
	// fmt.Println(strconv.Itoa(sum))
	return sum
}
