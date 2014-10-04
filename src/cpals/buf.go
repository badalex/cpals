package cpals

import "fmt"

func Fill(n int, b string) []byte {
	var out []byte
	for i := 0; i < n; i++ {
		out = append(out, b[0])
	}
	return out
}

func PrintEnc(u string, enc []byte) {

	var b = 16 * 3
	// we print up to 3 blocks per line
	for y := 0; y < len(enc); y += b {
		for i := 0; i < b; i++ {
			pos := i + y
			if pos >= len(enc) {
				break
			}

			if i%16 == 0 && i != 0 {
				fmt.Printf("  ")
			}

			if pos < len(u) {
				if u[pos] < 48 {
					fmt.Printf("%02x", u[pos])
				} else {
					fmt.Printf("%c_", u[pos])
				}
			} else {
				fmt.Printf("__")
			}
		}

		fmt.Printf("\n")
		for i := 0; i < b; i++ {
			pos := i + y
			if pos >= len(enc) {
				break
			}

			if i%16 == 0 && i != 0 {
				fmt.Printf("  ")
			}
			fmt.Printf("%02x", enc[pos])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}
