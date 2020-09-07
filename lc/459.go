package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(repeatedSubstringPattern("a"))
}

//func repeatedSubstringPattern(s string) bool {
//	if len(s) == 1 {
//		return true
//	}
//	lenS := len(s)
//	sqrtLen := int(math.Sqrt(float64(lenS)))
//	numbers := make([]int,0,1)
//	numbers = append(numbers,1)
//	for i := 2;i<=sqrtLen;i++ {
//		if lenS % i == 0 {
//			numbers = append(numbers,i,lenS/i)
//		}
//	}
//	for _,number := range numbers {
//		subStr := s[:number]
//		begin := number
//		subLen := number
//		for begin+subLen <= len(s){
//			if subStr == s[begin:begin+subLen] {
//				begin += subLen
//			} else {
//				break
//			}
//		}
//		if begin == len(s) {
//			return true
//		}
//	}
//	return false
//}

func repeatedSubstringPattern(s string) bool {
	return strings.Index((s + s)[1:], s) != len(s)
}
