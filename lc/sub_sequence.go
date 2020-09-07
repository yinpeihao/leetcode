package main

import "fmt"

func main() {
	fmt.Println(GetSubSequence([]int{1, 2, 3, 4, 5}))
}

func GetSubSequence(arr []int) [][]int {
	result := make([][]int, 0)
	for i := 2; i <= len(arr); i++ {
		result = append(result, Get(arr, i)...)
	}
	return result
}

func Get(arr []int, length int) [][]int {
	if length == 2 {
		result := make([][]int, 0)
		for i := 0; i < len(arr)-1; i++ {
			for j := i + 1; j < len(arr); j++ {
				curent := []int{arr[i], arr[j]}
				result = append(result, curent)
			}
		}
		return result
	}
	narr := Get(arr[1:], length-1)
	for i := range narr {
		narr[i] = append([]int{arr[0]}, narr[i]...)
	}
	return narr
}
