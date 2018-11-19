package goerkin

import (
	"fmt"
	"regexp"
)

type Definitions interface {
	Given(re string, body bodyFn)
	When(re string, body bodyFn)
	Then(re string, body bodyFn)
}

type definitions map[*regexp.Regexp]func()

func (d definitions) add(text string, body bodyFn) {
	re, err := regexp.Compile(text)
	if err != nil {
		panic(fmt.Sprintf("Could not compile %s with error %s", text, err))
	}

	d[re] = body
}

func (d definitions) Given(text string, body bodyFn) { d.add(text, body) }
func (d definitions) When(text string, body bodyFn) { d.add(text, body) }
func (d definitions) Then(text string, body bodyFn) { d.add(text, body) }

