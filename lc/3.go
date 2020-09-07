package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	mark := make(map[byte]bool)
	result := make([]int, len(s))
	max := 1
	for i := 0; i < len(s); i++ {
		result[i] = 1
		mark[s[i]] = true
		var j int
		if i == 0 {
			j = 1
		} else {
			j = result[i-1] + i - 1
		}
		if j == i {
			j++
		}
		for ; j < len(s); j++ {
			if mark[s[j]] {

				break
			}
			mark[s[j]] = true
		}
		result[i] = j - i
		if result[i] > max {
			max = result[i]
		}
		delete(mark, s[i])
	}
	//for _, v := range result {
	//	fmt.Println(v)
	//}
	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abc"))
	fmt.Println(lengthOfLongestSubstring("a"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
