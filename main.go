package main

import (
	"fmt"
	"strings"

	"github.com/inkyblackness/hacker/io"
	"github.com/inkyblackness/hacker/styling"
)

func main() {
	quit := false
	style := newStandardStyle()
	eval := io.NewEvaluater(style)

	for !quit {
		input := queryUserInput(style)

		if input != "" {
			if input == "quit" {
				quit = true
			} else {
				result := eval.Evaluate(input)
				fmt.Printf("%s\n", result)
			}
		}
	}
}

func queryUserInput(style styling.Style) string {
	var input string

	fmt.Printf(style.Prompt()("> "))
	fmt.Scanln(&input)

	return strings.Trim(input, " ")
}
