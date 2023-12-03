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
	// fmt.Println(line)
	hasNumber := true
	for hasNumber {
		firstIndexOfNumber := -1
		var keyOfFirstNumber int
		numberToDigit := getNumberToDigit()
		for key, word := range numberToDigit {
			index := strings.Index(line, word)
			if index == -1 {
				continue
			}
			if firstIndexOfNumber == -1 || index < firstIndexOfNumber {
				firstIndexOfNumber = index
				keyOfFirstNumber = key
			}
		}
		if firstIndexOfNumber != -1 {
			line = strings.ReplaceAll(line, numberToDigit[keyOfFirstNumber], strconv.Itoa(keyOfFirstNumber+1))
		} else {
			hasNumber = false
		}
	}

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
	// fmt.Println(line)
	// fmt.Println(strconv.Itoa(first))
	// fmt.Println(strconv.Itoa(last))
	// fmt.Println(strconv.Itoa(sum))
	return sum
}
