package calculator

import (
	"fmt"
	"testing"
)

func TestStack_Pop(t *testing.T) {
	stack := NewStack(2)
	stack.Push("a")
	stack.Push("b")
	stack.Push("c")
	stack.Pop()

	stack.Push("c")

	val := stack.Pop()
	fmt.Println(stack.slices)

	if val != "c" {
		t.Error("fail")
	}
}
