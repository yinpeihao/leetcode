package main

import "strings"

func main() {

}

func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	eFlag, dotFlag, numFlag := false, false, false

	for i, c := range s {
		if c >= '0' && c <= '9' {
			numFlag = true
		} else if (c == 'e' || c == 'E') && !eFlag && numFlag {
			eFlag = true
			numFlag = false //e后必须有数字
		} else if c == '.' && !eFlag && !dotFlag {
			dotFlag = true
		} else if (c == '+' || c == '-') && (i == 0 || s[i-1] == 'e' || s[i-1] == 'E') {

		} else {
			return false
		}
	}
	return numFlag
}
