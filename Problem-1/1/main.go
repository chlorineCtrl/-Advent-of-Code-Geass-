package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {

	content, _ := os.Open("input")
	defer content.Close()

	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
		// fmt.Printf("line: %s\n", scanner.Text())
		line := scanner.Text()
		sum = findFirstNLastDigit(line)
	}
	fmt.Println(sum)

}

var sum int

func findFirstNLastDigit(inLine string) int {
	re := regexp.MustCompile(`\d`)
	matches := re.FindAllString(inLine, -1)
	// fmt.Println(matches)

	if len(matches) > 0 {
		first := matches[0]
		last := matches[len(matches)-1]
		result := first + last
		resultint, _ := strconv.Atoi(result)
		sum += resultint
	}
	return sum
}
