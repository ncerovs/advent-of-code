package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	apiURL := "https://adventofcode.com/2023/day/1/input"

	sessionString := "<SESSION>"

	client := &http.Client{}

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(1)
	}

	req.Header.Set("Cookie", "session="+sessionString)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error making request:", err)
		os.Exit(1)
	}

	lines := strings.Split(string(body), "\n")

	sum := 0
	for _, line := range lines {
		fmt.Println(line)
		nums := regexp.MustCompile("[^0-9]").ReplaceAllString(line, "")

		runes := []rune(nums)

		firstSubstring := string(runes[:1])

		if err != nil {
			fmt.Println(err)
		}

		lastSubstring := firstSubstring
		if len(runes) > 1 {
			lastSubstring = string(runes[len(runes)-1:])
		}

		combined := fmt.Sprintf("%s%s", firstSubstring, lastSubstring)
		combinedNum, _ := strconv.Atoi(combined)
		sum += combinedNum
	}

	fmt.Println(sum)
}
