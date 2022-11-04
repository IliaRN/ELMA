package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type MaxStack struct {
	stack []int
	max   int
}

func NewMaxStack() MaxStack {
	return MaxStack{max: math.MinInt64}
}
func (ms *MaxStack) Push(item int) {
	ms.stack = append(ms.stack, item)
	if item > ms.max {
		ms.max = item
	}
}
func (ms *MaxStack) Pop() {
	length := len(ms.stack)
	last := ms.stack[length-1]
	ms.stack = ms.stack[:length-1]
	if last != ms.max {
		return
	}
	ms.max = math.MinInt64
	for _, x := range ms.stack {
		if x > ms.max {
			ms.max = x
		}
	}
}

func (ms *MaxStack) GetMax() int {
	return ms.max
}

func main() {
	var n, val int
	stack := NewMaxStack()

	for {
		fmt.Println(
			"\nChoose stacks operation \n 1: Push \n 2: Pop \n 3: Max element of stack \n\n To exit: any key")
		reader := bufio.NewReader(os.Stdin)
		_, err := fmt.Fscan(reader, &n)
		if err != nil {
			os.Exit(1)
		}
		switch n {
		case 1:
			fmt.Println("Enter the value you want to append:")
			_, err := fmt.Fscan(reader, &val)
			if err != nil {
				os.Exit(1)
			}
			stack.Push(val)
			fmt.Printf("%d was successfully pushed to stack. Stacks values: %v\n", val, stack.stack)
		case 2:
			toPop := stack.stack[len(stack.stack)-1]
			stack.Pop()
			fmt.Printf("%d was successfully removed from stack. Stacks values: %v\n", toPop, stack.stack)
		case 3:
			if len(stack.stack) != 0 {
				fmt.Printf("Max value of stack %v is %d", stack.stack, stack.GetMax())
			} else {
				fmt.Printf("Stack is empty. Max value is 0\n")
			}
		default:
			return
		}
	}
}
