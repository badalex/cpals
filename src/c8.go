package main

import (
	"bufio"
	"cpals"
	"os"
)

func main() {
	fh, err := os.Open("data/8.txt")

	scanner := bufio.NewScanner(fh)
	lineno, err := cpals.DetectECBScanner(16, scanner)
	if err != nil {
		panic(err)
	}
	fh.Close()

	if lineno != 133 {
		panic("cpals")
	}
}
