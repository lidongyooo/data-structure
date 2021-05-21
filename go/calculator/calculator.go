package calculator

import (
	"fmt"
	"strconv"
)

type Calculator struct {
	nums *StackNum
	operator *Stack
	expression string
}

var operatorPriority = map[string]int{
	"+" : 0,
	"-" : 0,
	"*" : 1,
	"/" : 1,
	"(" : 2,
	")" : 2,
}

func NewCalculator(expression string) *Calculator {
	return &Calculator{
		NewStackNum(10),
		NewStack(10),
		expression,
	}
}

func Run(expression string) int {
	c := NewCalculator(expression)
	return c.Calculate()
}

func (c *Calculator) Calculate() int {
	l := len(c.expression)

	for i := 0; i < l; i++ {
		
		switch val := c.expression[i]; val {
		
		case ' ':
			continue
			
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			j := i
			for j < l && c.expression[j] <= '9' && c.expression[j] >= '0' {
				j++
			}
			num,_ := strconv.Atoi(c.expression[i:j])
			i = j -1
			c.nums.Push(num)
			
		case '+', '-', '*', '/':
			prev := c.operator.Pop()

			if prev != "" && prev != "(" && operatorPriority[string(val)] <= operatorPriority[prev] {
				c.nums.Push(c.Calculation(prev))
				prev = c.operator.Pop()
			}

			if prev != "" {
				c.operator.Push(prev)
			}
			
			c.operator.Push(string(val))

		case '(':
			c.operator.Push(string(val))
		case ')':
			for o := c.operator.Pop(); o != "" && o != "("; o = c.operator.Pop() {
				c.nums.Push(c.Calculation(o))
			}
		default:
			panic("illegal character")
		}
	}

	for o := c.operator.Pop(); o != ""; o = c.operator.Pop() {
		c.nums.Push(c.Calculation(o))
	}

	return c.nums.Pop()
}

func (c *Calculator) Calculation(operator string) int {
	b := c.nums.Pop()
	a := c.nums.Pop()
	res := 0

	switch operator {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	}

	fmt.Printf("%d %s %d = %d\n", a, operator, b, res)

	return res
}