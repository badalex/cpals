package cpals

func XorBytes(a, b []byte) []byte {
	out := make([]byte, len(a))
	for i := 0; i < len(a); i++ {
		out[i] = byte(int(a[i]) ^ int(b[i%len(b)]))
	}
	return out
}
