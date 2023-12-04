package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func partTwo(file *os.File) int {
	_, err := file.Seek(0, 0)
	if err != nil {
		os.Exit(4)
	}
	sc := bufio.NewScanner(file)
	var gameIdSum int
	for i := 1; sc.Scan(); i++ {
		txt := sc.Text()
		pref := fmt.Sprintf("Game %d: ", i)
		txt, found := strings.CutPrefix(txt, pref)
		if !found {
			os.Exit(2)
		}

		repl := strings.NewReplacer(" ", "", ";", ",", "red", "r", "green", "g", "blue", "b")
		txt = repl.Replace(txt)

		cubSets := strings.Split(txt, ",")

		var rgbSet [3]int
		for _, set := range cubSets {
			color := set[len(set)-1:]
			amount, err := strconv.Atoi(set[:len(set)-1])
			if err != nil {
				os.Exit(3)
			}

			switch {
			case color == "r" && amount > rgbSet[0]:
				rgbSet[0] = amount
			case color == "g" && amount > rgbSet[1]:
				rgbSet[1] = amount
			case color == "b" && amount > rgbSet[2]:
				rgbSet[2] = amount
			}
		}

		gameIdSum += rgbSet[0] * rgbSet[1] * rgbSet[2]
	}

	return gameIdSum
}
