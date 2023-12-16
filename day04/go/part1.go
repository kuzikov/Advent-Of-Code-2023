package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func partOne(file *os.File) int {
	if _, err := file.Seek(0, 0); err != nil {
		log.Fatalln(err.Error())
	}
	sum := 0

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		txt := sc.Text()[10:]
		parts := strings.Split(txt, " | ")
		winPart := parts[0]
		ownPart := parts[1]

		winNumbers := strings.Fields(winPart)
		ownNumbers := strings.Fields(ownPart)

		match := make([]string, 0, len(ownNumbers))

		// search for a matches in both parts.
		for _, ownedNum := range ownNumbers {
			if contains(winNumbers, ownedNum) {
				match = append(match, ownedNum)
			}
		}

		if len(match) == 0 {
			continue
		}

		// fmt.Printf("%v\n", match)
		points := 1
		for i := 1; i < len(match); i++ {
			points = 2 * points
		}
		sum += points

	}

	return sum
}

func contains(list []string, target string) bool {
	for _, v := range list {
		if v == target {
			return true
		}
	}
	return false
}
