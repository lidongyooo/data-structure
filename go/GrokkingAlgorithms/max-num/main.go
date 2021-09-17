package main

import "fmt"

func max(list []int, low int) (int, int) {
	length := len(list)

	if length - 1 == low {
		return list[low], low
	}

	num, index := max(list, low + 1)

	if list[low] > num {
		return list[low], low
	} else {
		return num, index
	}
}

func main() {
	slice := []int{60, 4, 3}
	fmt.Println(max(slice, 0))
}
