package main

import (
	"bytes"
	"cpals"
	"fmt"
)

func main() {
	ct := cpals.B64ReadFile("data/7.txt")
	pt := cpals.ECBDecrypt([]byte("YELLOW SUBMARINE"), ct)

	if bytes.Contains(pt, []byte("Play that funky music, white boy Come on, Come on, Come on")) == false {
		fmt.Printf("%s", pt)
		panic("failed")
	}
}
