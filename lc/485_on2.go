package main

func main() {

}

func PredictTheWinner(nums []int) bool {
	return calc(0, len(nums)-1, nums) >= 0
}

func calc(begin, end int, nums []int) int {
	if begin == end {
		return nums[begin]
	}
	beginVal := nums[begin] - calc(begin+1, end, nums)
	endVal := nums[end] - calc(begin, end-1, nums)
	return MaxInt(beginVal, endVal)
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
