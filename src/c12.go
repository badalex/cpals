package main

import "cpals"

func oracle(key, mine, yours []byte) []byte {
	var toenc []byte
	toenc = append(toenc, mine...)
	toenc = append(toenc, yours...)
	return cpals.ECBEncrypt(key, toenc)
}

func main() {
	key := []byte("0123456790123456")
	b64 := `Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkg
aGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBq
dXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUg
YnkK`

	b64bit := cpals.B64ToBytes(b64)

	// find the blocksize...
	var b []byte
	blockSize := len(oracle(key, b, b64bit))
	for i := 0; i < 32; i++ {
		out := oracle(key, cpals.Fill(i, "A"), b64bit)
		if len(out) != blockSize {
			blockSize = len(out) - blockSize
			break
		}
	}
	if blockSize != 16 {
		panic("cpals")
	}

	if cpals.DetectECB(16, oracle(key, cpals.Fill(64, "A"), b64bit)) == false {
		panic("cpals")
	}

	var fk []byte
	for i := 0; i < len(b64bit); i++ {
		a := cpals.Fill(blockSize-(len(fk)%blockSize)-1, "A")
		needle := oracle(key, a, b64bit)
		dict := make(map[string]int, 256)
		for y := 0; y < 256; y++ {
			out := oracle(key, append(append(a, fk...), byte(y)), b64bit)

			// they want a dict...
			dict[string(out[0:len(a)+len(fk)+1])] = y
		}

		c, ok := dict[string(needle[0:len(a)+len(fk)+1])]
		if !ok {
			panic("didn't find")
		}
		fk = append(fk, byte(c))
	}
	//fmt.Printf("got: %s\n", fk)
}
