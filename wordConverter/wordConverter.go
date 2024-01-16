package wordConverter

func CaseChanger(newCase string, wordsToChange string) string { // change the case of a word based on case type
	wordRunes := []rune(wordsToChange)
	switch newCase {
	case "up":
		for i := 0; i < len(wordRunes); i++ {
			if wordRunes[i] > 96 && wordRunes[i] < 123 {
				wordRunes[i] = wordRunes[i] - 32
			}
		}
	case "low":
		for i := 0; i < len(wordRunes); i++ {
			if wordRunes[i] > 64 && wordRunes[i] < 91 {
				wordRunes[i] = wordRunes[i] + 32
			}
		}
	case "cap":
		if wordRunes[0] > 96 && wordRunes[0] < 123 { // change the first letter to uppercase
			wordRunes[0] = wordRunes[0] - 32
		}
		for i := 1; i < len(wordRunes); i++ { // change the rest to lower case
			if wordRunes[i] > 64 && wordRunes[i] < 91 {
				wordRunes[i] = wordRunes[i] + 32
			}
		}
	}
	return string(wordRunes)
}

var vowels = []string{"a", "e", "i", "o", "u"}

func CheckIfaORan(word string) string{ // checks if a word should have an or a before it
	firstLetter := CaseChanger("low", word[0:1])
	for _, vowel := range vowels {
		if vowel == firstLetter {
			return "an"
		}
	}
	return "a"
}