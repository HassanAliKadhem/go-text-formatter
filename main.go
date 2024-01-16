package main

import (
	"fmt"
	"strconv"
	"strings"
	"os"

	// my packages
	"main/numConverter"
	"main/wordConverter"
	"main/fileOperation"
)

func main() {
	if len(os.Args) == 3 {
		fileContent, err := fileOperation.ReadFile(os.Args[1])
		if err != nil {
			fmt.Print(err.Error() + "\n")
		} else {
			contentSlice := strings.Split(fileContent, " ")
			fmt.Println(fileContent)
			modifiedText := parseFile(contentSlice)
			saveErr := fileOperation.SaveFile(os.Args[2], modifiedText)
			if saveErr != nil {
				fmt.Print(saveErr.Error() + "\n")
			}
		}
	} else if len(os.Args) == 1 {
		fmt.Println("File name missing")
	} else {
		fmt.Println("You need two arguments, input file and output file")
	}
}

func parseFile(contentSlice []string) string {
	for i := 0; i < len(contentSlice); i++ {
		if strings.Contains(contentSlice[i], "(") {
			wordsToModify := 1
			if strings.Contains(contentSlice[i], ",") {
				wordsToModify, _ = strconv.Atoi(contentSlice[i+1][:len(contentSlice[i+1])-1]) // get the string without the last character which is )
			}
			wordsSliceToModify := contentSlice[i - wordsToModify: i]
			currentModification := contentSlice[i]
			for ii := 0; ii < len(wordsSliceToModify); ii++ {
				if strings.Contains(currentModification, "up") {
					wordsSliceToModify[ii] = wordConverter.CaseChanger("up", wordsSliceToModify[ii])
				} else if strings.Contains(currentModification, "cap") {
					wordsSliceToModify[ii] = wordConverter.CaseChanger("cap", wordsSliceToModify[ii])
				} else if strings.Contains(currentModification, "low") {
					wordsSliceToModify[ii] = wordConverter.CaseChanger("low", wordsSliceToModify[ii])
				} else if strings.Contains(currentModification, "hex") {
					wordsSliceToModify[ii] = numConverter.ConvertFromHex(wordsSliceToModify[ii])
				} else if strings.Contains(currentModification, "bin") {
					wordsSliceToModify[ii] = numConverter.ConvertFromBinary(wordsSliceToModify[ii])
				}
			}

			// remove after use
			if wordsToModify == 1 { // if only one
				contentSlice = append(contentSlice[:i], contentSlice[i+1:]...) // recreate the slice without the commands (up), (low), etc...
			} else { // if more than one
				contentSlice = append(contentSlice[:i], contentSlice[i+2:]...)
			}
		} else if contentSlice[i] == "a" || contentSlice[i] == "an" {
			contentSlice[i] = wordConverter.CheckIfaORan(contentSlice[i+1])
		}
	}
	return joinStringWithPunctuations(contentSlice)
}



func joinStringWithPunctuations(stringSlice []string) string { // join the string with spaces while using the correct spacing for punctuations
	result := ""
	result += stringSlice[0]
	for i := 1; i < len(stringSlice); i++ {
		if checkIfContainsPunctuation(stringSlice[i]) {
			result += fixPunctuationInString(stringSlice[i])
		} else {
			result += " " + stringSlice[i]
		}
	}

	return result
}

var punctuations = []string{".", ",", "!", "?", ":", ";"}
func checkIfContainsPunctuation(stringToTest string) bool{ // return true if the string contains punctuations
	for _, v := range punctuations {
		if strings.Index(stringToTest, v) != -1 { // check the index of the punctuation in the string, -1 will be returned if not found so this only runs if found
			return true
		}
	}
	return false
}

func fixPunctuationInString(stringToFix string) string{ // this function will fix the spacing around punctuations if it is with other characters
	result := ""
	if !checkIfContainsPunctuation(stringToFix[0:1]) {
		result = " "
	}
	letters := strings.Split(stringToFix, "")
	for i := 0; i < len(letters); i++ {
		result += letters[i]
		if i < len(letters) - 1 && checkIfContainsPunctuation(letters[i]) && !checkIfContainsPunctuation(letters[i+1]) { // check if the current and next characters is punctuations to avoid adding the space
			result += " "
		}
	}
	return result
}
