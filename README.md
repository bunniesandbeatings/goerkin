# goerkin
A Gherkin DSL for Ginkgo

# Goals

* Provide the gherkin format for stories
    * without a special `*.feature` format which.
* Local step definitions instead of shared steps which often force *the wrong abstraction*
    * of course you can still import shared definitions as methods
* Lean on Ginkgo so as not to create a whole other system that needs extensive design and testing
* Promote imperative style tests
    * Dissuade the use of BeforeEach/AfterEach    

# TODO

* Send Regex params to steps
* Document using afterEach for deferred actions
* Tests as living documentation 

# Sample
```go
    var _ = Describe("running a total", func() {
        var (
            total int
        )
    
        steps := Define(func(define Definitions) {
            define.Given("The current total is cleared", func() {
            	total = 0
            })
    
            define.When("^I add 5$", func() {
            	total = total + 5
            })
    
            define.When("^I add 3$", func() {
                total = total + 3
            })
    
            define.Then("^The total is 8$", func() {
                Expect(total).To(Equal(8))
            })
        })
    
        Scenario("Adding", func() {
            steps.Given("The current total is cleared")
            
            steps.When("I add 5")
            steps.And("I add 3")
            
            steps.Then("The total is 8")
        })

        Scenario("Subtracting with inline definitions", func() {
            steps.Given("The current total is cleared")
            
            steps.When("I add 5")
            steps.And("I subtract 3", func() {
            	total = total - 3
            })
            
            steps.Then("The total is 2", func() {
            	Expect(total).To(Equal(2))
            })
        })
    })
```