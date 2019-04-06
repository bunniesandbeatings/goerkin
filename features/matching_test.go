package features_test

import (
	. "github.com/bunniesandbeatings/goerkin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Over/Under-match experience", func() {
	steps := NewSteps()

	Scenario("Matching on more than one definition", func() {
		steps.Given("Two givens with similar regexes")
		steps.When("a given that matches both definitions is called")
		steps.Then("I see the given's causing me trouble")
	})

	steps.Define(func(define Definitions) {
		var (
			subject  *Steps
			failText string
		)

		var failFunc = func(message string, callerSkip ...int) {
			failText = failText + message + "\n"
		}

		define.Given(`^Two givens with similar regexes$`, func() {
			subject = NewSteps()
			subject.Fail = failFunc
			subject.Define(func(subjectDefinitions Definitions) {
				subjectDefinitions.Given(`I might (.*)`, func(match string) {})
				subjectDefinitions.Given(`I (.*) match`, func(might string) {})
			})
		})

		define.When(`^a given that matches both definitions is called$`, func() {
			subject.Given("I might match")
		})

		define.Then(`^I see the given's causing me trouble$`, func() {
			Expect(failText).To(ContainSubstring("Too many matches for `I might match`:"))
			Expect(failText).To(ContainSubstring("\t0: I might (.*)"))
			Expect(failText).To(ContainSubstring("\t1: I (.*) match"))
		})
	})
})
