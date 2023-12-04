package main

import (
	"fmt"
	test "ncerovs/advent-of-code/challenge-1"
	"os"
)

func main() {
	c, err := test.GetCalibrationValues()
	if err != nil {
		fmt.Printf("error getting calibration values: %w", err)
		os.Exit(1)
	}

	fmt.Println(c)
}
