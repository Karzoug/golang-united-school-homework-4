package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	if len(input) == 0 {
		return "", fmt.Errorf("StringSum: %w", errorEmptyInput)
	}

	s := make([]string, 0, 2)
	expr := strings.TrimSpace(input)
	prev := len(expr)
	for i := len(expr) - 1; i >= 0; i-- {
		if expr[i] == '-' || expr[i] == '+' {
			s = append(s, strings.ReplaceAll(expr[i:prev], " ", ""))
			prev = i
		}
	}

	if prev != 0 {
		s = append(s, strings.ReplaceAll(expr[0:prev], " ", ""))
	}

	if len(s) != 2 {
		return "", fmt.Errorf("StringSum: %w", errorNotTwoOperands)
	}

	if len(s[1]) == 0 {
		return "", fmt.Errorf("StringSum: first operand is empty\n %w", errorEmptyInput)
	}
	if len(s[0]) == 0 {
		return "", fmt.Errorf("StringSum: second operand is empty\n %w", errorEmptyInput)
	}

	num1, err := strconv.Atoi(s[1])
	if err != nil {
		return "", fmt.Errorf("StringSum: first operand not valid\n %w", err)
	}
	num2, err := strconv.Atoi(s[0])
	if err != nil {
		return "", fmt.Errorf("StringSum: second operand not valid\n %w", err)
	}

	return strconv.Itoa(num1 + num2), nil
}
