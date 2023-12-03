package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getNumberToDigit() []string {
	return []string{
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

const noIndex = -1

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var sum int
	for scanner.Scan() {
		line := scanner.Text()
		firstDigit, lastDigit := 0, 0
		indexOfFirstDigit, indexOfLastDigit := getIndicesOfDigits(line)
		// "word" meaning numbers like "one", "two" etc. up to "nine"
		sliceOfFirstWord := getSliceOfFirstWord(line)
		sliceOfLastWord := getSliceOfLastWord(line)

		// fmt.Println(line)
		// fmt.Println("indexOfFirstDigit: " + strconv.Itoa(indexOfFirstDigit))
		// fmt.Println("indexOfLastDigit: " + strconv.Itoa(indexOfLastDigit))
		// fmt.Printf("indexOfFirstWord: %v\n", sliceOfFirstWord)
		// fmt.Printf("indexOfLastWord: %v\n", sliceOfLastWord)

		if indexOfFirstDigit == noIndex && sliceOfFirstWord[0] == noIndex {
			continue
		}

		if indexOfFirstDigit == noIndex || (sliceOfFirstWord[0] != noIndex && sliceOfFirstWord[0] < indexOfFirstDigit) {
			firstDigit = getDigitByWordIndex(line, sliceOfFirstWord)
		} else if sliceOfFirstWord[0] == noIndex || (indexOfFirstDigit != noIndex && indexOfFirstDigit < sliceOfFirstWord[0]) {
			firstDigit, _ = strconv.Atoi(string(line[indexOfFirstDigit]))
		}

		if indexOfLastDigit == noIndex || (sliceOfLastWord[0] != noIndex && sliceOfLastWord[0] > indexOfLastDigit) {
			lastDigit = getDigitByWordIndex(line, sliceOfLastWord)
		} else if sliceOfLastWord[0] == noIndex || (indexOfLastDigit != noIndex && indexOfLastDigit > sliceOfLastWord[0]) {
			lastDigit, _ = strconv.Atoi(string(line[indexOfLastDigit]))
		}
		// fmt.Println(firstDigit*10 + lastDigit)

		sum += firstDigit*10 + lastDigit
	}

	fmt.Println("Result:")
	fmt.Println(strconv.Itoa(sum))
}

func getDigitByWordIndex(line string, slice []int) int {
	numberFromLine := line[slice[0]:slice[1]]
	for idx, number := range getNumberToDigit() {
		if numberFromLine == number {
			return idx + 1
		}
	}
	return 0
}

func getIndicesOfDigits(line string) (int, int) {
	firstIndex, lastIndex := noIndex, noIndex
	re := regexp.MustCompile(`[[:digit:]]`)
	matches := re.FindAllStringIndex(line, -1)
	if len(matches) == 0 {
		return firstIndex, lastIndex
	}
	firstIndex = matches[0][0]
	lastIndex = matches[len(matches)-1][0]

	return firstIndex, lastIndex
}

func getSliceOfFirstWord(line string) []int {
	re := regexp.MustCompile(strings.Join(getNumberToDigit(), "|"))
	matches := re.FindStringIndex(line)
	if len(matches) == 0 {
		return []int{noIndex, noIndex}
	}
	return matches
}

func getSliceOfLastWord(line string) []int {
	re := regexp.MustCompile(Reverse(strings.Join(getNumberToDigit(), "|")))
	matches := re.FindStringIndex(Reverse(line))
	if len(matches) == 0 {
		return []int{noIndex, noIndex}
	}
	return []int{
		(len(line) - 1) - (matches[1] - 1),
		(len(line) - 1) - matches[0] + 1,
	}
}

// https://stackoverflow.com/a/4965535
func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
