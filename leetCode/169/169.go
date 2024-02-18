package main

import "fmt"

func majorityElement(nums []int) int {
	counter := make(map[int]int)
	for _, v := range nums {
		counter[v]++
	}
	for k, v := range counter {
		if v > len(nums)/2 {
			return k
		}
	}
	return -1

}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3})) // 3
}
