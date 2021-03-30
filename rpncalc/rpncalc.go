package rpncalc

import (
	"fmt"
	"strconv"
	"strings"
)

type Calculator interface {
	Calculate(expr string) (int, error)
}

type calculatorImpl struct{}

func New() Calculator {
	return &calculatorImpl{}
}

func (c *calculatorImpl) Calculate(expr string) (int, error) {
	tokens := strings.Fields(expr)
	return c.eval(tokens)
}

func (c *calculatorImpl) eval(tokens []string) (int, error) {
	stack := []int{}
	for {
		fmt.Printf("tokens: %v\n", tokens)
		fmt.Printf("stack: %v\n", stack)
		if len(tokens) == 0 {
			if len(stack) == 1 {
				return stack[0], nil
			} else {
				return 0, fmt.Errorf("invalid syntax")
			}
		}
		token := tokens[0]
		tokens = tokens[1:]
		num, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, num)
			continue
		}
		y := stack[len(stack)-1]
		x := stack[len(stack)-2]
		stack = stack[:len(stack)-2]
		var answer int
		switch token {
		case "+":
			answer = x + y
		case "-":
			answer = x - y
		case "*":
			answer = x * y
		case "/":
			answer = x / y
		}
		stack = append(stack, answer)
	}
}
