package twelve

import (
	"fmt"
	"strings"
)

var gifts = strings.Split(" twelve Drummers Drumming, eleven Pipers Piping, ten Lords-a-Leaping, nine Ladies Dancing, eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree", ",")

var days = []string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}

func Verse(n int) string {
	if n == 1 {
		return fmt.Sprintf("On the first day of Christmas my true love gave to me: a Partridge in a Pear Tree.")
	}
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me:%s.", days[n-1], strings.Join(gifts[12-n:], ","))

}
func Song() string {
	verses := make([]string, 12)
	for i := 0; i < 12; i++ {
		verses[i] = Verse(i + 1)
	}
	return strings.Join(verses, "\n")
}
