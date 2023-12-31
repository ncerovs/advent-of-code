package challenge1

import (
	"fmt"
	"regexp"
	"strconv"
)

func GetCalibrationValues(lines []string) (int, error) {
	sum := 0
	for _, line := range lines {
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
