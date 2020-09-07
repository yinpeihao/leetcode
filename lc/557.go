package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}

func reverseWords(s string) string {
	result := make([]byte, 0, len(s))
	spaceList := []int{}
	for i, b := range s {
		if b == ' ' {
			spaceList = append(spaceList, i)
		}
	}
	spaceList = append(spaceList, len(s))
	current := -1
	for _, next := range spaceList {
		for i := next - 1; i > current; i-- {
			result = append(result, s[i])
		}
		if next != len(s) {
			result = append(result, ' ')
		}
		current = next
	}
	return *(*string)(unsafe.Pointer(&result))
}
