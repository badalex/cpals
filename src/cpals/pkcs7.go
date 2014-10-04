package cpals

import "errors"

func Pad7(blockSize int, in []byte) []byte {
	pad := blockSize - (len(in) % blockSize)

	if pad == 0 {
		return append(in, Fill(blockSize, string(blockSize))...)
	}

	for i := 0; i < pad; i++ {
		in = append(in, byte(pad))
	}

	return in
}

func Unpad7(blockSize int, in []byte) ([]byte, error) {
	l := len(in)
	b := int(in[l-1])

	if b > blockSize {
		return []byte(""), errors.New("bad padding")
	}

	if b == 0 {
		return []byte(""), errors.New("bad padding")
	}

	if b <= blockSize {
		for i := 0; i < b; i++ {
			if in[l-i-1] != byte(b) {
				return []byte(""), errors.New("bad padding")
			}
		}
	}

	return in[0 : l-b], nil
}
