package main

import (
	"cpals"
)

type top_score struct {
	bytes []byte
	score int
	has   bool
}

func put_top(tops []top_score, in []byte, s int) {
	for i := 0; i < cap(tops); i++ {
		if tops[i].has == false {
			tops[i] = top_score{in, s, true}
			return
		}

		if s > tops[i].score {
			tmp := top_score{in, s, true}
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
	//fmt.Printf("%d\n", ham_dist_str("karolin", "kathrin"))
	//fmt.Printf("%d\n", ham_dist_str("this is a test", "wokka wokka!!!"))
	//fmt.Printf("%d\n", ham_dist_bits([]byte("this is a test"), []byte("wokka wokka!!!")))

	out := cpals.B64ReadFile("data/6.txt")

	type top_keysize struct {
		dist float64
		ks   int
	}
	var top_ks [1]top_keysize

	for keysize := 2; keysize < 40; keysize++ {
		dist := 0

		for i := 0; i < 32; i += 2 {
			e1 := (i + 1) * keysize
			e2 := (i + 2) * keysize
			dist += cpals.HamBits(out[i*keysize:e1], out[e1:e2])
		}

		distf := float64(dist) / float64(keysize) * 32 / 2
		//fmt.Printf("dist: %d %f\n", keysize, distf)

		for i := 0; i < len(top_ks); i++ {
			if top_ks[i].ks == 0 || distf < top_ks[i].dist {
				top_ks[i] = top_keysize{distf, keysize}
				break
			}
		}
	}

	for i := 0; i < len(top_ks); i++ {
		//fmt.Printf("ks: %d dist: %f\n", top_ks[i].ks, top_ks[i].dist)

		var trans_blocks [][]byte
		trans_blocks = make([][]byte, (len(out) / top_ks[i].ks))
		for o := 0; o < len(out); o++ {
			//fmt.Printf("put %d in %d\n", o, o%top_ks[i].ks)
			trans_blocks[o%top_ks[i].ks] = append(trans_blocks[o%top_ks[i].ks], out[o])
		}

		var fin []byte
		for o := 0; o < len(trans_blocks); o++ {
			top_scores := make([]top_score, 1)
			for n := 0; n < 127; n++ {
				ns := string(n)
				nb := []byte(ns)

				xorout := cpals.XorBytes(trans_blocks[o], nb)

				s := cpals.Score(xorout)
				put_top(top_scores, nb, s)
			}

			for i := 0; i < len(top_scores); i++ {
				if top_scores[i].has && top_scores[i].score > 0 {
					//fmt.Printf("score: %d, '%s'\n", top_scores[i].score, top_scores[i].bytes)
					fin = append(fin, top_scores[i].bytes[0])
					if top_scores[i].bytes[0] == 0 {
						fin = append(fin, ' ')
					}
				}
			}
		}
		//fmt.Printf("final key: %s\n", fin)

		//fin = []byte("Terminator X: Bring the noise")

		//out2 := cpals.XorBytes(out, fin)
		//fmt.Printf("%s\n", out2)
	}

}
