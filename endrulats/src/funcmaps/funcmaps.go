package funcmaps

import (
	"fmt"
	"strings"
)

// https://github.com/Masterminds/sprig
func IndexCount(i int, a int) bool {
	if i%a == 0 {
		return true
	}
	return false
}

func RemoveBrackets(a []string) string {
	var aaaa string
	for _, aa := range a {
		aaa := strings.Replace(aa, "[", "", -1)
		aaaa = strings.Replace(aaa, "]", "", -1)
		fmt.Print(aaaa)
	}
	return aaaa
}
func CSP(a string) string {
	return a
}
