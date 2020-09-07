package main

import "fmt"

func main() {
	fmt.Println(getBitLen(4))
	fmt.Println(getBitLen(3))
	fmt.Println(getBitLen(2))
	fmt.Println(getBitLen(1))
	fmt.Println(getBitLen(0))

	a := []int{1, 2, 3, 4, 5}
	b := a[1:2]
	b = append(b,10)
	fmt.Println(a)
	fmt.Println(cap(a), len(a))
	fmt.Println(b)
	fmt.Println(cap(b), len(b))
}
func getBitLen(m int) int {
	bitlen := 0
	for {
		m = m >> 1
		if m != 0 {
			bitlen++
		} else {
			break
		}
	}
	return bitlen
}
