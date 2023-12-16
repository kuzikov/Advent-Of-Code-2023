package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("../adventofcode.com_2023_day_4_input.txt")
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Printf("part 1: %d\n", partOne(file))
}
