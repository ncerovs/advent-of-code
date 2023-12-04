package test

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func GetCalibrationValues() (int, error) {
	apiURL := "https://adventofcode.com/2023/day/1/input"

	sessionString := "53616c7465645f5feaa5ad72951c8b6a5127dd46c94b625c05b2796e83e75c0c80699ea2265895f2fc5991d5bc4a49944d4acf3224a42d0dc87ba5be4fe01048"

	client := &http.Client{}

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return 0, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Cookie", "session="+sessionString)

	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("error reading body: %w", err)
	}

	lines := strings.Split(string(body), "\n")

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
