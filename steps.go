package goerkin

import (
	"fmt"
	"reflect"

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

type matchT struct {
	body   interface{}
	params []string
}

func (s *Steps) run(method, text string, override []bodyFn) {
	if len(override) > 0 {
		ginkgo.By(text, override[0])
		return
	}

	match := matchT{}

	for re, body := range s.definitions {
		stringMatches := re.FindStringSubmatch(text)
		if stringMatches == nil {
			continue
		}

		if match.body != nil {
			ginkgo.Fail(fmt.Sprintf("Too many matches for `%s`", text))
			return
		}

		match.body = body
		match.params = stringMatches[1:]
	}

	if match.body == nil {
		templateBacktick := fmt.Sprintf("define.%s(`^%s$`, func() {})", method, text)
		templateDouble := fmt.Sprintf("define.%s(\"^%s$\", func() {})", method, text)
		// FIXME: matches fail here, they should show the definition that failed
		ginkgo.Fail(fmt.Sprintf("No match for `%s`, try adding:\n%s\nor:\n%s\n", text, templateBacktick, templateDouble))
		return
	}

	ginkgo.By(text, func() {
		switch match.body.(type) {
		case func():
			match.body.(func())()
		default:
			matchValue := reflect.ValueOf(match.body)

			in := make([]reflect.Value, len(match.params))

			for paramIndex := range in {
				in[paramIndex] = reflect.ValueOf(match.params[paramIndex])
			}

			matchValue.Call(in)

			//ginkgo.Fail(fmt.Sprintf("Could not match function call for \"%s\"\nlooking for:%v", text, reflect.TypeOf(match)))
		}
	})
}

func (s *Steps) Given(text string, body ...bodyFn) { s.run("Given", text, body) }
func (s *Steps) When(text string, body ...bodyFn)  { s.run("When", text, body) }
func (s *Steps) Then(text string, body ...bodyFn)  { s.run("Then", text, body) }
func (s *Steps) And(text string, body ...bodyFn)   { s.run("And", text, body) }

func (s *Steps) Run(text string) { s.run("Step", text, nil) }
