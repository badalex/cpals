package main

import (
	"bytes"
	"cpals"
)

func main() {
	s, err := cpals.Unpad7(16, []byte("ICE ICE BABY\x04\x04\x04\x04"))
	if err != nil {
		panic("failed")
	}

	if bytes.Equal([]byte("ICE ICE BABY"), s) == false {
		panic("failed")
	}

	_, err = cpals.Unpad7(16, []byte("ICE ICE BABY\x05\x05\x05\x05"))
	if err == nil {
		panic("no err")
	}

	_, err = cpals.Unpad7(16, []byte("ICE ICE BABY\x01\x02\x03\x04"))
	if err == nil {
		panic("no err")
	}

}
