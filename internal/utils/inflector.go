package utils

import (
	"regexp"
	"strings"
)

func Sanitize(str string, removePattern string) (string, error) {
	exp, err := regexp.Compile(removePattern)
	if err != nil {
		return "", err
	}

	return exp.ReplaceAllString(str, ""), nil
}

func Sentenize(str string) string {
	str = strings.TrimSpace(str)
	if str == "" {
		return ""
	}

	lastChar := str[len(str)-1:]
	if lastChar != "." && lastChar != "?" && lastChar != "!" {
		return str + "."
	}

	return str
}
