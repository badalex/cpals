package main

import (
	"bytes"
	"cpals"
)

func main() {
	iv := cpals.Fill(16, "\x00")

	//pt := []byte("01234567891234560123456789123456")
	//enc := CBCEncrypt([]byte("YELLOW SUBMARINE"), iv, pt)
	////fmt.Printf("%x %d\n", enc, len(enc))
	////dec := ECBDecrypt([]byte("YELLOW SUBMARINE"), enc)
	//cbcdec := CBCDecrypt([]byte("YELLOW SUBMARINE"), iv, enc)
	//fmt.Printf("P: %x\n", pt)
	//fmt.Printf("E: %x\n", enc)
	//fmt.Printf("F: %x %s %d\n", cbcdec, cbcdec, len(cbcdec))

	ct := cpals.B64ReadFile("data/10.txt")
	pt := cpals.CBCDecrypt([]byte("YELLOW SUBMARINE"), iv, ct)

	if bytes.Contains(pt, []byte("Play that funky music, white boy Come on, Come on, Come on")) == false {
		panic("cpals")
	}
}
