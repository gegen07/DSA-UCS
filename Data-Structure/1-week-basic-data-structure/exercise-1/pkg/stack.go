package stack

import(
	"log"
)

//Node is element of list
type Node struct {
	key interface{}
	next *Node
}

//Stack is built by singly linked list
type Stack struct {
	head *Node
	size int
}

//New function initialize a stack
func New() *Stack {
	return &Stack{ head: nil, size: 0 }
}

//Push function add a value
func (stack *Stack) Push(value interface{}) {
	newNode := new(Node)

	newNode.key = value
	newNode.next = stack.head
	stack.head = newNode
	stack.size++

}

//Pop funciton remove the first element
func (stack *Stack) Pop() (interface{}) {
	
	if !stack.Empty() {
		x := stack.head.key
		stack.head = stack.head.next
		stack.size--

		return x
	}
	log.Fatalf("underflow")
	return nil
}

//Top function return the first element
func (stack *Stack) Top() (interface{}) {

	if !stack.Empty() {
		return stack.head.key
	}

	return nil
}

//Empty function return if stack has some value
func (stack *Stack) Empty() (bool) {
	return stack.size == 0
}

//Size return th length of stack
func (stack *Stack) Size() (int) {
	return stack.size
}