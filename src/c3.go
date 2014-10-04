package main

import "cpals"

type top struct {
	bytes []byte
	score int
	has   bool
}

func put_top(tops []top, in []byte, s int) {
	for i := 0; i < cap(tops); i++ {
		if tops[i].has == false {
			tops[i] = top{in, s, true}
			return
		}

		if s > tops[i].score {
			tmp := top{in, s, true}
			for p := i; p < cap(tops); p++ {
				t2 := tops[p]
				tops[p] = tmp
				tmp = t2
			}
			return
		}
	}
}

func main() {
	bytes := cpals.HexStrToBytes("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")

	top_scores := make([]top, 10)
	for i := 0; i < 256; i++ {
		out := cpals.XorBytes(bytes, []byte(string(i)))

		s := cpals.Score(out)
		put_top(top_scores, out, s)
	}

	//for i := 0; i < len(top_scores); i++ {
	//	if top_scores[i].has {
	//		fmt.Printf("score: %d, %s\n", top_scores[i].score, top_scores[i].bytes)
	//	}
	//}

	if string(top_scores[0].bytes) != "Cooking MC's like a pound of bacon" {
		panic("failed")
	}
}
