package main

import "cpals"

func main() {
	out := cpals.XorBytes(cpals.HexStrToBytes("1c0111001f010100061a024b53535009181c"), cpals.HexStrToBytes("686974207468652062756c6c277320657965"))

	if cpals.HexBytesToStr(out) != "746865206b696420646f6e277420706c6179" {
		panic("failed")
	}
}
