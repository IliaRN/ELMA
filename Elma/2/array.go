package main

import "fmt"

func Solution(nums []int, k int) []int {
	if k < 0 || len(nums) == 0 {
		return nums
	}
	r := len(nums) - k%len(nums)
	nums = append(nums[r:], nums[:r]...)
	return nums
}

func SecondSolution(nums []int, k int) []int {
	for k > 0 {
		for i := len(nums) - 1; i > 0; i-- {
			if i > 0 {
				nums[i], nums[i-1] = nums[i-1], nums[i]
			}
		}
		k--
	}
	return nums
}

func main() {
	nums1 := []int{1, 2, 3, 4, 5, 6, 7}
	nums2 := []int{1, 2, 3, 4, 5, 6, 7}
	nums1 = Solution(nums1, 3)
	nums2 = SecondSolution(nums2, 3)
	fmt.Println(nums1)
	fmt.Println(nums2)
}
