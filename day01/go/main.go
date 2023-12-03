package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./adventofcode.com_2023_day_1_input.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	sc := bufio.NewScanner(file)

	n1, n2 := getSums(sc)

	fmt.Printf("\n\npart1: %d\npart2: %d\n", n1, n2)
}
