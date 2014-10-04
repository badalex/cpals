package main

import (
	"bytes"
	"cpals"
	"io/ioutil"
	"runtime"
)

type top struct {
	bytes []byte
	score int
	has   bool
	line  int
}

func put_top(tops []top, ntop top) {
	for i := 0; i < cap(tops); i++ {
		if tops[i].has == false {
			tops[i] = ntop
			return
		}

		if ntop.score > tops[i].score {
			tmp := ntop
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
	runtime.GOMAXPROCS(runtime.NumCPU())

	//	f, err := os.Create("c4.prof")
	//	if err != nil {
	//		panic(err)
	//	}
	//	pprof.StartCPUProfile(f)
	//	defer pprof.StopCPUProfile()

	read, err := ioutil.ReadFile("data/4.txt")
	if err != nil {
		panic(err)
	}

	lines := bytes.Split(read, []byte("\n"))
	read = nil

	ct := make(chan top)
	top_scores := make([]top, 1)
	for l := 0; l < len(lines); l++ {
		hexed := cpals.HexStrToBytes(string(lines[l]))
		for i := 0; i < 127; i++ {
			go func(hexed []byte, l int, i byte) {
				out := cpals.XorBytes(hexed, []byte(string(i)))

				//s := cpals.DictScore(out)
				s := cpals.Score(out)
				ntop := top{out, s, true, l + 1}
				ct <- ntop
			}(hexed, l, byte(i))
		}
	}

	for l := 0; l < len(lines); l++ {
		for i := 0; i < 127; i++ {
			ntop := <-ct
			put_top(top_scores, ntop)
		}
	}

	//for i := 0; i < len(top_scores); i++ {
	//	if top_scores[i].has && top_scores[i].score > 1 {
	//		fmt.Printf("score: %d line: %d, %s\n", top_scores[i].score, top_scores[i].line, top_scores[i].bytes)
	//	}
	//}

	if string(top_scores[0].bytes) != "Now that the party is jumping\n" {
		panic("failed")
	}
}
