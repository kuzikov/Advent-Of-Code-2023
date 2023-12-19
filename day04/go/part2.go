package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func partTwo(file *os.File) int {
	if _, err := file.Seek(0, 0); err != nil {
		log.Fatalln(err.Error())
	}
	cardCount := make([]int, 202)
	fill(cardCount, 1)

	sc := bufio.NewScanner(file)
	for i := 0; sc.Scan(); i++ {
		count := len(findMatches(sc.Text()))
		// fmt.Printf("Line: %d, match: %v, incStart: %d, incStop: %d\n", i, match, i+1, i+1+count)
		for j := i + 1; j < i+1+count; j++ {
			cardCount[j] += cardCount[i]
		}
	}

	return sumElements(cardCount)
}

func sumElements(cardStack []int) int {
	sum := 0
	for _, v := range cardStack {
		sum += v
	}
	return sum
}

func findMatches(card string) []string {
	txt := card[10:]
	parts := strings.Split(txt, " | ")
	winSide := parts[0]
	ownSide := parts[1]

	winNumbers := strings.Fields(winSide)
	ownNumbers := strings.Fields(ownSide)

	match := make([]string, 0, len(ownNumbers))

	for _, ownedNum := range ownNumbers {
		if contains(winNumbers, ownedNum) {
			match = append(match, ownedNum)
		}
	}
	return match
}

func fill(ar []int, n int) {
	for i := 0; i < len(ar); i++ {
		ar[i] = n
	}
}
