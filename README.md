# goerkin
A Gherkin DSL for Ginkgo

Inspired by [Robbie Clutton's simple_bdd](https://github.com/robb1e/simple_bdd)

# Goals

* Provide the gherkin format for stories
    * without a special `*.feature` format
* Local step definitions instead of shared steps which often force *the wrong abstraction*
    * of course you can still import shared definitions as methods
* Lean on Ginkgo so as not to create a whole other system that needs extensive design and testing
* Promote imperative style tests
    * Dissuade the use of BeforeEach/AfterEach    

# TODO

* Send Regex params to steps
* Tests as living documentation 

# Samples

## Simple usage
```go
    import (
        . "github.com/onsi/ginkgo"
        . "github.com/onsi/gomega"
        . "github.com/bunniesandbeatings/goerkin"
    )

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

## Calling steps from within other steps
```go
    var _ = Describe("running a total", func() {
        var (
            total int
            steps *Steps
        )
    
        steps = Define(func(define Definitions) {
            define.Given("The current total is cleared", func() {
            	total = 0
            })
    
            define.When("^I add 5$", func() {
            	total = total + 5
            })
    
            define.When("^I add 3$", func() {
                total = total + 3
            })

            define.When("^I add 5 and 3 to the total$", func() {
                steps.Run("I add 5")
                steps.Run("I add 3")
            })
            
            define.Then("^The total is 8$", func() {
                Expect(total).To(Equal(8))
            })
        })
    
        Scenario("Adding", func() {
            steps.Given("The current total is cleared")
            
            steps.When("I add 5 and 3 to the total")
            
            steps.Then("The total is 8")
        })
    })
```
## Features first

I like my features at the top of the file. You can do that:
```go
    import (
        . "github.com/onsi/ginkgo"
        . "github.com/onsi/gomega"
        . "github.com/bunniesandbeatings/goerkin"
    )

    var _ = Describe("running a total", func() {
        var (
            total int
        )
    
        steps := NewSteps()

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
        
        
        steps.Define(func(define Definitions) {
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
    
    })

```


## Cleanup Steps

`Givens` and `Whens` support cleanup methods (they become ginkgo AfterEach blocks)

```go
    import (
        . "github.com/onsi/ginkgo"
        . "github.com/onsi/gomega"
        . "github.com/bunniesandbeatings/goerkin"
    )

    var _ = Describe("Daemonize works", func() {
        var (
            app *exec.Cmd
        )
    
        steps := NewSteps()

        Scenario("Running", func() {
            steps.Given("My server is running")
            
            steps.When("I visit it's url")
            
            steps.Then("It responds")
        })

        
        
        steps.Define(func(define Definitions) {
            define.Given("My server is running",
            	func() {
            	    app := startMyServer()
                },
                func() {    // this becomes an AfterEach block.
                	stopMyServer(app)
                }
            )
    
            ... blah, blah blah blablah ...
        })
        
            
    })
```
