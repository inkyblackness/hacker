package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/inkyblackness/hacker/cmd"
	"github.com/inkyblackness/hacker/core"
	"github.com/inkyblackness/hacker/styling"
)

const (
	// Version contains the current version number
	Version = "0.1.0"
)

func main() {
	style := newStandardStyle()
	target := core.NewHacker(style)
	eval := cmd.NewEvaluater(style, target)
	scanner := bufio.NewScanner(os.Stdin)
	quit := false

	fmt.Printf("%s\n", style.Prompt()(`InkyBlackness Hacker v.`, Version))
	fmt.Printf("%s\n", style.Prompt()(`Type "quit" to exit`))
	fmt.Printf("%s\n", style.Prompt()(`Remember to keep backups! ...and to salt the fries!`))

	for !quit {
		input := queryUserInput(target.CurrentDirectory(), style, scanner)

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

func queryUserInput(prompt string, style styling.Style, scanner *bufio.Scanner) string {
	fmt.Printf(style.Prompt()(prompt, "> "))
	scanner.Scan()

	return strings.Trim(scanner.Text(), " ")
}
