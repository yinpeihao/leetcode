package main

import "fmt"


func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations("2"))
}

var NumberMap = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}


func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	return dfs(digits, []byte{})
}

func dfs(digits string, current []byte) []string {
	//fmt.Println(current)
	if len(digits) == len(current) {
		return []string{string(current)}
	}
	result := make([]string, 0)
	index := len(current)
	for i := 0; i < len(NumberMap[digits[index]]); i++ {
		result = append(result,dfs(digits, append(current, NumberMap[digits[index]][i]))...)
	}
	return result
}
