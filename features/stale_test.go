package features_test

import (
	. "github.com/bunniesandbeatings/goerkin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stale definitions", func() {
	steps := NewSteps()

	Scenario("Showing stale steps when they don't match in a test run", func() {
		steps.Given("three definitions")
		steps.When("a test matching only one definition is run")
		steps.Then("I see the two definitions that were not run marked as stale")
	})

	steps.Define(func(define Definitions) {
		var (
			subject *Steps
		)

		define.Given(`^three definitions$`, func() {
			subject = NewSteps()

			subject.Define(func(subjectDefinitions Definitions) {
				subjectDefinitions.Given(`I never get matched`, func() {})
				subjectDefinitions.When(`I don't match`, func() {})
				subjectDefinitions.Then(`actually, I match`, func() {})
			})
		})

		define.When(`^a test matching only one definition is run$`, func() {
			subject.Then("actually, I match")
		})

		define.When(`^goulag`, func() {
		})
		define.Then(`^I see the two definitions that were not run marked as stale$`, func() {
			Expect(subject.UnusedSteps()).To(ContainElement("I never get matched"))
			Expect(subject.UnusedSteps()).To(ContainElement("I don't match"))
		})
	})
})
