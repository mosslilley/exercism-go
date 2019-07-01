/*
Exercism Go Track Acronym

Takes in a phrase and makes an acronym from it
*/
package acronym

import "strings"
import "regexp"

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {

	acronym := ""

	regx := regexp.MustCompile(`[- ,_]`)
	words := regx.Split(s, -1)
	// fmt.Println(words)
	for _, word := range words {
		upperWord := strings.ToUpper(word)

		if len(word) > 0{
			firstLetter := string(upperWord[0])
			acronym += firstLetter
		}
	}
	return acronym
}
