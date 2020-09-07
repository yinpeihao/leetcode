package main

import (
	"fmt"

)

func main() {
	fmt.Println(getPermutation(3, 3))
}

func getPermutation(n int, k int) string {
	fact := make([]int, n+1)
	fact[0] = 1
	for i := 1; i < n; i++ {
		fact[i] = i * fact[i-1]
	}
	i := n - 1
	ans := make([]byte, 0, n)
	mp := make([]int, n+1)
	k--
	for len(ans) != n {
		index := k/fact[i] + 1
		k = k % fact[i]
		i--
		for j := 1; j <= n; j++ {
			if mp[j] == 0 {
				index--
			}
			if index == 0 {
				mp[j] = 1
				ans = append(ans, byte(j)+'0')
				//fmt.Println(ans)
				break
			}
		}
	}
	return string(ans)
}
