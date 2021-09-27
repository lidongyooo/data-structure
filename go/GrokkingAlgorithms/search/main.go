package main

import (
	"GrokkingAlgorithms/search/queue"
	"fmt"
)

func graphSearch(graph map[string][]string) {
	queue := queue.Init()
	queue.Push(graph["yours"]...)

	checked := make(map[string]uint8)

	for !queue.IsEmpty() {
		v := queue.Pop()

		if _, ok := checked[v]; ok {
			continue
		}

		if v == "d" {
			fmt.Printf("got it：%v", v)
			return
		} else {
			if _, ok := graph[v]; ok {
				queue.Push(graph[v]...)
			}

			checked[v] = 0
		}
	}
}

func main() {
	graph := make(map[string][]string)
	graph["yours"] = []string{"a1", "a2", "a3", "a4"}
	graph["a1"] = []string{"b1", "b2", "b3"}
	graph["a2"] = []string{"b1"}
	graph["b2"] = []string{"c1", "c2"}
	graph["c1"] = []string{"d"}

	graphSearch(graph)
}
