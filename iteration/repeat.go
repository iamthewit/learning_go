package iteration

import "strings"

// Repeat given character 5 times
func Repeat(character string, repeatCount int) string {
	// var repeated string

	// for i := 0; i < repeatCount; i++ {
	// 	repeated += character
	// }

	// return repeated

	return strings.Repeat(character, repeatCount)
}
