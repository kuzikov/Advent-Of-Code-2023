package main

import (
	"os"
	"strconv"
	"strings"
	"unicode"
)

func starTwo(str string) int {

	repl := strings.NewReplacer("one", "o1e", "two", "t2o", "three", "th3ee", "four", "fo4r", "five", "fi5ve", "six", "s6x", "seven", "se7en", "eight", "ei8ht", "nine", "ni9e")
	txt := repl.Replace(str)

	repl2 := strings.NewReplacer("one", "1", "two", "2", "three", "3", "four", "4", "five", "5", "six", "6", "seven", "7", "eight", "8", "nine", "9")
	txt = repl2.Replace(txt)

	var acc []rune
	for _, v := range txt {
		if unicode.IsDigit(v) {
			acc = append(acc, v)
		}
	}
	n, err := strconv.Atoi(string(acc[0]) + string(acc[len(acc)-1]))
	if err != nil {
		os.Exit(3)
	}
	return n
}
