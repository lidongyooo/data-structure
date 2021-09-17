package main

import (
	"fmt"
)

func binarySearch(list []int, target int) int {
	low, high := 0, len(list) - 1
	var (
		mid int
		guess int
	)

	for low <= high {
		mid = (low + high) / 2
		guess = list[mid]

		if guess > target {
			high = mid - 1
		} else if guess < target {
			low = mid + 1
		} else {
			return mid
		}
	}

	return -1
}

func binarySearchPlus(list []int, target int, low int, high int) int {
	mid := (low + high) / 2
	guess := list[mid]

	if low > high {
		return -1
	}

	if guess == target {
		return mid
	} else if guess < target {
		low = mid + 1
	} else {
		high = mid - 1
	}

	return binarySearchPlus(list, target, low, high)
}

func main() {
	list := []int{1, 5, 6, 9, 10, 90, 300}
	fmt.Println(binarySearchPlus(list, 444, 0, len(list) - 1))

	fmt.Println(binarySearch(list, 444))
}
