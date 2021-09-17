package main

import "fmt"

func sum(arr []int) int {
	length := len(arr)

	if length == 0 {
		return 0
	}

	return arr[length - 1] + sum(arr[0 : (length - 1)])
}

func main() {
	slices := []int{1, 2, 3}

	fmt.Println(sum(slices))
}
