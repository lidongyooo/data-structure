package main

import "fmt"

func quickSort(list []int) []int {
	length := len(list)

	if length < 2 {
		return list
	}

	pivot := list[0]
	var less []int
	var more []int
	for _, num := range list[1:] {
		if num > pivot {
			more = append(more, num)
		} else {
			less = append(less, num)
		}
	}

	return append(quickSort(less), append([]int{pivot}, quickSort(more)...)...)
}

func main() {
	slice := []int{8, 9, 120, 3, 50, 60, 5, 9}
	fmt.Println(quickSort(slice))
}
