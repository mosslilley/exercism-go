/*
Exercism Go Track Bob

Bob is a basic chat bot with limited replies through his hey fuction
*/
package bob

import "strings"

func Hey(remark string) string {
	if string == strings.ToUpper(string) && strings.Contains(string, "?"){
		response := "'Calm down, I know what I'm doing!'"
	} else if strings.Contains(string, "?") {
		response := "Sure."
	} else if string == strings.ToUpper(string){
		response := "Whoa, chill out!"
	} else if strings.TrimSpace(string) == ""{
		response := "Fine. Be that way!"
	} else {
		response := "Whatever."
	}
	return response
}
