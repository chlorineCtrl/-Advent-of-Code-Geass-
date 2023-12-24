package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var NumberMap = map[string]int{
	"zero":  0,
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

	var ans int = 0
	content, _ := os.Open("input")
	defer content.Close()

	scanner := bufio.NewScanner(content)

	for scanner.Scan() {
		var line string = scanner.Text()
		// fmt.Println(line)
		ans += summer(line)
	}
	fmt.Println(ans)
}

func summer(line string) int {
	// summer wind
	var sum int = 0
	var firstNum int
	var lastNum int
	var isSet bool = false

	for i := 0; i < len(line); i++ {
		if line[i] >= '0' && line[i] <= '9' {
			var num int = int(line[i] - '0')
			// fmt.Printf("%d", num)
			if isSet == false {
				firstNum = num
				isSet = true
			}
			lastNum = num

		} else {
			for word, num := range NumberMap {
				if wordCheck(line, i, word) {
					firstNum = num
					isSet = true
				}
				lastNum = num

			}
		}
		sum += (firstNum * 10) + lastNum
	}
	return sum
}

func wordCheck(lineString string, index int, word string) bool {
	return strings.HasPrefix(lineString[index:], word)
}
