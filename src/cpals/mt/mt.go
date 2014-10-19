package mt

import "fmt"

var _ = fmt.Sprintf("")

type MT struct {
	state [624]uint
	idx   int
}

func (mt *MT) Seed(seed uint) {
	mt.state[0] = seed
	for i := uint(1); i < uint(len(mt.state)); i++ {
		mt.state[i] = 0xffffffff & (1812433253*(mt.state[i-1]^(mt.state[i-1]>>30)) + i)
	}
}

func (mt *MT) Rand() uint {
	if mt.idx == 0 {
		mt.generate()
	}

	y := mt.state[mt.idx]
	y = y ^ (y >> 11)
	y = y ^ (y<<7)&2636928640
	y = y ^ (y<<15)&4022730752
	y = y ^ (y >> 18)

	mt.idx = (mt.idx + 1) % 624

	return y
}

func (mt *MT) generate() {
	for i := range mt.state {
		//fmt.Printf("%d %d\n", i, (i+397)%624)
		y := (mt.state[i] & 0x80000000) + (mt.state[(i+1)%624] & 0x7fffffff)
		mt.state[i] = mt.state[(i+397)%624] ^ (y >> 1)
		if y%2 != 0 {
			mt.state[i] = mt.state[i] ^ 0x9908b0df
		}
	}
}
