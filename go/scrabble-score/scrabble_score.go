package scrabble

import "strings"

type RangeValues []byte

func (lv RangeValues) Has(b byte) bool {
	for _, v := range lv {
		if v == b {
			return true
		}
	}
	return false
}

func Score(w string) int {
	rangeTo1 := RangeValues("aeioulnrst")
	rangeTo2 := RangeValues("dg")
	rangeTo3 := RangeValues("bcmp")
	rangeTo4 := RangeValues("fhvwy")
	rangeTo5 := RangeValues("k")
	rangeTo8 := RangeValues("jx")
	rangeTo10 := RangeValues("qz")
	val := 0
	for _, l := range []byte(strings.ToLower(w)) {
		switch {
		case rangeTo1.Has(l):
			val += 1
		case rangeTo2.Has(l):
			val += 2
		case rangeTo3.Has(l):
			val += 3
		case rangeTo4.Has(l):
			val += 4
		case rangeTo5.Has(l):
			val += 5
		case rangeTo8.Has(l):
			val += 8
		case rangeTo10.Has(l):
			val += 10
		}
	}
	return val
}
