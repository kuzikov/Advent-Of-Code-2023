package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("adventofcode.com_2023_day_2_input.txt")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	fmt.Println(partOne(file))
	fmt.Println(partTwo(file))

}
