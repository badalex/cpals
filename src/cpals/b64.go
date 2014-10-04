package cpals

import (
	"encoding/base64"
	"io/ioutil"
)

func B64ReadFile(path string) []byte {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return B64ToBytes(string(dat))
}

func B64ToBytes(b64 string) []byte {
	dat, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		panic(err)
	}
	return dat
}

var b64_enc = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func StrToB64(in string) string {
	var out []byte

	// base64 works on 3 bytes at a time, taking the value of 6 bits at a time
	// as an index into b64_enc
	var l = len(in)
	for i := 0; i < l; i++ {
		var n uint8

		// first 6 bits
		n = in[i] >> 2
		out = append(out, b64_enc[n])

		// 2nd 6 bits
		if i+1 < l {

			// keep last 2 bits
			n = in[i] & 0x03
			// move them up
			n = n << 4

			// keep first 4 bits
			t := in[i+1] & 0xf0
			// move them down
			t = t >> 4

			// combined them
			n = n | t

			out = append(out, b64_enc[n])
		}

		// 3rd and 4th
		if i+2 < l {

			// 3rd
			// keep last 4 bits
			n = in[i+1] & 0x0f
			// move them up
			n = n << 2

			// keep first 2 bits
			t := in[i+2] & 0xc0
			// move them down
			t = t >> 6

			// combined them
			n = n | t
			out = append(out, b64_enc[n])

			// 4th
			// keep last 6 bits
			n = in[i+2] & 0x3f
			out = append(out, b64_enc[n])
		}

		i += 2
	}

	return string(out)
}
