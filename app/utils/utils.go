package utils

import (
	"regexp"
)

func OnlyNumbersString(str string) string {
	reg := regexp.MustCompile("[^0-9]+")

	return reg.ReplaceAllString(str, "")
}