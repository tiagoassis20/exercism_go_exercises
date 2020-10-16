package isogram

import "strings"

func IsIsogram(w string) bool {
	letters := make(map[byte]int)
	for _, c := range []byte(strings.ToLower(w)) {
		if a, ok := letters[c]; !ok || (c < 'a' || c > 'z') {
			letters[c] = a + 1
		} else {
			return false
		}
	}
	return true
}
