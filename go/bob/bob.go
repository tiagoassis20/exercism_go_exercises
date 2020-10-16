// Bob is a lackadaisical teenager. In conversation, his responses are very limited.
// Bob answers 'Sure.' if you ask him a question, such as "How are you?".
// He answers 'Whoa, chill out!' if you YELL AT HIM (in all capitals).
// He answers 'Calm down, I know what I'm doing!' if you yell a question at him.
// He says 'Fine. Be that way!' if you address him without actually saying anything.
// He answers 'Whatever.' to anything else.
// Bob's conversational partner is a purist when it comes to written communication and always follows normal rules regarding sentence punctuation in English.

package bob

import (
	"regexp"
	"strings"
)

// Hey function gives the bob answer to a given remark
func Hey(remark string) string {
	var onlyNumber, question, silence *regexp.Regexp
	var err error
	if onlyNumber, err = regexp.Compile(`^([0-9]|[[:punct:]]|\s)*$`); err != nil {
		panic(err)
	}
	if question, err = regexp.Compile(`\?\s*$`); err != nil {
		panic(err)
	}
	if silence, err = regexp.Compile(`^\s*$`); err != nil {
		panic(err)
	}

	if silence.MatchString(remark) {
		return "Fine. Be that way!"
	}
	if question.MatchString(remark) {
		if onlyNumber.MatchString(remark) || remark != strings.ToUpper(remark) {
			return "Sure."
		} else {
			return "Calm down, I know what I'm doing!"
		}
	} else {
		if onlyNumber.MatchString(remark) || remark != strings.ToUpper(remark) {
			return "Whatever."
		} else {
			return "Whoa, chill out!"
		}
	}
}
