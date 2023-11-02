package stack

import (
	"errors"
)

type Stack[T any] struct {
    memory []T
}

func New[T any]() *Stack[T] {
    return &Stack[T]{}
}

func (stack *Stack[T]) Push(v T) {
    stack.memory = append(stack.memory, v)
}

func (stack *Stack[T]) Pop() (T, error) {
    if stack.IsEmpty() {
        return stack.memory[0], errors.New("Stack is empty")
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
        return stack.memory[0], errors.New("Stack is empty")
    }
    return stack.memory[len(stack.memory) -1], nil
}

func (stack *Stack[T]) Size() int {
    return len(stack.memory)
}
