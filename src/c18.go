package main

import (
	"bytes"
	"cpals"
	"fmt"
)

func main() {
	//key := []byte("0123456789123456")
	//out := cpals.CTREncrypt(key, []byte("ab"), 0)

	key := []byte("YELLOW SUBMARINE")
	ct := cpals.B64ToBytes("L77na/nrFsKvynd6HzOoG7GHTLXsTVu9qvY/2syLXzhPweyyMTJULu/6/kXX0KSvoOLSFQ==")
	pt := cpals.CTRDecrypt(key, ct, 0)
	ct2 := cpals.CTREncrypt(key, pt, 0)

	if bytes.Equal(pt, []byte("Yo, VIP Let's kick it Ice, Ice, baby Ice, Ice, baby ")) == false {
		fmt.Printf("%x\n%x\n", pt, "Yo, VIP Let's kick it Ice, Ice, baby Ice, Ice, baby ")
		panic("failed")
	}

	if bytes.Equal(ct2, ct) == false {
		panic("failed")
	}
}
