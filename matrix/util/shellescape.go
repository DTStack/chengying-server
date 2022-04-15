package util

import (
	"regexp"
	"strings"
)

var pattern = regexp.MustCompile(`[^\w@%+=:,./-]`)

// ShellQuote returns a shell-escaped version of the string s. The returned value
// is a string that can safely be used as one token in a shell command line.
func ShellQuote(s string) string {
	if len(s) == 0 {
		return "''"
	}
	if pattern.MatchString(s) {
		return "'" + strings.Replace(s, "'", "'\"'\"'", -1) + "'"
	}

	return s
}
