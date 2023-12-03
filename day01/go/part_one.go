package main

import (
	"os"
	"strconv"
	"strings"
)

func partOne(str string) int {
	txt := strings.Map(func(r rune) rune {
		if r > '9' || r < '0' {
			return -1
		}
		return r
	}, str)

	n, err := strconv.Atoi(string([]byte{txt[0], txt[len(txt)-1]}))
	if err != nil {
		os.Exit(3)
	}
	return n
}
