package calculator

type Stack struct {
	counter uint
	capacity uint
	slices []string
}

func NewStack(capacity uint) *Stack {
	return &Stack{
		0,
		capacity,
		make([]string, capacity),
	}
}

func (s *Stack) Pop () string {
	if s.counter == 0 {
		return ""
	}

	s.counter--
	return s.slices[s.counter]
}

func (s *Stack) Push (val string) bool {
	if s.counter >= s.capacity {
		s.slices = append(s.slices, val)
		s.capacity = uint(cap(s.slices))
	} else {
		s.slices[s.counter] = val
	}

	s.counter++
	return true
}

type StackNum struct {
	counter uint
	capacity uint
	slices []int
}

func NewStackNum(capacity uint) *StackNum {
	return &StackNum{
		0,
		capacity,
		make([]int, capacity),
	}
}

func (s *StackNum) Pop () int {
	if s.counter == 0 {
		return -999999
	}

	s.counter--
	return s.slices[s.counter]
}

func (s *StackNum) Push (val int) bool {
	if s.counter >= s.capacity {
		s.slices = append(s.slices, val)
		s.capacity = uint(cap(s.slices))
	} else {
		s.slices[s.counter] = val
	}

	s.counter++
	return true
}