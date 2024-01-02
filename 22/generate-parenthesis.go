package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%v\n", generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	if n < 1 {
		return []string{}
	}

	var chars, stack []string
	backtrack(n, 0, 0, &chars, &stack)
	return stack
}

func backtrack(n, openN, closeN int, chars, stack *[]string) {
	if openN == closeN && closeN == n {
		*stack = append(*stack, strings.Join(*chars, ""))
		return
	}

	if openN < n {
		*chars = append(*chars, "(")
		backtrack(n, openN+1, closeN, chars, stack)
		*chars = (*chars)[0 : len(*chars)-1]
	}

	if closeN < openN {
		*chars = append(*chars, ")")
		backtrack(n, openN, closeN+1, chars, stack)
		*chars = (*chars)[0 : len(*chars)-1]
	}
}
