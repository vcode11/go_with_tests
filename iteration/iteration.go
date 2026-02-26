package iteration

import "strings"

func Repeat(text string, numTimes int) string{
	var sb strings.Builder
	for range numTimes{
		sb.WriteString(text)
	}
	return sb.String()
}