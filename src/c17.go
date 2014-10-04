package main

import (
	"cpals"
	"fmt"
)

var key []byte

var blockSize = 16

func one() (ct, iv []byte) {
	if len(key) == 0 {
		//key = []byte("0123456789123456")
		key = cpals.RandBytes(blockSize)
	}

	b64 := []string{
		"MDAwMDAwTm93IHRoYXQgdGhlIHBhcnR5IGlzIGp1bXBpbmc=",
		"MDAwMDAxV2l0aCB0aGUgYmFzcyBraWNrZWQgaW4gYW5kIHRoZSBWZWdhJ3MgYXJlIHB1bXBpbic=",
		"MDAwMDAyUXVpY2sgdG8gdGhlIHBvaW50LCB0byB0aGUgcG9pbnQsIG5vIGZha2luZw==",
		"MDAwMDAzQ29va2luZyBNQydzIGxpa2UgYSBwb3VuZCBvZiBiYWNvbg==",
		"MDAwMDA0QnVybmluZyAnZW0sIGlmIHlvdSBhaW4ndCBxdWljayBhbmQgbmltYmxl",
		"MDAwMDA1SSBnbyBjcmF6eSB3aGVuIEkgaGVhciBhIGN5bWJhbA==",
		"MDAwMDA2QW5kIGEgaGlnaCBoYXQgd2l0aCBhIHNvdXBlZCB1cCB0ZW1wbw==",
		"MDAwMDA3SSdtIG9uIGEgcm9sbCwgaXQncyB0aW1lIHRvIGdvIHNvbG8=",
		"MDAwMDA4b2xsaW4nIGluIG15IGZpdmUgcG9pbnQgb2g=",
		"MDAwMDA5aXRoIG15IHJhZy10b3AgZG93biBzbyBteSBoYWlyIGNhbiBibG93",
	}
	pt := cpals.B64ToBytes(b64[9])

	//pt := []byte("asdf")
	//iv = cpals.Fill(blockSize, "0")

	iv = cpals.RandBytes(blockSize)
	ct = cpals.CBCEncrypt(key, iv, pt)
	cpals.PrintEnc(string(cpals.Pad7(16, pt)), ct)
	//fmt.Printf("%x\n", pt)

	return ct, iv
}

func two(ct, iv []byte) bool {
	pt := cpals.CBCDecrypt(key, iv, ct)
	//fmt.Printf("t: %x\n", pt)
	_, err := cpals.Unpad7(blockSize, pt)

	if err == nil {
		return true
	}

	return false
}

func main() {
	ct, iv := one()
	if two(ct, iv) != true {
		panic("failed a")
	}

	//fmt.Printf("%x\n", ct)

	pt := []byte("")
	// start at the last block
	for i := 0; i < len(ct); i += blockSize {
		c2 := ct[len(ct)-i-blockSize : len(ct)-i]
		var c1 []byte

		if len(ct)-i-blockSize*2 >= 0 {
			c1 = ct[len(ct)-i-blockSize*2 : len(ct)-i-blockSize]
		} else {
			c1 = iv
		}

		p2 := []byte("")
		i2 := []byte("")
		// brute the block
		for y := 0; y < blockSize; y++ {
			k := len(i2) + 1
			my := cpals.Fill(blockSize-k, "\x00") //cpals.RandBytes(blockSize - k)
			//my := cpals.RandBytes(blockSize - k)
			found := false

			for z := 0; z < 256; z++ {
				c1t := append(my, byte(z))

				for x := 0; x < len(i2); x++ {
					c1t = append(c1t, i2[x]^byte(k))
				}

				if len(c1t) != blockSize {
					fmt.Printf("%x %d %d\n", c1t, len(c1t), k)
					panic("cpals")
				}

				try := append(c1t, c2...)
				if two(try, iv) == true {
					found = true

					b := append([]byte(""), byte(z)^byte(k))
					i2 = append(b, i2...)

					b = append([]byte(""), byte(c1[len(c1)-(k)]^(i2[0])))
					p2 = append(b, p2...)
					//fmt.Printf("found %d %d %d %d %d\n", i2[0], z, k, c2[len(c2)-k], p2[0])
					//fmt.Printf("%x %x %x %x\n", c2, c1, c1t, try)
					break
				}
			}

			if found == false {
				fmt.Printf("FAIL\ni2 %x\np2 %x\nk %d\n", i2, p2, k)
				panic("failed")
			}
		}

		b := append([]byte(""), p2...)
		pt = append(b, pt...)
	}

	pt, err := cpals.Unpad7(16, pt)
	if err != nil {
		panic("cpals")
	}
	fmt.Printf("%s\n", pt)
}
