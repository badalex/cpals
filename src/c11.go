package main

import "cpals"

func main() {
	plain := []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	out, mode := cpals.EncOracle(cpals.RandBytes(16), plain)
	if cpals.DetectECB(16, out) && mode != "ECB" {
		panic("cpals")
	}
}
