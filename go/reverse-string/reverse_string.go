package reverse

import "strings"

func Reverse(input string) string {
	rev := ""
	for _, c := range strings.Split(input, "") {
		rev = c + rev
	}
	return rev
}
