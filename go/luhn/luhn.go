package luhn

import (
	"regexp"
	"strings"
)

//Given a number determine whether or not it is valid per the Luhn formula.
func Valid(card string) bool {
	//remove white spaces
	card = strings.Join(strings.Split(card, " "), "")
	//verify non digits and size
	if ok, _ := regexp.MatchString(`^[0-9]{2,}$`, card); !ok {
		return false
	}

	l := len(card)
	sum := 0
	//do the sum
	for i, e := range card {
		//convert rune to number
		n := int(e - '0')
		//if index is even from right ot left
		if (l-i)%2 == 0 {
			sum += (2 * n)
			if n >= 5 {
				sum -= 9
			}
		} else {
			sum += n
		}
	}
	//return if is valid
	return sum%10 == 0
}
