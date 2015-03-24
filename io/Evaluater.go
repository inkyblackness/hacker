package io

import (
	"github.com/inkyblackness/hacker/styling"
)

// Evaluater wraps the Evaluate function to process some input
type Evaluater struct {
	style styling.Style
}

// NewEvaluater returns an evaluater processing input strings.
func NewEvaluater(style styling.Style) *Evaluater {
	eval := &Evaluater{style: style}

	return eval
}

// Evaluate takes the given input, processes it and returns an evaluation result.
func (eval *Evaluater) Evaluate(input string) string {
	return eval.style.Error()("Unknown command: <", input, ">")
}
