package utils

import (
	"regexp"
	"strconv"
	"unicode"
)

func hasSymbolOrSpace(str string) bool {
	for _, letter := range str {
		if unicode.IsSymbol(letter) {
			return true
		}
		if unicode.IsSpace(letter) {
			return true
		}
	}
	return false
}

func IsValidNameString(str string) bool {
	if len(str) == 0 {
		return false
	}
	if hasSymbolOrSpace(str) {
		return false
	}
	return true
}

func IsValidEmail(str string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(str)
}

func IsNumber(str string) bool {
	if _, err := strconv.Atoi(str); err != nil {
		return false
	}
	return true
}
