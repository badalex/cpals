package cpals

func HamStr(s1, s2 string) int {
	dist := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			dist++
		}
	}
	return dist
}

func HamBits(s1, s2 []byte) int {
	dist := 0
	for i := 0; i < len(s1); i++ {
		val := s1[i] ^ s2[i]
		for val != 0 {
			dist++
			val &= val - 1
		}
	}
	return dist
}
