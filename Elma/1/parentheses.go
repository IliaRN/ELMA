package main

import "fmt"

func ParenthesesTask(inputString string) string {
	if inputString == "" {
		return "YES"
	}

	mapOfCases := map[rune]rune{
		'}': '{', // { = 123, } = 125
		')': '(', // ( = 40,  ) = 41
		']': '[', // [ = 91,  ] = 93
	}

	stack := make([]rune, 0, 1)

	for _, sym := range inputString {
		switch sym {
		case 123, 40, 91:
			stack = append(stack, sym)
		case 125, 41, 93:
			if len(stack) == 0 || stack[len(stack)-1] != mapOfCases[sym] {
				return "НЕТ"
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) == 0 {
		return "YES"
	}
	return "НЕТ"
}

func main() {
	s := "(123{]})"
	result := ParenthesesTask(s)
	fmt.Println(result)
}
