package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const noIndex = -1

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var sum int
	for scanner.Scan() {
		line := scanner.Text()
		firstDigit, lastDigit := 0, 0

		forwardReplacer := strings.NewReplacer(getReplacerArgs(false)...)
		forwardLine := forwardReplacer.Replace(line)
		backwardReplacer := strings.NewReplacer(getReplacerArgs(true)...)
		backwardLine := backwardReplacer.Replace(Reverse(line))
		re := regexp.MustCompile(`[^\d]*(\d?)`)
		firstDigit, _ = strconv.Atoi(re.FindStringSubmatch(forwardLine)[1])
		lastDigit, _ = strconv.Atoi(re.FindStringSubmatch(backwardLine)[1])
		sum += firstDigit*10 + lastDigit
	}

	fmt.Println("Result:")
	fmt.Println(strconv.Itoa(sum))
}

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

func getReplacerArgs(reverse bool) (result []string) {
	for idx, number := range getNumberToDigit() {
		if reverse {
			number = Reverse(number)
		}
		result = append(result, number, strconv.Itoa(idx+1))
	}
	return
}

// https://stackoverflow.com/a/4965535
func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
