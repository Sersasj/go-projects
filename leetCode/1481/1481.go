// 1481. Least Number of Unique Integers after K Removals
// Medium
// Topics
// Companies
// Hint
// Given an array of integers arr and an integer k. Find the least number of unique integers after removing exactly k elements.

// Example 1:

// Input: arr = [5,5,4], k = 1
// Output: 1
// Explanation: Remove the single 4, only 5 is left.
// Example 2:
// Input: arr = [4,3,1,1,3,3,2], k = 3
// Output: 2
// Explanation: Remove 4, 2 and either one of the two 1s or three 3s. 1 and 3 will be left.

// Constraints:

// 1 <= arr.length <= 10^5
// 1 <= arr[i] <= 10^9
// 0 <= k <= arr.length
package main

import (
	"fmt"
	"sort"
)

func findLeastNumOfUniqueInts(arr []int, k int) int {

	// count the frequency of each number
	counter := make(map[int]int)
	for _, v := range arr {
		counter[v]++
	}

	// sort frequencies
	freqs := make([]int, 0, len(counter))
	for _, v := range counter {
		freqs = append(freqs, v)
	}
	sort.Ints(freqs)

	numbers_removed := 0
	for _, count := range freqs {
		if k >= count {
			k -= count
			numbers_removed++
		} else {
			break
		}
	}

	return len(freqs) - numbers_removed
}

func main() {
	fmt.Println(findLeastNumOfUniqueInts([]int{5, 5, 4}, 1))             // 1
	fmt.Println(findLeastNumOfUniqueInts([]int{4, 3, 1, 1, 3, 3, 2}, 3)) // 2
	fmt.Println(findLeastNumOfUniqueInts([]int{1, 2, 3, 4, 5}, 2))       // Output: 3
	fmt.Println(findLeastNumOfUniqueInts([]int{1, 1, 3, 3, 3}, 3))       // Output: 3

}
