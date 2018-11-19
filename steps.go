package goerkin

import (
	"fmt"
	"github.com/onsi/ginkgo"
)

type Steps struct {
	definitions definitions
}

type defineBodyFn func(Definitions)

func NewSteps() *Steps {
	return &Steps{
		definitions: definitions{},
	}
}

func Define(body ...interface{}) *Steps {
	steps := NewSteps()

	if len(body) > 0 {
		steps.Define(body[0].(func(Definitions)))
	}

	return steps
}

func (s *Steps) Define(body defineBodyFn) {
	body(s.definitions)
}

func (s *Steps) run(method, text string, override []bodyFn) {
	if len(override) > 0 {
		ginkgo.By(text, override[0])
	}

	var match bodyFn

	for re, body := range s.definitions {
		if re.MatchString(text) {
			if match != nil {
				ginkgo.Fail(fmt.Sprintf("Too many matches for `%s`", text))
				return
			}
			match = body
		}
	}

	if match == nil {
		template := fmt.Sprintf("define.%s(\"^%s$\", func() {})", method, text)
		ginkgo.Fail(fmt.Sprintf("No match for `%s`, try adding:\n%s", text, template))
		return
	}

	ginkgo.By(text, match)
}

func (s *Steps) Given(text string, body ...bodyFn) { s.run("Given", text, body) }
func (s *Steps) When(text string, body ...bodyFn)  { s.run("When", text, body) }
func (s *Steps) Then(text string, body ...bodyFn)  { s.run("Then", text, body) }
func (s *Steps) And(text string, body ...bodyFn)   { s.run("And", text, body) }

func (s *Steps) Run(text string)   { s.run("Step", text, nil) }
