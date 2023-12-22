package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var wordToDigit = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {

	content, _ := os.Open("input")
	defer content.Close()

	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
		fmt.Printf("line: %s\n", scanner.Text())
		line := scanner.Text()
		sum = findFirstNLastDigit(line)
	}
	fmt.Println(sum)

}

var sum int

func findFirstNLastDigit(inLine string) int {
	re := regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`)
	matches := re.FindAllString(inLine, -1)
	fmt.Println(matches)

	var numericValues []string
	for _, match := range matches {
		numericValue := match

		if val, ok := wordToDigit[match]; ok {
			numericValue = strconv.Itoa(val)
		}

		numericValues = append(numericValues, numericValue)
	}

	fmt.Println(numericValues)

	if len(numericValues) > 0 {
		first := numericValues[0]
		last := numericValues[len(numericValues)-1]
		result := first + last
		fmt.Println(result)
		resultint, _ := strconv.Atoi(result)
		sum += resultint
	}
	return sum
}
