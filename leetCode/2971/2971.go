package main

import (
	"fmt"
	"sort"
)

// 2971. Find Polygon With the Largest Perimeter
// Medium
// Topics
// Companies
// Hint
// You are given an array of positive integers nums of length n.

// A polygon is a closed plane figure that has at least 3 sides. The longest side of a polygon is smaller than the sum of its other sides.

// Conversely, if you have k (k >= 3) positive real numbers a1, a2, a3, ..., ak where a1 <= a2 <= a3 <= ... <= ak and a1 + a2 + a3 + ... + ak-1 > ak, then there always exists a polygon with k sides whose lengths are a1, a2, a3, ..., ak.

// The perimeter of a polygon is the sum of lengths of its sides.

// Return the largest possible perimeter of a polygon whose sides can be formed from nums, or -1 if it is not possible to create a polygon.
func largestPerimeter(nums []int) int64 {
	sort.Ints(nums) //
	fmt.Println(nums)
	var previousElementsSum int64 = 0
	var resp int64 = -1
	for i := 0; i < len(nums); i++ {
		if i >= 2 && nums[i] < int(previousElementsSum) {
			resp = previousElementsSum + int64(nums[i])
		}
		previousElementsSum += int64(nums[i])
	}
	return resp
}

func main() {
	fmt.Println(largestPerimeter([]int{1, 12, 1, 2, 5, 50, 3})) // 12
	fmt.Println(largestPerimeter([]int{1, 2, 1}))               // 0
	//fmt.Println(largestPerimeter([]int{5, 5, 50}))              // 10

}
