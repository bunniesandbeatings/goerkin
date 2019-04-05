package features_test

import (
	"strconv"

	. "github.com/bunniesandbeatings/goerkin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Inline Definitions", func() {
	var total int

	steps := NewSteps()

	Scenario("Everything inline", func() {
		steps.Given("the current total is cleared", func() {
			total = 0
		})

		steps.When("3 is added", func() {
			total = total + 3
		})

		steps.When("2 is subtracted", func() {
			total = total - 2
		})

		steps.Then("the total is 1", func() {
			Expect(total).To(Equal(1))
		})

	})

})

var _ = Describe("Reuse Definitions", func() {
	var total int

	steps := Define(func(define Definitions) {
		define.Given("^the current total is cleared$", func() {
			total = 0
		})

		define.When("^3 is added$", func() {
			total = total + 3
		})

		define.When("^2 is subtracted$", func() {
			total = total - 2
		})

		define.Then("^the total is 1$", func() {
			Expect(total).To(Equal(1))
		})
	})

	Scenario("No definitions", func() {
		steps.Given("the current total is cleared")

		steps.When("3 is added")
		steps.And("2 is subtracted")

		steps.Then("the total is 1")
	})
})

var _ = Describe("Scenario first", func() {
	steps := NewSteps()

	Scenario("Scenario before steps", func() {
		steps.Then("it works")
	})

	steps.Define(func(define Definitions) {
		define.Then("^it works$", func() {
			Succeed()
		})
	})
})

var _ = Describe("After Definitions", func() {
	var (
		givenCalled      = false
		whenCalled       = false
		thenCalled       = false
		givenAfterCalled = false
		whenAfterCalled  = false
		thenAfterCalled  = false
	)
	steps := NewSteps()

	Scenario("After functions are called after all of the other functions", func() {
		steps.Given("given calls are set to true", func() {
			// This runs 1st
			givenCalled = true
			Expect(givenCalled).To(BeTrue())
			Expect(whenCalled).To(BeFalse())
			Expect(thenCalled).To(BeFalse())

			Expect(givenAfterCalled).To(BeFalse())
			Expect(whenAfterCalled).To(BeFalse())
			Expect(thenAfterCalled).To(BeFalse())
		}, func() {
			// This runs 4th
			givenAfterCalled = true
			Expect(givenCalled).To(BeTrue())
			Expect(whenCalled).To(BeTrue())
			Expect(thenCalled).To(BeTrue())

			Expect(givenAfterCalled).To(BeTrue())
			Expect(whenAfterCalled).To(BeFalse())
			Expect(thenAfterCalled).To(BeFalse())
		})

		steps.When("when calls are set to true", func() {
			// This runs 2nd
			whenCalled = true
			Expect(givenCalled).To(BeTrue())
			Expect(whenCalled).To(BeTrue())
			Expect(thenCalled).To(BeFalse())

			Expect(givenAfterCalled).To(BeFalse())
			Expect(whenAfterCalled).To(BeFalse())
			Expect(thenAfterCalled).To(BeFalse())
		}, func() {
			// This runs 5th
			whenAfterCalled = true
			Expect(givenCalled).To(BeTrue())
			Expect(whenCalled).To(BeTrue())
			Expect(thenCalled).To(BeTrue())

			Expect(givenAfterCalled).To(BeTrue())
			Expect(whenAfterCalled).To(BeTrue())
			Expect(thenAfterCalled).To(BeFalse())
		})

		steps.Then("then calls are set to true", func() {
			// This runs 3rd
			thenCalled = true
			Expect(givenCalled).To(BeTrue())
			Expect(whenCalled).To(BeTrue())
			Expect(thenCalled).To(BeTrue())

			Expect(givenAfterCalled).To(BeFalse())
			Expect(whenAfterCalled).To(BeFalse())
			Expect(thenAfterCalled).To(BeFalse())
		}, func() {
			// This runs 6th
			thenAfterCalled = true
			Expect(givenCalled).To(BeTrue())
			Expect(whenCalled).To(BeTrue())
			Expect(thenCalled).To(BeTrue())

			Expect(givenAfterCalled).To(BeTrue())
			Expect(whenAfterCalled).To(BeTrue())
			Expect(thenAfterCalled).To(BeTrue())
		})
	})
})

var _ = Describe("Groups as string params", func() {
	var total int
	steps := NewSteps()

	Scenario("Use RE groups as params", func() {
		steps.Given("4 and 6")
		steps.Then("the answer is 10")
	})

	steps.Define(func(define Definitions) {
		define.Given(`^(\d+) and (\d+)$`, func(first, second string) {
			firstInt, _ := strconv.Atoi(first)
			secondInt, _ := strconv.Atoi(second)
			total = firstInt + secondInt
		})
		define.Then(`^the answer is (\d+)$`, func(answer string) {
			answerInt, _ := strconv.Atoi(answer)
			Expect(total).To(Equal(answerInt))
		})
	})

})

var _ = Describe("Shared Steps without the framework", func() {
	steps := NewSteps()

	Scenario("Use a shared step", func() {
		steps.Given("I am a shared step")
		steps.Then("I can depend upon it")
	})

	steps.Define(func(define Definitions) {
		sharedSteps(define) // inline addition

		define.Then(`^I can depend upon it$`, func() {
			Expect(sharedValue).To(Equal("shared step called"))
		})
	})
})

var _ = Describe("Shared Steps with the framework", func() {
	steps := NewSteps()

	Scenario("Use a shared step", func() {
		steps.Given("I am a shared step")
		steps.Then("I can depend upon it")
	})

	steps.Define(
		sharedSteps, // framework addition
		func(define Definitions) {
			define.Then(`^I can depend upon it$`, func() {
				Expect(sharedValue).To(Equal("shared step called"))
			})
		},
	)
})
