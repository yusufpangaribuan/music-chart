package util

import "strings"

/*
	RemoveAllWhiteSpace
	example:
		1. RemoveAllWhiteSpace("            tes") = "tes"
		2. RemoveAllWhiteSpace("        ") = ""
*/
func RemoveAllWhiteSpace(str string) string {
	return strings.Replace(str, " ", "", -1)
}
