package features_test

import . "github.com/bunniesandbeatings/goerkin"

var sharedValue string

var sharedSteps = func(define Definitions) {
	define.Given(`^I am a shared step$`, func() {
		sharedValue = "shared step called"
	}, func() {
		sharedValue = ""
	})
}
