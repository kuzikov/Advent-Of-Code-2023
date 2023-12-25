package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("../adventofcode.com_2023_day_5_input.txt")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer file.Close()

	fmt.Println(partOne(file))
}
