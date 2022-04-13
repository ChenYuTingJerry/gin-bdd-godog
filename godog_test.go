package main

import (
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/spf13/pflag"

	"github.com/ChenYuTingJerry/gin-bdd-godog/features/cucumber"
)

var opts = godog.Options{
	//Output: colors.Colored(os.Stdout),
	Format: "progress",
}

func init() {
	//godog.BindCommandLineFlags("godog.", &opts)
}

func TestMain(m *testing.M) {
	pflag.Parse()
	opts.Paths = pflag.Args()

	status := godog.TestSuite{
		Name:                 "godogs",
		TestSuiteInitializer: MainScenarioSuite,
		ScenarioInitializer:  MainScenarioContext,
		Options:              &opts,
	}.Run()

	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}

func MainScenarioSuite(ctx *godog.TestSuiteContext) {
	cucumber.InitializeTestSuite(ctx)
}

func MainScenarioContext(s *godog.ScenarioContext) {
	cucumber.InitializeScenario(s)
}
