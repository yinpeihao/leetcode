package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//fmt.Println(shortestPalindrome("aacecaaa"))
	//fmt.Println(shortestPalindrome("abcd"))
	//fmt.Println(shortestPalindrome(""))
	fmt.Println(shortestPalindrome("abb"))
}

func shortestPalindrome(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}
	ns := InitManacher(s)
	maxLen := Manacher(ns) - 1
	fmt.Println(maxLen)
	prefixLen := len(s) - maxLen
	result := make([]byte, len(s)+len(s)-maxLen)
	for i := len(s) - 1; i >= maxLen; i-- {
		result[len(s)-1-i] = s[i]
	}
	copy(result[prefixLen:], s)
	return *(*string)(unsafe.Pointer(&result))
}

func InitManacher(s string) string {
	ns := make([]byte, 0, len(s)*2+1)
	ns = append(ns, '#')
	for i := range s {
		ns = append(ns, s[i], '#')
	}
	return *(*string)(unsafe.Pointer(&ns))
}

func Manacher(s string) int {
	dp := make([]int, len(s))
	maxRight := 0
	maxLen := 1
	pos := 0
	for i := range s {
		j := 2*pos - i
		if i < maxRight {
			dp[i] = minInt(dp[j], maxRight-i)
		} else {
			dp[i] = 1
		}
		for i-dp[i] >= 0 && i+dp[i] < len(s) && s[i-dp[i]] == s[i+dp[i]] {
			dp[i]++
		}
		if dp[i]+i-1 > maxRight {
			pos = i
			maxRight = dp[i] + i - 1
			if maxRight-pos == pos {
				maxLen = maxInt(maxLen, dp[i])
			}
		}
	}
	return maxLen
}

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}
