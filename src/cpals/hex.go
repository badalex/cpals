package cpals

import (
	"encoding/hex"
	"io/ioutil"
)

func htob(i byte) int {
	if i >= '0' && i <= '9' {
		return int(i - '0')
	}

	if i >= 'A' && i <= 'F' {
		return int(i-'A') + 10
	}

	if i >= 'a' && i <= 'f' {
		return int(i-'a') + 10
	}

	panic("bad hex")
}

func HexStrToBytes(in string) []byte {
	var out []byte

	var l = len(in)
	for i := 0; i < l; i++ {
		var str [2]byte
		n := htob(in[i]) << 4
		str[0] = in[i]

		i++
		if i < l {
			n = n | htob(in[i])
			str[1] = in[i]
		}

		out = append(out, byte(n))
	}

	return out
}

func HexBytesToStr(in []byte) string {
	return hex.EncodeToString(in)
}

func HexReadFile(path string) []byte {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	dat = HexStrToBytes(string(dat))
	if err != nil {
		panic(err)
	}

	return dat
}
