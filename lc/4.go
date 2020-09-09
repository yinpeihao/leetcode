package main

import "fmt"

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{0, 0, 0, 0, 0}, []int{-1, 0, 0, 0, 0, 0, 1}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	k := len(nums1) + len(nums2)
	if k%2 == 1 {
		return float64(findKth(nums1, nums2, k/2))
	} else {
		return float64(findKth(nums1, nums2, k/2-1)+findKth(nums1, nums2, k/2)) / 2.0
	}
}

func findKth(nums1, nums2 []int, k int) int {
	for {
		//fmt.Println(k)
		if len(nums1) == 0 {
			return nums2[k]
		}
		if len(nums2) == 0 {
			return nums1[k]
		}
		if k == 0 {
			return min(nums1[0], nums2[0])
		}
		mid := (k - 1) / 2
		i1 := min(mid, len(nums1)-1)
		i2 := min(mid, len(nums2)-1)
		num1, num2 := nums1[i1], nums2[i2]
		if num1 <= num2 {
			nums1 = nums1[i1+1:]
			k -= i1 + 1
		} else {
			nums2 = nums2[i2+1:]
			k -= i2 + 1
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
