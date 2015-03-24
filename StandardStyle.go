package main

import (
	"github.com/inkyblackness/hacker/styling"

	"github.com/fatih/color"
)

type standardStyle struct {
	prompt styling.StyleFunc
	err    styling.StyleFunc
}

func newStandardStyle() *standardStyle {
	style := &standardStyle{
		prompt: color.New(color.FgGreen).SprintFunc(),
		err:    color.New(color.FgRed, color.Bold).SprintFunc()}

	return style
}

func (style *standardStyle) Prompt() styling.StyleFunc {
	return style.prompt
}

func (style *standardStyle) Error() styling.StyleFunc {
	return style.err
}
