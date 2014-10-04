package cpals

import (
	"bufio"
	"os"
	"strings"
)

var dict map[string]struct{}

func DictLoad() {
	file, err := os.Open("/usr/share/dict/words")
	if err != nil {
		panic(err)
	}

	var empty struct{}
	dict = make(map[string]struct{}, 1024)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		if len(s) > 3 && len(s) <= 7 && strings.Index(s, "'") == -1 {
			if strings.IndexAny(s, "aeio") != -1 {
				dict[strings.ToLower(s)+" "] = empty
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic("failed reading /usr/share/dict/words")
	}

	file.Close()
}

func DictScore(in string) int {
	lower := strings.ToLower(in)

	score := 0
	for k := range dict {
		if strings.Index(lower, k) != -1 {
			score++
		}
	}
	return score
}
