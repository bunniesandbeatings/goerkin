package goerkin

import (
	"fmt"
	"github.com/onsi/ginkgo"
	"regexp"
)

type Definitions interface {
	Given(re string, given bodyFn, after ...func())
	When(re string, when bodyFn, after ...func())
	Then(re string, then bodyFn)
}

type definitions map[*regexp.Regexp]func()

func (defs definitions) add(text string, body bodyFn, after []func()) {
	if len(after) > 0 {
		ginkgo.AfterEach(after[0])
	}

	re, err := regexp.Compile(text)
	if err != nil {
		panic(fmt.Sprintf("Could not compile %s with error %s", text, err))
	}

	defs[re] = body
}

func (defs definitions) Given(re string, given bodyFn, after ...func()) { defs.add(re, given, after) }
func (defs definitions) When(re string, when bodyFn, after ...func())   { defs.add(re, when, after) }
func (defs definitions) Then(re string, then bodyFn)                    { defs.add(re, then, nil) }

