package main

import (
	"bytes"
	"cpals"
	"fmt"
)

func main() {
	pad := cpals.Pad7(20, []byte("YELLOW SUBMARINE"))

	if !bytes.Equal(pad, []byte("YELLOW SUBMARINE\x04\x04\x04\x04")) {
		panic(fmt.Sprintf("failed to pad7: %x", pad))
	}
}
