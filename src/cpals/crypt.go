package cpals

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"encoding/binary"
	"fmt"
	"math/rand"
	"time"
)

func ECBDecrypt(key, in []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	out := make([]byte, len(in))
	for i := 0; i < len(in); i += len(key) {
		block.Decrypt(out[i:i+len(key)], in[i:i+len(key)])
	}

	return out
}

func ECBEncrypt(key, in []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	in = Pad7(len(key), in)
	if len(in)%len(key) != 0 {
		fmt.Printf("input not padded right %d %d %d", len(key), len(in), len(in)%len(key))
		panic("crap")
	}

	out := make([]byte, len(in))

	for i := 0; i < len(in); i += len(key) {
		cbc := in[i : i+len(key)]
		block.Encrypt(out[i:i+len(key)], cbc)
	}

	return out
}

func CTREncrypt(key, pt []byte, nonce int64) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	keystream := []byte("")
	var counter int64 = 0
	for i := 0; i < len(pt); i += 16 {
		tmp := make([]byte, 16)
		buf := new(bytes.Buffer)
		binary.Write(buf, binary.LittleEndian, nonce)
		binary.Write(buf, binary.LittleEndian, counter)
		block.Encrypt(tmp, buf.Bytes())
		keystream = append(keystream, tmp...)
		counter++
	}

	return XorBytes(pt, keystream)
}

func CTRDecrypt(key, ct []byte, nonce int64) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	keystream := []byte("")
	var counter int64 = 0
	for i := 0; i < len(ct); i += 16 {
		tmp := make([]byte, 16)
		buf := new(bytes.Buffer)
		binary.Write(buf, binary.LittleEndian, nonce)
		binary.Write(buf, binary.LittleEndian, counter)
		block.Encrypt(tmp, buf.Bytes())
		keystream = append(keystream, tmp...)
		counter++
	}

	return XorBytes(ct, keystream)
}
func CBCEncrypt(key, iv, in []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	in = Pad7(len(key), in)
	if len(in)%len(key) != 0 {
		fmt.Printf("input not padded right %d %d %d", len(key), len(in), len(in)%len(key))
		panic("crap")
	}

	out := make([]byte, len(in))
	tmp := iv

	for i := 0; i < len(in); i += len(key) {
		cbc := in[i : i+len(key)]
		cbc = XorBytes(tmp, cbc)

		block.Encrypt(out[i:i+len(key)], cbc)

		//fmt.Printf("exor: %x %x %x %s\n", tmp, cbc, out[i:i+len(key)], in[i:i+len(key)])
		tmp = out[i : i+len(key)]
	}

	return out
}

func CBCDecrypt(key, iv []byte, in []byte) []byte {

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	out := make([]byte, len(in))
	dec := make([]byte, 16)

	tmp := iv
	for i := 0; i < len(in); i += len(key) {
		cbc := in[i : i+len(key)]
		block.Decrypt(dec, cbc)
		fin := XorBytes(dec, tmp)
		tmp = []byte{}
		tmp = append(tmp, cbc...)

		copy(out[i:i+len(key)], fin)
	}

	return out
}

func DetectECBScanner(blockSize int, scanner *bufio.Scanner) (lineno int, err error) {
	lineno = 0
	for scanner.Scan() {
		line := scanner.Bytes()
		if DetectECB(blockSize, line) {
			return lineno + 1, nil
		}
		lineno++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return 0, nil
}

// I dont see a reason to make this nice until I have 3 different callers
// returns lineno on sucess, starting at 1, <= 0 means we did not find anything
func DetectECB(blockSize int, line []byte) bool {
	var blocks map[string]int = make(map[string]int, (len(line)/blockSize)/2)

	var matched = 0
	for i := 0; i < len(line); i += blockSize {
		if blocks[string(line[i:i+blockSize])] > 0 {
			matched++
		}
		blocks[string(line[i:i+blockSize])]++
	}
	if matched > 1 {
		return true
	}
	return false
}

var seeded bool

func seed() {
	if !seeded {
		rand.Seed(time.Now().UTC().UnixNano())
		seeded = true
	}
}

func RandBytes(size int) []byte {
	seed()

	var out []byte
	for i := 0; i < size; i++ {
		out = append(out, byte(rand.Intn(255)))
	}

	return out
}

func EncOracle(key, in []byte) ([]byte, string) {
	seed()

	var tmp []byte

	// prepend 5-10 random bytes
	num := 5 + rand.Intn(5)
	for i := 0; i < num; i++ {
		tmp = append(tmp, byte(rand.Intn(255)))
	}
	toenc := append(make([]byte, len(in)+len(tmp)), tmp...)
	toenc = append(toenc, in...)

	// append 5-10 random bytes after
	tmp = make([]byte, 1)
	num = 5 + rand.Intn(5)
	for i := 0; i < num; i++ {
		tmp = append(tmp, byte(rand.Intn(255)))
	}
	toenc = append(toenc, tmp...)

	if rand.Intn(100) > 50 {
		return CBCEncrypt(key, RandBytes(len(key)), toenc), "CBC"
	}
	return ECBEncrypt(key, toenc), "ECB"
}
