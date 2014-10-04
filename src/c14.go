package main

import (
	"bytes"
	"cpals"
	"fmt"
)

func oracle(key, junk, mine, yours []byte) []byte {
	var toenc []byte
	toenc = append(toenc, junk...)
	toenc = append(toenc, mine...)
	toenc = append(toenc, yours...)
	return cpals.ECBEncrypt(key, toenc)
}

func main() {
	//key := []byte("0123456790123456")
	key := cpals.RandBytes(16)
	b64 := `Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkg
aGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBq
dXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUg
YnkK`

	b64bit := cpals.B64ToBytes(b64)
	randomPrefix := cpals.RandBytes(42)
	blockSize := 16

	//cpals.PrintEnc(string(append(randomPrefix, b64bit...)), oracle(key, randomPrefix, nil, b64bit))

	//prefixLen := stupidBrute(key, blockSize, randomPrefix, b64bit)
	prefixLen, prefixBlock := getPrefix(key, blockSize, randomPrefix, b64bit)
	if prefixLen != len(randomPrefix) {
		panic("failed")
	}

	dec := []byte("")
	for i := 0; i < len(b64bit); i++ {
		a := cpals.Fill(blockSize-(len(dec)%blockSize)-1, "A")

		if prefixLen != 0 {
			a = append(a, cpals.Fill((prefixBlock*blockSize)+blockSize-prefixLen, "A")...)
		}

		needle := oracle(key, randomPrefix, a, b64bit)
		//cpals.PrintEnc(string(append(append(randomPrefix, a...), b64bit...)), needle)
		dict := make(map[string]int, 256)
		for y := 0; y < 256; y++ {
			out := oracle(key, randomPrefix, append(append(a, dec...), byte(y)), b64bit)

			// they want a dict...
			dict[string(out[0:prefixLen+len(a)+len(dec)+1])] = y
		}

		c, ok := dict[string(needle[0:prefixLen+len(a)+len(dec)+1])]
		if !ok {
			fmt.Printf("%d\n", i)
			panic("didn't find")
		}
		dec = append(dec, byte(c))
	}
	//fmt.Printf("got: %s\n", dec)
}

func getPrefix(key []byte, blockSize int, randomPrefix, b64bit []byte) (int, int) {
	out := oracle(key, randomPrefix, []byte(""), b64bit)
	out1 := oracle(key, randomPrefix, []byte("0"), b64bit)

	// first find which block the prefix ends in
	block := -1
	blockPos := -1
	for i := 0; i < len(out1); i += blockSize {
		if bytes.Equal(out[i:i+blockSize], out1[i:i+blockSize]) == false {
			block = i / blockSize
			blockPos = i
			break
		}
	}

	// now brute the block
	for i := 0; i <= blockSize; i++ {
		a := cpals.Fill(2*blockSize+i, "0")
		try := oracle(key, randomPrefix, a, b64bit)

		//cpals.PrintEnc(string(append(append(randomPrefix, a...), b64bit...)), try)
		//fmt.Printf("%x\n%x\n", try[blockPos+blockSize:blockPos+blockSize*2], try[blockPos+blockSize*2:blockPos+blockSize*3])

		// compare the block after the prefix to the next block, if they are the same we now how long the prefix was
		if bytes.Equal(try[blockPos+blockSize:blockPos+blockSize*2], try[blockPos+blockSize*2:blockPos+blockSize*3]) {
			return blockPos + blockSize - i, block
		}
	}
	return 0, 0
}

// my first attempt, only works on prefixes > blockSize, its also really dumb
func stupidBrute(key []byte, blockSize int, randomPrefix, b64bit []byte) int {
	// we have random data, so we need to pad our cpals out to the next block size
	// do that by trying to attack the first block and adding up the the blockSize
	prefixLen := -1
	//for tryPrefix := 0; tryPrefix < len(randomPrefix)+len(b64bit); tryPrefix++ {
	for tryPrefix := 0; tryPrefix < blockSize; tryPrefix++ {

		dec := []byte("")
		prefixOk := 0
		fmt.Printf("==== tryPrefix %d ====\n", tryPrefix)
		for i := 0; i < blockSize; i++ {
			a := cpals.Fill(blockSize-(len(dec)%blockSize)-1, "A")

			if tryPrefix > 0 {
				a = append(a, cpals.Fill(tryPrefix, "A")...)
			}

			needle := oracle(key, randomPrefix, a, b64bit)
			cpals.PrintEnc(string(append(append(randomPrefix, a...), b64bit...)), needle)

			dict := make(map[string]int, 256)
			for y := 0; y < 256; y++ {
				out := oracle(key, randomPrefix, append(append(a, dec...), byte(y)), b64bit)

				// they want a dict...
				dict[string(out[0:tryPrefix+len(a)+len(dec)+1])] = y
			}

			c, ok := dict[string(needle[0:tryPrefix+len(a)+len(dec)+1])]
			if ok {
				//fmt.Printf("ok %d\n", c)
				prefixOk++
				dec = append(dec, byte(c))
			} else {
				//fmt.Printf("failed\n")
				break
			}
		}

		//panic("cpals")

		// false positive if its 255
		if prefixOk == 16 && dec[0] != 255 && dec[1] != 255 {
			fmt.Printf("p: %d %s\n", prefixOk, dec)
			prefixLen = tryPrefix
			break
		}
	}

	return prefixLen
}
