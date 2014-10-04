package cpals

func Score(in []byte) int {
	s := 0
	for i := 0; i < len(in); i++ {
		if in[i] >= 'A' && in[i] <= 'Z' || in[i] >= 'a' && in[i] <= 'z' {
			s += 2
		}

		if in[i] == '.' || in[i] == '!' || in[i] == ' ' {
			s++
		}

		// non printable
		if in[i] != '\n' && in[i] != '\r' && (in[i] < ' ' || in[i] > '~') {
			return -1
		}
	}
	return s
}
