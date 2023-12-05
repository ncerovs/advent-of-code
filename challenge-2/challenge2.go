package challenge2

import (
	"fmt"
	"regexp"
	"strconv"
)

var digits = map[string]string{
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

func Recalibrate(lines []string) (int, error) {
	sum := 0
	for i, line := range lines {
		for word, number := range digits {
			line = regexp.MustCompile(word).ReplaceAllLiteralString(line, word+number+word)
		}

		nums := regexp.MustCompile("[^0-9]").ReplaceAllString(line, "")

		runes := []rune(nums)

		firstSubstring := string(runes[:1])

		lastSubstring := firstSubstring
		if len(runes) > 1 {
			lastSubstring = string(runes[len(runes)-1:])
		}

		combinedNum, _ := strconv.Atoi(fmt.Sprintf("%s%s", firstSubstring, lastSubstring))
		sum += combinedNum
	}

	return sum, nil
}
