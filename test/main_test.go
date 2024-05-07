package test

import (
	"testing"

	"github.com/cucumber/godog"
	"github.com/denniskreussel/scm/test/cucumber/step_definitions"
)

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: func(s *godog.ScenarioContext) {
			// Add step definitions here.
			s.Step(`^I add up the numbers (\d+) and (\d+)$`, step_definitions.IAddUpTheNumbersAnd)
			s.Step(`^I should get (\d+)$`, step_definitions.IShouldGet)
		},
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"cucumber/features"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
