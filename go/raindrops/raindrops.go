package raindrops

import "fmt"

func Convert(n int) string {
	ret := ""
	if n%3 == 0 {
		ret += "Pling"
	}
	if n%5 == 0 {
		ret += "Plang"
	}
	if n%7 == 0 {
		ret += "Plong"
	}

	if ret == "" {
		return fmt.Sprintf("%d", n)
	}

	return ret
}
