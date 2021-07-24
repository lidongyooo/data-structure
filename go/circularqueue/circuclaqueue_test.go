package circularqueue

import (
	"fmt"
	"testing"
)

func TestMyCircularQueue_DeQueue(t *testing.T) {
	obj := Constructor(3)
	obj.EnQueue(1)
	obj.EnQueue(2)
	obj.EnQueue(3)
	obj.EnQueue(4)
	param_1 := obj.Rear()
	param_2 := obj.IsFull()

	obj.EnQueue(1)

	param_3 := obj.DeQueue()
	param_4 := obj.EnQueue(4)
	param_5 := obj.Rear()

	fmt.Println(param_1, param_2, param_3, param_4, param_5)
}