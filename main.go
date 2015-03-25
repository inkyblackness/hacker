package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/inkyblackness/hacker/cmd"
	"github.com/inkyblackness/hacker/styling"
)

type testTarget struct {
}

func (target *testTarget) Load(path1, path2 string) string {
	return "hello <" + path1 + ">, <" + path2 + ">"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	quit := false
	style := newStandardStyle()
	target := &testTarget{}
	eval := cmd.NewEvaluater(style, target)

	for !quit {
		input := queryUserInput(style, scanner)

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

func queryUserInput(style styling.Style, scanner *bufio.Scanner) string {
	fmt.Printf(style.Prompt()("> "))
	scanner.Scan()

	return strings.Trim(scanner.Text(), " ")
}
