package main

import (
	"unsafe"
)

func main() {
	//fmt.Println(shortestPalindrome("aacecaaa"))
	//fmt.Println(shortestPalindrome("abcd"))
	//fmt.Println(shortestPalindrome(""))
}

func shortestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	ans := s[0:1]
	for l := 0; l < len(s); l++ {
		for i := 0; i < len(s)-l; i++ {
			j := i + l
			if l == 0 {
				dp[i][j] = true
			} else if l == 1 && s[i] == s[j] {
				dp[i][j] = true
			} else if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			}
			if dp[i][j] && l+1 > len(ans) && i == 0 {
				ans = s[i : j+1]
			}
		}
	}
	if len(ans) == len(s) {
		return s
	}
	prefixLen := len(s) - len(ans)
	result := make([]byte, len(s)+len(s)-len(ans))
	for i := len(s) - 1; i >= len(ans); i-- {
		result[len(s)-1-i] = s[i]
	}
	copy(result[prefixLen:], s)

	return *(*string)(unsafe.Pointer(&result))
}
