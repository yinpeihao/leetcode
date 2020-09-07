package main

import "fmt"

func main() {
	fmt.Println(getPermutation(3, 3))
	fmt.Println(getPermutation(4, 9))
}

var (
	cnt, gn, gk int
	tag         []bool
	ans         []byte
)

func getPermutation(n int, k int) string {
	gn, gk = n, k
	cnt = 0
	tag = make([]bool, n+1)
	seq := make([]byte, 0, n)
	dfs(seq)
	return string(ans)
}

func dfs(seq []byte) {
	if len(seq) == gn {
		cnt++
		ans = seq
		return
	}
	for i := 1; i <= gn; i++ {
		if !tag[i] {
			tag[i] = true
		} else {
			continue
		}
		dfs(append(seq, byte(i+int('0'))))
		if cnt == gk {
			return
		}
		tag[i] = false
	}
	return
}
