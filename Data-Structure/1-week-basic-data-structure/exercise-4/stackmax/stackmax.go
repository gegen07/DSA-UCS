package stackmax

import (
	stack "github.com/gegen07/Courses/Data-Structure/1-week-basic-data-structure/exercise-4/stack"
)

// StackMax is a special Stack for supporting max/push/pop functions
type StackMax struct {
	stk *stack.Stack
	maxElem int
}

// New function init a stackMax
func New() *StackMax {
	return &StackMax {stk: stack.New(), maxElem: 0}
}

// Push function insert values into StackMax
func (sm *StackMax) Push(value int) {
	if sm.stk.Empty() {
		sm.maxElem = value
		sm.stk.Push(value)
		return
	} else if value > sm.maxElem {
		sm.stk.Push(2*value-sm.maxElem)
		sm.maxElem = value
	} else {
		sm.stk.Push(value)
	}
}

// Pop function remove value from StackMax
func (sm *StackMax) Pop() bool {
	if sm.stk.Empty() {
		return false
	}

	topValue, _ := sm.stk.Top().(int)
	sm.stk.Pop()

	if topValue > sm.maxElem {
		sm.maxElem = 2 * sm.maxElem - topValue
	}

	return true
}

// Max funtion return StackMax's MaxElem 
func (sm *StackMax) Max() (int, bool) {
	if sm.stk.Empty() {
		return -1, true
	}

	return sm.maxElem, false
}

