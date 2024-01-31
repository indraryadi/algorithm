package main

import (
	"fmt"
	"math"
)

func binarySearch(data []int, find int) {
	left := 0
	right := len(data) - 1
	mid := int(math.Floor(float64(left) + float64(right)/2))

	// base case
	if left == mid && right == mid {
		if data[mid] == find {
			fmt.Println("found")
		} else {
			fmt.Println("not found")
		}
		return
	}

	if find <= data[mid] {
		right = mid
		binarySearch(data[:right], find)
	} else {
		left = mid
		binarySearch(data[left+1:], find)
	}
}

func main() {
	data := []int{0, 1, 3, 4, 6, 8, 12, 14, 15, 17, 19, 22, 25, 37, 38}
	binarySearch(data, 13)
}
