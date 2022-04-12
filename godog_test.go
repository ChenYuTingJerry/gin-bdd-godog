package main

import (
	"github.com/cucumber/godog"

	"github.com/ChenYuTingJerry/gin-bdd-godog/features/cucumber"
)

func MainScenarioContext(s *godog.ScenarioContext) {
	cucumber.InitializeScenario(s)
}
