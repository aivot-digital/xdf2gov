package utils

import (
	"regexp"
	"strings"
)

var re = regexp.MustCompile(" +")

func CleanString(str string) string {
	tmp := strings.Replace(str, "\n", "", -1)
	tmp = strings.Replace(tmp, "\r", "", -1)
	tmp = strings.Trim(tmp, " \n\r")
	tmp = re.ReplaceAllString(tmp, " ")
	return tmp
}
