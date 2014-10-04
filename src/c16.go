package main

import (
	"bytes"
	"cpals"
	"strings"
)

func enc(key []byte, user string) []byte {
	data := "comment1=cooking%20MCs;userdata=" + clean(user) + ";comment2=%20like%20a%20pound%20of%20bacon"

	padded := cpals.Pad7(16, []byte(data))
	//fmt.Printf("%s\n", data)

	return cpals.CBCEncrypt(key, cpals.Fill(16, "0"), padded)
}

func isAdmin(key, cookie []byte) bool {
	padded := cpals.CBCDecrypt(key, cpals.Fill(16, "0"), cookie)

	//fmt.Printf("%s\n", padded)
	if bytes.Index(padded, []byte(";admin=true;")) != -1 {
		return true
	}
	return false
}

func clean(s string) string {
	return strings.Replace(strings.Replace(s, ";", "", -1), "=", "", -1)
}

func main() {
	key := []byte("0123456789123456")

	// we need a block so we can flip the bits
	// not sure if there is a way to do this without knowing where :admin<true: is...
	str := "XXXXXXXXXXXXXXXX:admin<true:"
	//str := ":admin<true:"
	cookie := enc(key, str)

	// now flip the bit in the X block, when decoding it will also flip the bit
	// in the corresponding admin block, thanks CBC
	cookie[32] = cookie[32] ^ 1 // first :
	cookie[38] = cookie[38] ^ 1 // <
	cookie[43] = cookie[43] ^ 1 // :
	if isAdmin(key, cookie) == false {
		panic("failed to get admin")
	}

}
