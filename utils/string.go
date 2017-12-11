package utils

import "strings"

// Remove whitespace and punctuation from start and end of strings.
func CleanStrings(data [][]string) [][]string {
	for i := 0; i < len(data); i++ {
		line := data[i]
		for j := 0; j < len(line); j++ {
			line[j] = strings.Trim(line[j], "\"")
		}
	}
	return data
}

// Return the longest string.
func Longest(s1 string, s2 string) string {
	if len(s1) > len (s2) {
		return s1
	} else {
		return s2
	}
}

func PadRight(str, pad string, length int) string {
	for {
		str += pad
		if len(str) > length {
			return str[0:length]
		}
	}
}

// Truncate the string and add an ellipsis.
func Truncate (s string, l int) string {
	if l < 0 || len(s) < l {
		return s
	} else {
		return s[0:l] + "..."
	}
}