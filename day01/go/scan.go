package main

import (
	"bufio"
	"os"
)

func getSums(scanner *bufio.Scanner) (int, int) {
	pt1, pt2 := 0, 0
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			os.Exit(2)
		}

		txt := scanner.Text()

		n1 := partOne(txt)
		n2 := starTwo(txt)

		pt1 += n1
		pt2 += n2

	}
	return pt1, pt2
}
