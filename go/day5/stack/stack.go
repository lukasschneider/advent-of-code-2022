package stack

import (
	"errors"
)

type Number interface {
    uint8 | int16 | int32 | int64 
}


type Stack[T Number] struct {
    memory []T
}

func New[T Number](elems ...T) *Stack[T] {
    stack := &Stack[T]{}
    for _, e := range elems {
        stack.Push(e)
    }

    return stack
}

func (stack *Stack[T]) Push(v T) {
    stack.memory = append(stack.memory, v)
}

func (stack *Stack[T]) Pop() (T, error) {
    if stack.IsEmpty() {
        return 0, errors.New("Stack is empty")
    }
    var item T = stack.memory[len(stack.memory) -1]
    stack.memory = stack.memory[:len(stack.memory) -1]
    return item, nil
}

func (stack *Stack[T]) IsEmpty() bool {
    return len(stack.memory) == 0
}

func (stack *Stack[T]) Peek() (T, error) {
    if stack.IsEmpty() {
        return 0, errors.New("Stack is empty")
    }
    return stack.memory[len(stack.memory) -1], nil
}

func (stack *Stack[T]) Size() int {
    return len(stack.memory)
}
