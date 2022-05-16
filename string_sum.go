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
	s := strings.Split(input, "+")
	if len(s) != 2 {
		return "", fmt.Errorf("StringSum: %w", errorNotTwoOperands)
	}
	s1 := strings.TrimSpace(s[0])
	if len(s1) == 0 {
		return "", fmt.Errorf("StringSum: first operand is empty\n %w", errorEmptyInput)
	}
	s2 := strings.TrimSpace(s[1])
	if len(s2) == 0 {
		return "", fmt.Errorf("StringSum: second operand is empty\n %w", errorEmptyInput)
	}

	num1, err := strconv.Atoi(s1)
	if err != nil {
		return "", fmt.Errorf("StringSum: first operand not valid\n %w", err)
	}
	num2, err := strconv.Atoi(s2)
	if err != nil {
		return "", fmt.Errorf("StringSum: second operand not valid\n %w", err)
	}

	return strconv.Itoa(num1 + num2), nil
}
