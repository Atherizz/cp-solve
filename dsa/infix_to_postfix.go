package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
	if len(s.items) == 0 {
		var zero T
		return zero
	} else {
		top := s.items[len(s.items)-1]
		s.items = s.items[:len(s.items)-1]
		return top
	}
}

func (s *Stack[T]) Peek() T {
	return s.items[len(s.items)-1]
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func Postfix(expr string) string {

	degree := map[rune]int{
		'(': 0,
		')': 0,
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
		'^': 3,
	}

	stack := Stack[rune]{
		items: []rune{},
	}

	var postfix string
	var num string

	for i, ch := range expr {
		if unicode.IsDigit(ch) {
			num += string(ch)
			if i != len(expr) - 1 && !unicode.IsDigit(rune(expr[i+1])) {
				postfix += num + " "
				num = ""
			} else {

			}
		} else if ch == '(' {
			stack.Push(ch)
		} else if ch == ')' {
			for !stack.IsEmpty() && stack.Peek() != '(' {
				operand := stack.Pop()
				postfix += string(operand) + " "
			}
			stack.Pop()
		} else {
			for !stack.IsEmpty() && degree[ch] <= degree[stack.Peek()] {
				operator := stack.Pop()
				postfix += string(operator) + " "
			}
			stack.Push(ch)
		}

	}

	if num != "" {
		postfix += num + " "
	}

	for !stack.IsEmpty() {
		postfix += string(stack.Pop()) + " "
	}

	return postfix
}

func Calculate(postfix string) int {

	equation := strings.Split(postfix, " ")

	stack := Stack[int]{
		items: []int{},
	}

	var x, y, h int
	var opt string

	for _, e := range equation {
		num, err := strconv.Atoi(e)
		if err == nil {
			stack.Push(num)
		} else {
			opt = e
			x = stack.Pop()
			y = stack.Pop()

			if opt == "+" {
				h = y + x
			} else if opt == "-" {
				h = y - x
			} else if opt == "*" {
				h = y * x
			} else if opt == "/" {
				h = y / x
			} else if opt == "^" {
				h = int(math.Pow(float64(y), float64(x)))
			} 
			stack.Push(h)
		}

	}
	
	return stack.Pop()
}

func main() {
	postfix := Postfix("2+14*(9-5)/3")
	// postfix := Postfix("3+-2")
	fmt.Println(postfix)
	fmt.Println(Calculate(postfix))
}
