package drunktext

import (
	"math/rand"
	"strings"
	"time"
)

// MakeDrunk makes it's inputText look drunk based on passed
// "percentSober" (1-100) value
func MakeDrunk(inputText string, percentSober int) string {
	rand.Seed(time.Now().UnixNano())
	var outputText string
	for _, inputLetter := range inputText {
		changeLetter := rand.Intn(100) > percentSober
		var missedLetter rune
		letter := []rune(strings.ToLower(string(inputLetter)))[0]

		if len(string(MissedKeys[letter])) > 0 {
			missedLetter = MissedKeys[letter][rand.Intn(3)]
		} else {
			missedLetter = letter
		}

		if changeLetter {
			outputText = outputText + string(missedLetter)
		} else {
			outputText = outputText + string(letter)
		}
	}
	return outputText
}
