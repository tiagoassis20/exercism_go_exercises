// Convert a phrase to its acronym.
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate a phrase to its acronym
func Abbreviate(s string) string {
	wordParse := regexp.MustCompile(`[[:alpha:]']+`)
	words := wordParse.FindAllString(s, -1)
	abbr := make([]byte, len(words))
	for i, w := range words {
		abbr[i] += strings.ToUpper(w)[0]
	}
	return string(abbr)
}
